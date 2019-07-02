package zap

import (
	"github.com/antizhou/goutils/file"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	logger *zap.SugaredLogger
)

func init() {
	errorPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	infoPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel || lvl == zapcore.WarnLevel
	})
	debugPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel
	})

	filePath := "./error.log"
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	if file.IsExist(filePath) {
		os.Create(filePath)
	}

	errorWriteSyncer := zapcore.AddSync(f)
	infoWriteSyncer := zapcore.Lock(os.Stdout)
	debugWriteSyncer := zapcore.Lock(os.Stdout)

	errorEncoderConfig := zap.NewDevelopmentEncoderConfig()
	infoEncoderConfig := zap.NewDevelopmentEncoderConfig()
	debugEncoderConfig := zap.NewDevelopmentEncoderConfig()

	errorEncoder := zapcore.NewConsoleEncoder(errorEncoderConfig)
	infoEncoder := zapcore.NewConsoleEncoder(infoEncoderConfig)
	debugEncoder := zapcore.NewConsoleEncoder(debugEncoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(errorEncoder, errorWriteSyncer, errorPriority),
		zapcore.NewCore(infoEncoder, infoWriteSyncer, infoPriority),
		zapcore.NewCore(debugEncoder, debugWriteSyncer, debugPriority),
	)

	logger = zap.New(core).Sugar()
	//defer logger.Sync()

	logger.Info("constructed a logger")

}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

func Warn(v ...interface{}) {
	logger.Warn(v...)
}

func WarnF(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

func Error(v ...interface{}) {
	logger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}
