package distributed

import (
    "context"
    "log"

    "google.golang.org/grpc"
)

type Worker struct {
    id     int
    master string
    conn   *grpc.ClientConn
    client FuzzingServiceClient
}

func NewWorker(id int, masterAddress string) (*Worker, error) {
    conn, err := grpc.Dial(masterAddress, grpc.WithInsecure())
    if err != nil {
        return nil, err
    }

    return &Worker{
        id:     id,
        master: masterAddress,
        conn:   conn,
        client: NewFuzzingServiceClient(conn),
    }, nil
}

func (w *Worker) Start() {
    log.Printf("Worker %d started.", w.id)

    for {
        task, err := w.client.GetTask(context.Background(), &WorkerRequest{WorkerId: int32(w.id)})
        if err != nil {
            log.Printf("Worker %d encountered an error: %v", w.id, err)
            break
        }

        success := w.processTask(task)
        _, err = w.client.ReportResult(context.Background(), &FuzzingResult{
            TaskId:  task.TaskId,
            Success: success,
        })
        if err != nil {
            log.Printf("Worker %d failed to report result: %v", w.id, err)
        }
    }

    w.conn.Close()
}

func (w *Worker) processTask(task *Task) bool {
    log.Printf("Worker %d processing task %d", w.id, task.TaskId)

    // 1. Load the protocol definition
    proto, err := protocol.LoadProtocolDefinition("path/to/protocol.json")
    if err != nil {
        log.Printf("Worker %d failed to load protocol definition: %v", w.id, err)
        return false
    }

    // 2. Initialize the fuzzing engine with the protocol definition
    engine := engine.NewEngine(proto)

    // 3. Generate a packet based on the task data and mutate it
    packet := engine.GeneratePacket()

    // Apply mutation strategies if needed
    mutatedPacket := engine.MutatePacket(packet)

    // 4. Send the packet to the target network service
    success := w.sendPacket(mutatedPacket)

    return success
}

func (w *Worker) sendPacket(packet []byte) bool {
    // Placeholder target address and port
    targetAddr := "127.0.0.1:8080"

    // Create a connection to the target service
    conn, err := net.Dial("tcp", targetAddr)
    if err != nil {
        log.Printf("Worker %d failed to connect to target: %v", w.id, err)
        return false
    }
    defer conn.Close()

    // Send the packet
    _, err = conn.Write(packet)
    if err != nil {
        log.Printf("Worker %d failed to send packet: %v", w.id, err)
        return false
    }

    // Receive the response
    response := make([]byte, 1024)
    _, err = conn.Read(response)
    if err != nil {
        log.Printf("Worker %d failed to receive response: %v", w.id, err)
        return false
    }

    // 5. Analyze the response (Placeholder logic)
    if isValidResponse(response) {
        log.Printf("Worker %d received valid response", w.id)
        return true
    } else {
        log.Printf("Worker %d received invalid response", w.id)
        return false
    }
}

func isValidResponse(response []byte) bool {
    // Placeholder: Implement response validation logic
    // Example: Check if the response contains an expected pattern
    expectedPattern := []byte("OK")
    return bytes.Contains(response, expectedPattern)
}

