package main

import (
    "flag"
    "log"

    "network-protocol-fuzzer/pkg/fuzzer"
    "network-protocol-fuzzer/internal/logger"
)

func main() {
    protocolFile := flag.String("protocol", "protocol.json", "Path to the protocol definition file")
    logFile := flag.String("log", "fuzzer.log", "Path to the log file")
    verbose := flag.Bool("v", false, "Enable verbose logging")

    flag.Parse()

    logger.Initialize(*logFile, *verbose)

    fuzzerInstance, err := fuzzer.NewFuzzer(*protocolFile)
    if err != nil {
        log.Fatalf("Error initializing fuzzer: %v", err)
    }

    if err := fuzzerInstance.Run(); err != nil {
        log.Fatalf("Fuzzer encountered an error: %v", err)
    }
}
