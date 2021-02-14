package logger

import (
	"go.uber.org/zap"
)

var (
	defaultLogger, _ = zap.NewProduction()
)

// Log returns zap.Logger instance
func Log() *zap.Logger {
	return defaultLogger
}
