.PHONY: build run clean docker

build:
    go build -o bin/fuzzer ./cmd/fuzzer

run: build
    ./bin/fuzzer --protocol=protocol.json

test:
    go test ./...

clean:
    go clean
    rm -rf bin/

docker:
    docker build -t network-protocol-fuzzer .
