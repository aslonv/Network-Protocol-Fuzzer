package distributed

import (
    "context"
    "log"
    "net"
    "sync"

    "google.golang.org/grpc"
)

type Master struct {
    tasks   chan Task
    results chan FuzzingResult
    wg      sync.WaitGroup
    UnimplementedFuzzingServiceServer
}

func NewMaster() *Master {
    return &Master{
        tasks:   make(chan Task, 100),
        results: make(chan FuzzingResult, 100),
    }
}

func (m *Master) StartServer(port string) error {
    listener, err := net.Listen("tcp", port)
    if err != nil {
        return err
    }

    grpcServer := grpc.NewServer()
    RegisterFuzzingServiceServer(grpcServer, m)

    log.Printf("Master node listening on %s", port)
    return grpcServer.Serve(listener)
}

func (m *Master) GetTask(ctx context.Context, req *WorkerRequest) (*Task, error) {
    task := <-m.tasks
    return &task, nil
}

func (m *Master) ReportResult(ctx context.Context, result *FuzzingResult) (*ResultAck, error) {
    m.results <- *result
    return &ResultAck{Message: "Result received"}, nil
}

func (m *Master) DistributeTasks() {
    for i := 0; i < 1000; i++ {
        m.tasks <- Task{TaskId: int32(i), Data: []byte{byte(i)}}
    }
    close(m.tasks)
    m.wg.Wait()
}
