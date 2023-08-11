package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger struct {
	zap.Logger
}

func NewLogger() *Logger {
	logger := &Logger{}
	l, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("logger failed to initialize. %s", err))
	}
	defer l.Sync()

	logger.Logger = *l

	return logger
}

func NewTestLogger() *Logger {
	logger := Logger{}
	logger.Logger = *zap.NewNop()

	return &logger
}
