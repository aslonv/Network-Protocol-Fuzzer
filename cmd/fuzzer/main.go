// Entry point for the network protocol fuzzer.
package main

import (
	"flag"
	"fmt"
	"os"

	"network-protocol-fuzzer/internal/logger"
	"network-protocol-fuzzer/pkg/fuzzer"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// Executes the main application logic and handles errors.
func run() error {
	protocolFile := flag.String("protocol", "protocol.json", "protocol definition file path")
	logFile := flag.String("log", "fuzzer.log", "log file path")
	verbose := flag.Bool("v", false, "enable verbose logging")
	flag.Parse()

	if err := logger.Initialize(*logFile, *verbose); err != nil {
		return fmt.Errorf("logger initialization: %w", err)
	}

	fuzzerInstance, err := fuzzer.NewFuzzer(*protocolFile)
	if err != nil {
		return fmt.Errorf("fuzzer initialization: %w", err)
	}

	if err := fuzzerInstance.Run(); err != nil {
		return fmt.Errorf("fuzzer execution: %w", err)
	}

	return nil
}
