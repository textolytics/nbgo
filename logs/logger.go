package logs

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"
)

// LogLevel represents the severity level of a log message
type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

// LogEntry represents a single log entry
type LogEntry struct {
	Level      LogLevel
	Timestamp  time.Time
	Message    string
	Module     string
	Fields     map[string]interface{}
	StackTrace string
}

// Logger defines the logging interface
type Logger interface {
	// Debug logs a debug message
	Debug(msg string, fields ...interface{})

	// Info logs an info message
	Info(msg string, fields ...interface{})

	// Warning logs a warning message
	Warning(msg string, fields ...interface{})

	// Error logs an error message
	Error(msg string, fields ...interface{})

	// Fatal logs a fatal message and exits
	Fatal(msg string, fields ...interface{})

	// WithModule returns a logger with a module name
	WithModule(module string) Logger

	// SetLevel sets the minimum log level
	SetLevel(level LogLevel)
}

// StandardLogger implements the Logger interface
type StandardLogger struct {
	mu       sync.RWMutex
	writers  []io.Writer
	level    LogLevel
	module   string
	maxLevel LogLevel
}

// NewStandardLogger creates a new standard logger
func NewStandardLogger(writers ...io.Writer) *StandardLogger {
	return &StandardLogger{
		writers:  writers,
		level:    LevelInfo,
		maxLevel: LevelFatal,
	}
}

// log writes a log entry
func (sl *StandardLogger) log(level LogLevel, msg string, fields ...interface{}) {
	if level < sl.level {
		return
	}

	entry := &LogEntry{
		Level:     level,
		Timestamp: time.Now(),
		Message:   msg,
		Module:    sl.module,
		Fields:    make(map[string]interface{}),
	}

	// Parse fields into a map (simple key-value pairs)
	for i := 0; i < len(fields)-1; i += 2 {
		key := fmt.Sprintf("%v", fields[i])
		value := fields[i+1]
		entry.Fields[key] = value
	}

	// Write to all writers
	sl.mu.RLock()
	defer sl.mu.RUnlock()

	for _, w := range sl.writers {
		fmt.Fprintf(w, "[%s] %s - %s: %s\n",
			entry.Timestamp.Format(time.RFC3339),
			levelToString(level),
			entry.Module,
			entry.Message)
	}
}

// Debug logs a debug message
func (sl *StandardLogger) Debug(msg string, fields ...interface{}) {
	sl.log(LevelDebug, msg, fields...)
}

// Info logs an info message
func (sl *StandardLogger) Info(msg string, fields ...interface{}) {
	sl.log(LevelInfo, msg, fields...)
}

// Warning logs a warning message
func (sl *StandardLogger) Warning(msg string, fields ...interface{}) {
	sl.log(LevelWarning, msg, fields...)
}

// Error logs an error message
func (sl *StandardLogger) Error(msg string, fields ...interface{}) {
	sl.log(LevelError, msg, fields...)
}

// Fatal logs a fatal message
func (sl *StandardLogger) Fatal(msg string, fields ...interface{}) {
	sl.log(LevelFatal, msg, fields...)
}

// WithModule returns a logger with a module name
func (sl *StandardLogger) WithModule(module string) Logger {
	return &StandardLogger{
		writers: sl.writers,
		level:   sl.level,
		module:  module,
	}
}

// SetLevel sets the minimum log level
func (sl *StandardLogger) SetLevel(level LogLevel) {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	sl.level = level
}

// levelToString converts LogLevel to string
func levelToString(level LogLevel) string {
	switch level {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarning:
		return "WARNING"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Router manages log routing
type Router struct {
	mu            sync.RWMutex
	loggers       map[string]Logger
	defaultLogger Logger
}

// NewRouter creates a new log router
func NewRouter(defaultLogger Logger) *Router {
	return &Router{
		loggers:       make(map[string]Logger),
		defaultLogger: defaultLogger,
	}
}

// RegisterLogger registers a logger for a topic
func (r *Router) RegisterLogger(topic string, logger Logger) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.loggers[topic] = logger
}

// GetLogger gets a logger for a topic
func (r *Router) GetLogger(topic string) Logger {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if logger, exists := r.loggers[topic]; exists {
		return logger
	}
	return r.defaultLogger
}

// Route routes a log entry based on topic
func (r *Router) Route(ctx context.Context, topic string, entry *LogEntry) error {
	logger := r.GetLogger(topic)
	if logger == nil {
		return fmt.Errorf("no logger found for topic: %s", topic)
	}
	return nil
}
