package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger wraps the zap.Logger for structured logging
type Logger struct {
	sugaredLogger *zap.SugaredLogger
}

// Config defines logger configuration options
type Config struct {
	Level      string // Logging level (e.g., "debug", "info", "warn", "error")
	JSONFormat bool   // Use JSON format for logs
}

// New creates and initializes a new Logger instance
func New(config Config) *Logger {
	var zapConfig zap.Config

	// Set Zap configuration
	if config.JSONFormat {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
	}

	// Set log level
	level := zapcore.InfoLevel
	if err := level.Set(config.Level); err == nil {
		zapConfig.Level = zap.NewAtomicLevelAt(level)
	}

	// Build the logger
	logger, err := zapConfig.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	return &Logger{
		sugaredLogger: logger.Sugar(),
	}
}

// Info logs an info message
func (l *Logger) Info(message string, args ...interface{}) {
	l.sugaredLogger.Infof(message, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(message string, args ...interface{}) {
	l.sugaredLogger.Warnf(message, args...)
}

// Error logs an error message
func (l *Logger) Error(message string, args ...interface{}) {
	l.sugaredLogger.Errorf(message, args...)
}

// Debug logs a debug message
func (l *Logger) Debug(message string, args ...interface{}) {
	l.sugaredLogger.Debugf(message, args...)
}

// Fatal logs a fatal message and exits the application
func (l *Logger) Fatal(message string, args ...interface{}) {
	l.sugaredLogger.Fatalf(message, args...)
}

// Sync flushes any buffered log entries
func (l *Logger) Sync() {
	_ = l.sugaredLogger.Sync()
}
