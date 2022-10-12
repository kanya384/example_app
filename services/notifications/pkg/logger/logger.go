package logger

import (
	"fmt"
	"os"
	"strings"

	formatters "github.com/fabienm/go-logrus-formatters"
	"github.com/sirupsen/logrus"

	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
)

// Interface -.
type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

// Logger -.
type Logger struct {
	logger *logrus.Logger
}

var _ Interface = (*Logger)(nil)

// New -.
func New(level string, serviceName string, graylogHost string) *Logger {
	var l logrus.Level

	switch strings.ToLower(level) {
	case "error":
		l = logrus.ErrorLevel
	case "warn":
		l = logrus.WarnLevel
	case "info":
		l = logrus.InfoLevel
	case "debug":
		l = logrus.DebugLevel
	default:
		l = logrus.InfoLevel
	}

	fmtr := formatters.NewGelf(serviceName)
	hooks := []logrus.Hook{graylog.NewGraylogHook(graylogHost, map[string]interface{}{})}

	logger := logrus.New()
	logger.Formatter = fmtr
	for _, hook := range hooks {
		logger.AddHook(hook)
	}
	logger.Level = l

	return &Logger{
		logger: logger,
	}
}

// Debug -.
func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.msg("debug", message, args...)
}

// Info -.
func (l *Logger) Info(message string, args ...interface{}) {
	l.log(message, args...)
}

// Warn -.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.log(message, args...)
}

// Error -.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	if l.logger.GetLevel() == logrus.DebugLevel {
		l.Debug(message, args...)
	}

	l.msg("error", message, args...)
}

// Fatal -.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.msg("fatal", message, args...)

	os.Exit(1)
}

func (l *Logger) log(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Info(message)
	} else {
		l.logger.Infof(message, args...)
	}
}

func (l *Logger) msg(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(msg.Error(), args...)
	case string:
		l.log(msg, args...)
	default:
		l.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
