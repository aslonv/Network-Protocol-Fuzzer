package protocol

import (
    "testing"
    "os"
    "encoding/json"
)

func TestLoadProtocolDefinition(t *testing.T) {
    // Create a temporary file with a protocol definition
    tempFile, err := os.CreateTemp("", "protocol.json")
    if err != nil {
        t.Fatalf("Failed to create temp file: %v", err)
    }
    defer os.Remove(tempFile.Name())

    proto := Protocol{
        Name:    "TestProtocol",
        Version: "1.0",
        Fields: []Field{
            {Name: "Header", Type: "fixed", Length: 4, Mutation: "none"},
            {Name: "Payload", Type: "variable", Length: 1024, Mutation: "random"},
        },
    }

    data, _ := json.Marshal(proto)
    tempFile.Write(data)
    tempFile.Close()

    loadedProto, err := LoadProtocolDefinition(tempFile.Name())
    if err != nil {
        t.Fatalf("Failed to load protocol definition: %v", err)
    }

    if loadedProto.Name != proto.Name {
        t.Errorf("Expected protocol name %s, got %s", proto.Name, loadedProto.Name)
    }
    if loadedProto.Version != proto.Version {
        t.Errorf("Expected protocol version %s, got %s", proto.Version, loadedProto.Version)
    }
}

func TestCalculateChecksum(t *testing.T) {
    data := []byte("test data")
    checksum := CalculateChecksum(data)
    if len(checksum) != 2 {
        t.Errorf("Checksum length should be 2 bytes, got %d", len(checksum))
    }
}
