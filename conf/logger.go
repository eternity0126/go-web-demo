package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.DebugLevel
	if viper.GetString("env") == "dev" {
		logMode = zapcore.DebugLevel
	} else if viper.GetString("env") == "release" {
		logMode = zapcore.InfoLevel
	} else {
		panic("unknown `env` value, choose from 'dev' or 'release'!")
	}
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(
		getWriteSyncer(),
		zapcore.AddSync(os.Stdout),
	), logMode)

	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSyncer() zapcore.WriteSyncer {
	// 记录到log文件夹下
	stSeparator := string(filepath.Separator)
	stRootDir, _ := os.Getwd()
	stLogFilePath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format(time.DateOnly) + ".txt"
	fmt.Println(stLogFilePath)

	// 日志分割器
	lumberjackSyncer := &lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxSize:    viper.GetInt("log.max_size"),
		MaxBackups: viper.GetInt("log.max_backups"),
		MaxAge:     viper.GetInt("log.max_age"),
		Compress:   false,
	}

	return zapcore.AddSync(lumberjackSyncer)

}
