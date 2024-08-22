package protocol

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "log"
)

type Protocol struct {
    Name    string  `json:"name"`
    Version string  `json:"version"`
    Fields  []Field `json:"fields"`
}

type Field struct {
    Name     string `json:"name"`
    Type     string `json:"type"`
    Length   int    `json:"length"`
    Mutation string `json:"mutation"`
}

func LoadProtocolDefinition(filepath string) (*Protocol, error) {
    data, err := ioutil.ReadFile(filepath)
    if err != nil {
        return nil, fmt.Errorf("failed to read protocol definition: %w", err)
    }

    var protocol Protocol
    if err := json.Unmarshal(data, &protocol); err != nil {
        return nil, fmt.Errorf("failed to unmarshal protocol definition: %w", err)
    }

    if err := validateProtocol(&protocol); err != nil {
        return nil, fmt.Errorf("protocol validation failed: %w", err)
    }

    log.Printf("Loaded protocol: %s (version: %s)", protocol.Name, protocol.Version)
    return &protocol, nil
}

func validateProtocol(p *Protocol) error {
    if p.Name == "" || p.Version == "" {
        return errors.New("protocol name and version must be specified")
    }
    if len(p.Fields) == 0 {
        return errors.New("protocol must have at least one field")
    }
    for _, field := range p.Fields {
        if field.Name == "" || field.Type == "" {
            return errors.New("field name and type must be specified")
        }
        if field.Length <= 0 {
            return errors.New("field length must be greater than zero")
        }
    }
    return nil
}
