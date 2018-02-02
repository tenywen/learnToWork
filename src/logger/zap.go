package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

import (
	"cfg"
)

var sugarLogger *zap.SugaredLogger
var GOPATH string

func StartLogger(fileName string) {
	GOPATH = os.Getenv("GOPATH")
	if GOPATH == "" {
		fmt.Println("GOPATH not set", GOPATH)
		os.Exit(-1)
	}

	dir := GOPATH + "/" + cfg.Config.LOG.Path
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("mkdir ", err)
		os.Exit(-1)
	}
	name := dir + fileName
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   name,
		MaxSize:    cfg.Config.LOG.MaxSize,
		MaxBackups: cfg.Config.LOG.MaxBackups,
		MaxAge:     cfg.Config.LOG.MaxAge,
		Compress:   cfg.Config.LOG.Compress,
		LocalTime:  cfg.Config.LOG.LocalTime,
	})

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		zapcore.Level(cfg.Config.LOG.Level),
	)

	logger := zap.New(core)
	sugarLogger = logger.Sugar()
}

func DEBUG(v ...interface{}) {
	sugarLogger.Debug(v...)
}

func DEBUGF(format string, v ...interface{}) {
	sugarLogger.Debugf(format, v...)
}

func WARN(v ...interface{}) {
	sugarLogger.Warn(v...)
}

func WARNF(format string, v ...interface{}) {
	sugarLogger.Warnf(format, v...)
}

func ERROR(v ...interface{}) {
	sugarLogger.Error(v...)
}

func ERRORF(format string, v ...interface{}) {
	sugarLogger.Errorf(format, v...)
}

func INFO(v ...interface{}) {
	sugarLogger.Info(v...)
}

func INFOF(format string, v ...interface{}) {
	sugarLogger.Infof(format, v...)
}
