package logger

import (
	"go.uber.org/zap"
)

// New is ...
func New() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()
	defer logger.Sync()
	return slogger
}
