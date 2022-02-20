package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger(l *logrus.Logger) *Logger {
	return &Logger{
		logger: l,
	}
}

var _ LoggerI = Logger{}

func (l Logger) Debug(msg ...interface{}) {
	l.logger.Debug(msg...)
}

func (l Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l Logger) Info(msg ...interface{}) {
	l.logger.Info(msg...)
}

func (l Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l Logger) Warn(msg ...interface{}) {
	l.logger.Warn(msg...)
}

func (l Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l Logger) Error(msg ...interface{}) {
	l.logger.Error(msg...)
}

func (l Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l Logger) Fatal(msg ...interface{}) {
	l.logger.Fatal(msg...)
}

func (l Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}
