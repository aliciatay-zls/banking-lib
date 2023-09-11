package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Wrap logger
var appLogger *zap.Logger //declare
var originalAppLogger *zap.Logger

func init() {
	config := zap.NewProductionConfig()

	customConfig := zap.NewProductionEncoderConfig()
	customConfig.TimeKey = "timestamp"
	customConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	customConfig.StacktraceKey = ""

	config.EncoderConfig = customConfig

	var err error
	appLogger, err = config.Build(zap.AddCallerSkip(1)) //initialize
	if err != nil {
		panic(err)
	}

	originalAppLogger = appLogger //save previous logger settings
}

// Expose its functionalities
func Info(message string, fields ...zap.Field) { //[]zap.Field
	appLogger.Info(message, fields...) //unpack slice
}

func Debug(message string, fields ...zap.Field) {
	appLogger.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	appLogger.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	appLogger.Fatal(message, fields...)
}

func MuteLogger() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{os.DevNull}

	var err error
	appLogger, err = config.Build()
	if err != nil {
		panic(err)
	}
}

func UnmuteLogger() {
	appLogger = originalAppLogger //restore previous logger settings
}
