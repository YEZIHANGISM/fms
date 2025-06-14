package log

import (
	"fms/configs"

	"go.uber.org/zap"
)

var (
	Logger  *zap.Logger
	SLogger *zap.SugaredLogger
)

func InitLogger() {
	var err error
	if configs.Config.Basic.Debug {
		Logger, err = zap.NewDevelopment()
	} else {
		Logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(err)
	}
	SLogger = Logger.Sugar()
}
