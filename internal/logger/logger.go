package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	logger *logrus.Logger
}

func New() *Logger {
	return &Logger{
		logger: logrus.New(),
	}
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}
