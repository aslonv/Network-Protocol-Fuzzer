// Entry point for the network protocol fuzzer.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"network-protocol-fuzzer/internal/logger"
	"network-protocol-fuzzer/pkg/fuzzer"
	"network-protocol-fuzzer/pkg/metrics"
	"network-protocol-fuzzer/pkg/config"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := parseFlags()
	if err != nil {
		return fmt.Errorf("parsing flags: %w", err)
	}

	if err := logger.Initialize(cfg.LogFile, cfg.Verbose); err != nil {
		return fmt.Errorf("logger initialization: %w", err)
	}
	defer logger.Cleanup()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	setupGracefulShutdown(cancel)

	metricsCollector := metrics.NewCollector()
	fuzzerInstance, err := fuzzer.NewFuzzer(cfg.ProtocolFile, fuzzer.WithMetrics(metricsCollector))
	if err != nil {
		return fmt.Errorf("fuzzer initialization: %w", err)
	}

	if err := fuzzerInstance.Run(ctx); err != nil {
		return fmt.Errorf("fuzzer execution: %w", err)
	}

	metricsCollector.DisplaySummary()

	return nil
}

func parseFlags() (*config.Config, error) {
	cfg := &config.Config{}
	flag.StringVar(&cfg.ProtocolFile, "protocol", "protocol.json", "protocol definition file path")
	flag.StringVar(&cfg.LogFile, "log", "fuzzer.log", "log file path")
	flag.BoolVar(&cfg.Verbose, "v", false, "enable verbose logging")
	flag.IntVar(&cfg.Threads, "threads", runtime.NumCPU(), "number of fuzzing threads")
	flag.DurationVar(&cfg.Timeout, "timeout", 5*time.Minute, "fuzzing timeout")
	flag.Parse()

	return cfg, cfg.Validate()
}

func setupGracefulShutdown(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nReceived interrupt signal. Shutting down gracefully...")
		cancel()
	}()
}
