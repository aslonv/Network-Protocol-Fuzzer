# Network Protocol Fuzzer

The Network Protocol Fuzzer is a high-performance distributed network protocol fuzzer implemented in Go. It is designed to discover vulnerabilities in network services by generating and mutating packets based on protocol definitions.

- **Flexible Protocol Definitions**: Define network protocols using a JSON schema.
- **High-Speed Packet Generation**: Capable of generating millions of packets per second.
- **Stateful Fuzzing Engine**: Supports complex mutation strategies and stateful testing.
- **Distributed Fuzzing**: Uses gRPC for distributed task management across multiple machines.

## Structure

- **cmd/**: Contains the entry points for the application.
  - `fuzzer/`: Main application for running the fuzzer.
- **pkg/**: Contains reusable packages.
  - `protocol/`: Protocol definition and handling.
  - `engine/`: Packet generation and mutation.
  - `fuzzer/`: Main fuzzing logic.
- **internal/**: Contains internal packages.
  - `logger/`: Logging and metrics.
  - `distributed/`: gRPC server and client for distributed fuzzing.
- **docs/**: Documentation and protocol examples.
- **Dockerfile**: Docker configuration for building and running the application.
- **.github/workflows/ci.yml**: CI configuration for GitHub Actions.
