package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

const (
	LevelTrace = slog.Level(-8) // More verbose than DEBUG
	LevelFatal = slog.Level(12) // More severe than ERROR
)

func main() {

	opts := &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				switch level {
				case LevelTrace:
					a.Value = slog.StringValue("TRACE")
				case LevelFatal:
					a.Value = slog.StringValue("FATAL")
				}
			}
			return a
		},
	}
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelDebug,
		ReplaceAttr: opts.ReplaceAttr,
	})
	logger := slog.New(handler)

	logger.Log(context.Background(), slog.LevelInfo, "an info message")
	logger.LogAttrs(context.Background(), slog.LevelInfo, "an info message")
	logger.Info("incoming request", "method", "GET", "status", "200")
	logger.Warn("Permission denied",
		slog.Int("user_id", 12345),
		slog.String("resource", "/api/admin"))

	if logger.Enabled(context.Background(), slog.LevelDebug) {
		logger.Debug("operation complete", "data", getExpensiveDebugData())
	}
	logger.Log(context.Background(), LevelFatal, "database connection lost")

}

func getExpensiveDebugData() error {
	return fmt.Errorf("Expensive Debug Data")
}
