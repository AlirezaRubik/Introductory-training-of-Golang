package main

import (
	"os"
	"log/slog"
	"github.com/go-stack/stack"
)

func main() {
	//default log system working with strings but slog using json
	//we have different os std : stdout,stderr,stdin 
	//std output = terminal
	// Create a new logger with JSON output
	logger := slog.New(slog.NewJSONHandler(os.Stderr,&slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Log a simple info message
	logger.Info("User login successful", "module", "auth", "request_id", 12345)

	// Log a warning message
	logger.Warn("Invalid login attempt", "module", "auth", "request_id", 12345)

	// Log an error message with a stack trace
	logger.Error("Critical error occurred", "stack", stack.Trace().String(), "request_id", 12345)

	// Log an info message with additional fields
	logger.Info("Data processing started", "module", "data_pipeline", "request_id", 12345, "processing_time", "500ms")
}
