package logger

import (
	"encoding/json"
	"fmt"
	"github.com/gouef/standards"
	"os"
	"sync"
	"time"
)

type FileLogger struct {
	mu      sync.Mutex
	file    *os.File
	enabled map[standards.LogLevel]bool
}

// NewFileLogger creates a new FileLogger instance.
func NewFileLogger(filePath string, levels ...standards.LogLevel) (*FileLogger, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	enabledLevels := make(map[standards.LogLevel]bool)
	if len(levels) == 0 {
		levels = []standards.LogLevel{
			standards.EMERGENCY,
			standards.CRITICAL,
			standards.ERROR,
			standards.ALERT,
			standards.WARNING,
		}
	}

	for _, level := range levels {
		enabledLevels[level] = true
	}

	return &FileLogger{
		file:    file,
		enabled: enabledLevels,
	}, nil
}

// Log writes a log entry with the specified level, message, and optional context.
func (l *FileLogger) Log(level standards.LogLevel, message string, context []any) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if !l.enabled[level] {
		return nil
	}

	timestamp := time.Now().Format(time.RFC3339)
	logEntry := fmt.Sprintf("[%s] [%s] %s", timestamp, level, message)

	if context != nil {
		contextJSON, err := json.Marshal(context)

		if err != nil {
			return err
		}

		logEntry += " " + string(contextJSON)
	}

	if _, err := l.file.WriteString(logEntry + "\n"); err != nil {
		return fmt.Errorf("failed to write log entry: %w", err)
	}

	return nil
}

// Close closes the underlying log file.
func (l *FileLogger) Close() error {
	return l.file.Close()
}

// Emergency System is unusable
func (l *FileLogger) Emergency(message string, context []any) error {
	return l.Log(standards.EMERGENCY, message, context)
}

// Alert Action must be taken immediately.
func (l *FileLogger) Alert(message string, context []any) error {
	return l.Log(standards.ALERT, message, context)
}

// Critical Critical conditions.
func (l *FileLogger) Critical(message string, context []any) error {
	return l.Log(standards.CRITICAL, message, context)
}

// Error Runtime errors that do not require immediate action but should typically
func (l *FileLogger) Error(message string, context []any) error {
	return l.Log(standards.ERROR, message, context)
}

// Warning Exceptional occurrences that are not errors.
func (l *FileLogger) Warning(message string, context []any) error {
	return l.Log(standards.WARNING, message, context)
}

// Notice Normal but significant events.
func (l *FileLogger) Notice(message string, context []any) error {
	return l.Log(standards.NOTICE, message, context)
}

// Info Interesting events
func (l *FileLogger) Info(message string, context []any) error {
	return l.Log(standards.INFO, message, context)
}

// Debug Detailed debug information.
func (l *FileLogger) Debug(message string, context []any) error {
	return l.Log(standards.DEBUG, message, context)
}
