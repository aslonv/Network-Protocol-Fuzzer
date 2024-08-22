# Network Protocol Fuzzer

This project is a high-performance network protocol fuzzer written in Go. It is designed to discover vulnerabilities in network services by generating and mutating network packets based on protocol definitions.

## Features

- **Protocol Definition**: Flexible JSON-based protocol definitions.
- **Packet Generation**: High-speed packet generation engine.
- **Mutation Strategies**: Genetic algorithm-based mutation for comprehensive fuzzing.
- **Distributed Fuzzing**: Leverage multiple machines for increased fuzzing throughput.

## Getting Started

1. **Clone the repository**:
    ```sh
    git clone https://github.com/yourusername/network-protocol-fuzzer.git
    cd network-protocol-fuzzer
    ```

2. **Build and run the fuzzer**:
    ```sh
    make build
    ./bin/fuzzer --protocol=protocol.json
    ```

3. **Run tests**:
    ```sh
    make test
    ```

## Docker

To build and run the fuzzer using Docker:

```sh
make docker
docker run --rm network-protocol-fuzzer
