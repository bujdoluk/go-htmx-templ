package logger

import (
	"encoding/json"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

// Level represents severity of a log message.
type Level int8

const (
	// LevelDebug represents detailed information messages.
	LevelDebug Level = iota
	// LevelInfo represents general informational messages.
	LevelInfo
	// LevelWarning indicates a potential issue or unexpected behavior.
	LevelWarning
	// LevelError indicates an error that occurred during execution.
	LevelError
	// LevelFatal indicates a critical error after which the application will terminate.
	LevelFatal
	// LevelOff disables all logging.
	LevelOff
)

func (l Level) String() string {
	switch l {
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
		return ""
	}
}

// Logger
type Logger struct {
	out      io.Writer
	minLevel Level
	mu       sync.Mutex
}

// New represents new instance of logger
func New(out io.Writer, minLevel Level) *Logger {
	return &Logger{
		out:      out,
		minLevel: minLevel,
	}
}

// Writes message with level error to print method
func (l *Logger) Write(message []byte) (n int, err error) {
	return l.print(LevelError, string(message), nil)
}

// Prints debug message with level
func (l *Logger) PrintDebug(message string, properties map[string]string) {
	l.print(LevelDebug, message, properties)
}

// Prints info message with level
func (l *Logger) PrintInfo(message string, properties map[string]string) {
	l.print(LevelInfo, message, properties)
}

// Prints warning message with level
func (l *Logger) PrintWarning(message string, properties map[string]string) {
	l.print(LevelWarning, message, properties)
}

// Prints error with level
func (l *Logger) PrintError(err error, properties map[string]string) {
	l.print(LevelError, err.Error(), properties)
}

// Prints error with level and exits an app
func (l *Logger) PrintFatal(err error, properties map[string]string) {
	l.print(LevelFatal, err.Error(), properties)
	os.Exit(1)
}

func (l *Logger) print(level Level, message string, properties map[string]string) (int, error) {
	if level < l.minLevel {
		return 0, nil
	}

	aux := struct {
		Level      string            `json:"level"`
		Time       string            `json:"time"`
		Message    string            `json:"message"`
		Properties map[string]string `json:"properties,omitempty"`
		Trace      string            `json:"trace,omitempty"`
	}{
		Level:      level.String(),
		Time:       time.Now().UTC().Format(time.RFC3339),
		Message:    message,
		Properties: properties,
	}

	if level >= LevelError {
		aux.Trace = string(debug.Stack())
	}

	var line []byte

	line, err := json.Marshal(aux)
	if err != nil {
		line = []byte(LevelError.String() + ": unable to marshal log message" + err.Error())
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	return l.out.Write(append(line, '\n'))
}
