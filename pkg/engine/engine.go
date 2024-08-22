package engine

import (
    "crypto/rand"
    "log"
    "network-protocol-fuzzer/pkg/protocol"
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

    // Example mutation: Flip random bits
    for i := range mutatedPacket {
        if rand.Intn(2) == 0 { // 50% chance to flip the bit
            mutatedPacket[i] ^= 1 << uint(rand.Intn(8))
        }
    }

    log.Printf("Mutated packet: %v", mutatedPacket)
    return mutatedPacket
}
