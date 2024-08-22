package logger

import (
    "log"
    "os"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
)

var (
    PacketsProcessed = promauto.NewCounter(prometheus.CounterOpts{
        Name: "fuzzer_packets_processed_total",
        Help: "The total number of processed packets",
    })

    PacketGenerationErrors = promauto.NewCounter(prometheus.CounterOpts{
        Name: "fuzzer_packet_generation_errors_total",
        Help: "The total number of packet generation errors",
    })
)

func Initialize(logFile string, verbose bool) {
    file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }

    log.SetOutput(file)
    if verbose {
        log.SetFlags(log.LstdFlags | log.Lshortfile)
    } else {
        log.SetFlags(log.LstdFlags)
    }

    go func() {
        http.Handle("/metrics", promhttp.Handler())
        log.Fatal(http.ListenAndServe(":2112", nil))
    }()
}
