package logger

import "github.com/gouef/standards"

type MultiLogger struct {
	loggers []standards.Logger
}

func NewMultiLogger(loggers ...standards.Logger) *MultiLogger {
	return &MultiLogger{loggers: loggers}
}

// Log Logs with an arbitrary level.
func (l *MultiLogger) Log(level standards.LogLevel, message string, context []any) error {
	for _, logger := range l.loggers {
		if err := logger.Log(level, message, context); err != nil {
			return err
		}
	}

	return nil
}

// Emergency System is unusable
func (l *MultiLogger) Emergency(message string, context []any) error {
	return l.Log(standards.EMERGENCY, message, context)
}

// Alert Action must be taken immediately.
func (l *MultiLogger) Alert(message string, context []any) error {
	return l.Log(standards.ALERT, message, context)
}

// Critical Critical conditions.
func (l *MultiLogger) Critical(message string, context []any) error {
	return l.Log(standards.CRITICAL, message, context)
}

// Error Runtime errors that do not require immediate action but should typically
func (l *MultiLogger) Error(message string, context []any) error {
	return l.Log(standards.ERROR, message, context)
}

// Warning Exceptional occurrences that are not errors.
func (l *MultiLogger) Warning(message string, context []any) error {
	return l.Log(standards.WARNING, message, context)
}

// Notice Normal but significant events.
func (l *MultiLogger) Notice(message string, context []any) error {
	return l.Log(standards.NOTICE, message, context)
}

// Info Interesting events
func (l *MultiLogger) Info(message string, context []any) error {
	return l.Log(standards.INFO, message, context)
}

// Debug Detailed debug information.
func (l *MultiLogger) Debug(message string, context []any) error {
	return l.Log(standards.DEBUG, message, context)
}
