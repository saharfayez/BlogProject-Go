// middlewares/logger.go
package middlewares

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	ZapLogger   *zap.Logger
	atomicLevel zap.AtomicLevel
)

func InitLogger(level zapcore.Level) {

	atomicLevel = zap.NewAtomicLevelAt(level)
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.TimeKey = "timestamp"
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapConfig.EncoderConfig),
		zapcore.AddSync(os.Stdout),
		atomicLevel,
	)

	ZapLogger = zap.New(core, zap.AddCaller())
}

//func SetLevel(level zapcore.Level) {
//	atomicLevel.SetLevel(level)
//}
