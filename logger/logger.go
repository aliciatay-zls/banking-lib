package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
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

// ReplaceWithTestLogger overrides init() that was called when this method was called with an observer logger,
// then returns a collection of the observed logs that can be used directly in tests.
func ReplaceWithTestLogger() *observer.ObservedLogs {
	observerCore, observedLogs := observer.New(zap.InfoLevel)
	observerLogger := zap.New(observerCore)
	appLogger = observerLogger

	return observedLogs
}

//Notes on ReplaceWithTestLogger()
// observer.New(zap.InfoLevel)
// Param LevelEnabler: decides whether a given logging level is enabled when logging a message.
// Each concrete Level value implements a static LevelEnabler which returns true for itself and all higher logging levels.
// = hence here info level and above messages (i.e. all levels) will be logged (messages of Info(), Error(), etc calls will be captured)
//
// Reference: https://medium.com/go-for-punks/handle-zap-log-messages-in-a-test-8503b25fe38f
