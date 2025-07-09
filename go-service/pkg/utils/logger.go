package utils

import "github.com/sirupsen/logrus"

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Trace(args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	SetLogLevel(level logrus.Level)
	WithError(err error) *logrus.Entry
	WithFields(fields logrus.Fields) *logrus.Entry
}

type LogrusLogger struct {
	*logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{logrus.New()}
}

func (l *LogrusLogger) SetLogLevel(level logrus.Level) {
	l.SetLevel(level)
}

func (l *LogrusLogger) GetLogger() *logrus.Logger {
	return l.Logger
}

func (l *LogrusLogger) WithError(err error) *logrus.Entry {
	return l.Logger.WithError(err)
}

func (l *LogrusLogger) WithFields(fields logrus.Fields) *logrus.Entry {
	return l.Logger.WithFields(fields)
}

func ParseLevel(l string) logrus.Level {
	level, err := logrus.ParseLevel(l)
	if err != nil {
		return 6
	}
	return level
}
