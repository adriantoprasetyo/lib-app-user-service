// logger/logger.go
package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func Init() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	Log = logger.Sugar()
}

func Sync() {
	_ = Log.Sync()
}
