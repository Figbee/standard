package initialize

import (
	"fmt"
	"os"
	"standard/internal/global"
	"testing"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestInitLogger(t *testing.T) {
	now := time.Now()
	path := global.Conf.Logs.Path
	filename := fmt.Sprintf("%s/%04d-%02d-%02d.log", path, now.Year(), now.Month(), now.Day())
	hook := lumberjack.Logger{
		Filename:   filename,
		MaxSize:    50,
		MaxAge:     30,
		MaxBackups: 10,
		Compress:   false,
	}
	enConfig := zap.NewProductionEncoderConfig()
	enConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(enConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		zap.DebugLevel,
	)

	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	global.Logger = log.Sugar()
	global.Logger.Debug("初始化日志完成")
}
