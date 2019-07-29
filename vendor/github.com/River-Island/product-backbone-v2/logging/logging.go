package logging

import (
	"runtime"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LogLevel represents the logging level.
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
	PANIC
	DISABLED
)

// LogKey is the type each key that appears in the log should be.
type LogKey string

// String returns the string representation of the LogKey
func (lk LogKey) String() string {
	return string(lk)
}

// Each LogKey appearing in the logs is defined in the following const block.
const (
	CallerKey  LogKey = "caller"
	UUIDKey    LogKey = "uuid"
	ServiceKey LogKey = "service"
)

var (
	NopLogger = &Logger{Zero: zerolog.Nop()}
)

// Logger implements the LogChainer interface and relies on http://github.com/rs/zerolog
type Logger struct {
	Zero zerolog.Logger
}

// LogChainer defines the behavior of the logger.
// Exposes a function for each loggable field which maps to a LogKey.
// The functions invocations can be chained and terminated by one of the levelled function calls (Fatal, Error, Warn, Info).
type LogChainer interface {
	GetOption(s string) LogChainer

	// These are the last functions that should be called on a log chain.
	// These will execute and log all the information
	Panic(msg string)
	Fatal(msg string, err error)
	Error(msg string, err error)
	Warn(msg string)
	Info(msg string)
	Debug(msg string)
}

// CustomField - instructs the logger to log a custom field of any type.
// It should be used for logging some specific details and additional information that won't be used for querying.
func (l *Logger) GetOption(s string) LogChainer {
	lcopy := *l
	lcopy.Zero = l.Zero.With().Str(UUIDKey.String(), s).Logger()
	return &lcopy
}

// Panic logs the message at panic level.
// It stops the ordinary flow of a goroutine.
// The log payload will contain everything else the logger has been instructed to log.
func (l *Logger) Panic(msg string) {
	l.Zero.Panic().Str(CallerKey.String(), getCallerFunctionName()).Msg(msg)
}

// Fatal logs the message and the error at fatal level.
// It after exits with os.Exit(1).
// The log payload will contain everything else the logger has been instructed to log.
func (l *Logger) Fatal(msg string, err error) {
	l.Zero.Fatal().Str(CallerKey.String(), getCallerFunctionName()).AnErr("error", err).Msg(msg)
}

// Error logs the message and the error at error level.
// The log payload will contain everything else the logger has been instructed to log.
func (l *Logger) Error(msg string, err error) {
	l.Zero.Error().Str(CallerKey.String(), getCallerFunctionName()).AnErr("error", err).Msg(msg)
}

// Warn logs the message at warning level.
// The log payload will contain everything else the logger has been instructed to log.
func (l *Logger) Warn(msg string) {
	l.Zero.Warn().Str(CallerKey.String(), getCallerFunctionName()).Msg(msg)
}

// Info logs the message at info level.
// The log payload will contain everything else the logger has been instructed to log.
func (l *Logger) Info(msg string) {
	l.Zero.Info().Str(CallerKey.String(), getCallerFunctionName()).Msg(msg)
}

// Debug logs the message at debug level.
// The log payload will contain everything else the logger has been instructed to log.
func (l *Logger) Debug(msg string) {
	l.Zero.Debug().Str(CallerKey.String(), getCallerFunctionName()).Msg(msg)
}

// getCallerFunctionName returns the function caller's name.
func getCallerFunctionName() string {
	// Skip GetCallerFunctionName and the function to get the caller of
	return getFrame(2).Function
}

func getFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"} // fallback caller name
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}

// GetLogger returns a pointer to a Logger that logs from logLevel and above.
// The logger is instructed to include in each log message the name of the service received in input.
func GetLogger(service string, logLevel LogLevel) *Logger {
	switch logLevel {
	case DEBUG:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case INFO:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case WARNING:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case ERROR:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case FATAL:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case PANIC:
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case DISABLED:
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}
	return &Logger{
		Zero: log.With().Str(ServiceKey.String(), service).Logger(),
	}
}

// GetLoggerString - alternative Logger constructor that returns a pointer to a Logger based on a string defining
// a log level.
// The default value is INFO.
func GetLoggerString(service string, logLevel string) *Logger {
	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "WARNING":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "FATAL":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "PANIC":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "DISABLED":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	return &Logger{
		Zero: log.With().Str(ServiceKey.String(), service).Logger(),
	}
}
