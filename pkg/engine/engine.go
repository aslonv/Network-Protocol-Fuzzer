package engine

import (
    "crypto/rand"
    "log"
    "math/rand"
    "network-protocol-fuzzer/pkg/protocol"
    "time"
)

type Engine struct {
    protocol *protocol.Protocol
}

func NewEngine(proto *protocol.Protocol) *Engine {
    return &Engine{
        protocol: proto,
    }
}

func (e *Engine) GeneratePacket() []byte {
    packet := []byte{}

    for _, field := range e.protocol.Fields {
        generatedData := e.generateFieldData(field)
        packet = append(packet, generatedData...)
    }

    protocol.RecalculateComputedFields(packet, e.protocol.Fields)
    log.Printf("Generated packet: %v", packet)
    return packet
}

func (e *Engine) generateFieldData(field protocol.Field) []byte {
    data := make([]byte, field.Length)
    _, err := rand.Read(data)
    if err != nil {
        log.Fatalf("Error generating random data for field %s: %v", field.Name, err)
    }
    return data
}

func (e *Engine) MutatePacket(packet []byte) []byte {
    mutatedPacket := make([]byte, len(packet))
    copy(mutatedPacket, packet)

    // Example mutation: Genetic Algorithm-based mutation
    for i := range mutatedPacket {
        if shouldMutate() {
            mutatedPacket[i] = byte(rand.Intn(256))
        }
    }

    protocol.RecalculateComputedFields(mutatedPacket, e.protocol.Fields)
    log.Printf("Mutated packet: %v", mutatedPacket)
    return mutatedPacket
}

func shouldMutate() bool {
    rand.Seed(time.Now().UnixNano())
    return rand.Float64()
