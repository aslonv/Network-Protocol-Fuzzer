package engine

import (
    "testing"
    "network-protocol-fuzzer/pkg/protocol"
)

func TestGeneratePacket(t *testing.T) {
    proto := &protocol.Protocol{
        Name:    "TestProtocol",
        Version: "1.0",
        Fields: []protocol.Field{
            {Name: "Header", Type: "fixed", Length: 4, Mutation: "none"},
            {Name: "Payload", Type: "variable", Length: 1024, Mutation: "random"},
        },
    }
    engine := NewEngine(proto)
    packet := engine.GeneratePacket()

    if len(packet) != 1028 { // 4 bytes header + 1024 bytes payload
        t.Errorf("Generated packet length should be 1028 bytes, got %d", len(packet))
    }
}

func TestMutatePacket(t *testing.T) {
    proto := &protocol.Protocol{
        Name:    "TestProtocol",
        Version: "1.0",
        Fields: []protocol.Field{
            {Name: "Header", Type: "fixed", Length: 4, Mutation: "none"},
            {Name: "Payload", Type: "variable", Length: 1024, Mutation: "random"},
        },
    }
    engine := NewEngine(proto)
    packet := engine.GeneratePacket()
    mutatedPacket := engine.MutatePacket(packet)

    if len(mutatedPacket) != len(packet) {
        t.Errorf("Mutated packet length should be same as original packet length")
    }
}
