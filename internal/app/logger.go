package app

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

// CompactTextFormatter formats logs without unnecessary spacing and with color coding
type CompactTextFormatter struct {
	TimestampFormat string
	ForceColors     bool
}

func (f *CompactTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b bytes.Buffer

	// Apply color based on the log level
	levelColor := f.getColor(entry.Level)

	if f.ForceColors {
		b.WriteString(levelColor)
	}

	b.WriteString(fmt.Sprintf("[%s]", strings.ToUpper(entry.Level.String())))

	if entry.HasCaller() {
		b.WriteString(fmt.Sprintf("[%s:%d] ", entry.Caller.File, entry.Caller.Line))
	}

	b.WriteString(entry.Message)

	if len(entry.Data) > 0 {
		b.WriteString(" - ")
		for key, value := range entry.Data {
			b.WriteString(fmt.Sprintf("%s=%v ", key, value))
		}
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *CompactTextFormatter) getColor(level logrus.Level) string {
	switch level {
	case logrus.InfoLevel:
		return "\033[1;34m" // Blue
	case logrus.WarnLevel:
		return "\033[1;33m" // Yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return "\033[1;31m" // Red
	case logrus.DebugLevel:
		return "\033[0;36m" // Cyan
	case logrus.TraceLevel:
		return "\033[0;35m" // Magenta
	default:
		return "\033[0m" // Reset
	}
}

type Logger struct {
	log *logrus.Logger
}

func NewLogger() *Logger {
	log := logrus.New()

	log.SetFormatter(&CompactTextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return &Logger{
		log: log,
	}
}

func (l *Logger) timeEntry() *logrus.Entry {
	return l.log.WithFields(
		logrus.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})

}

func (l *Logger) Info(args ...interface{}) {
	l.timeEntry().Info(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.timeEntry().Infof(format, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.timeEntry().Error(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.timeEntry().Errorf(format, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.timeEntry().Fatal(args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.timeEntry().Fatalf(format, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.timeEntry().Debug(args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.timeEntry().Debugf(format, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.timeEntry().Warn(args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.timeEntry().Warnf(format, args...)
}

func (l *Logger) Trace(args ...interface{}) {
	l.timeEntry().Trace(args...)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	l.timeEntry().Tracef(format, args...)
}

func (l *Logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return l.timeEntry().WithFields(fields)
}
