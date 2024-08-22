package logger

import (
    "os"
    "testing"
)

func TestInitializeLogger(t *testing.T) {
    logFile := "test.log"
    Initialize(logFile, true)
    
    // Check if log file was created
    if _, err := os.Stat(logFile); os.IsNotExist(err) {
        t.Errorf("Log file %s was not created", logFile)
    }
    os.Remove(logFile)
}
