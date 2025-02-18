package utils

import "go.uber.org/zap"

func Getlogger() *zap.Logger {

	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout"}
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	return logger
}
