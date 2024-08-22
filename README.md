# Network Protocol Fuzzer

## Overview

The Network Protocol Fuzzer is a high-performance distributed network protocol fuzzer implemented in Go. It is designed to discover vulnerabilities in network services by generating and mutating packets based on protocol definitions.

## Features

- **Flexible Protocol Definitions**: Define network protocols using a JSON schema.
- **High-Speed Packet Generation**: Capable of generating millions of packets per second.
- **Stateful Fuzzing Engine**: Supports complex mutation strategies and stateful testing.
- **Distributed Fuzzing**: Uses gRPC for distributed task management across multiple machines.

## Project Structure

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

## Installation

1. **Clone the Repository**

   ```sh
   git clone https://github.com/yourusername/network-protocol-fuzzer.git
   cd network-protocol-fuzzer

2. **Build the Application**

    ```sh
    go build -o fuzzer ./cmd/fuzzer
    ```

## Usage

1. **Run the Fuzzer**

    To run the fuzzer, use the following command:

    ```sh
    ./fuzzer -protocol path/to/protocol.json
    ```

    Replace path/to/protocol.json with the path to your protocol definition file.

2. **Run the Distributed System**

    - Start the Master Node

    ```sh
    go run ./cmd/distributed/master.go --port :50051
    ```

    - Start the Worked Nodes

    ```sh
    go run ./cmd/distributed/worker.go --master-address localhost:50051 --worker-id 1
    ```

## Testing

    Run the tests using the following command:

    ```sh
    go test ./...
    ```

## Docker

    To build and run the application using Docker, use the following commands:

    **1. Build the Docker Image**

    ```sh
    docker build -t network-protocol-fuzzer .
    ```

    **2. Run the Docker Container**

    ```sh
    docker run -it --rm network-protocol-fuzzer
    ```