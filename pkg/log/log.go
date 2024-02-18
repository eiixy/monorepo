package log

import (
	"fmt"
	"github.com/eiixy/monorepo/internal/pkg/config"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"time"
)

func NewLoggerFromConfig(conf config.Log, name string, keyvals ...interface{}) log.Logger {
	options := &ZapOptions{
		Filename:     path.Join(conf.Dir, name),
		Level:        Level(conf.Level),
		MaxAge:       time.Duration(conf.MaxAge) * time.Hour * 24,
		RotationTime: time.Duration(conf.RotationTime) * time.Hour * 24,
	}

	keyvals = append(keyvals, "ts", log.DefaultTimestamp)
	keyvals = append(keyvals, "caller", log.DefaultCaller)
	return log.With(NewLogger(options), keyvals...)
}

func NewLogger(options *ZapOptions) log.Logger {
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		LevelKey:    "level",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
	})

	level := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= options.Level
	})

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(writer(options)), level),
	)

	return &zapLogger{
		logger: zap.New(core),
	}
}

type zapLogger struct {
	logger *zap.Logger
}

func (l *zapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.logger.Warn(fmt.Sprint("keyvals must appear in pairs: ", keyvals))
		return nil
	}

	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}
	switch level {
	case log.LevelDebug:
		l.logger.Debug("", data...)
	case log.LevelInfo:
		l.logger.Info("", data...)
	case log.LevelWarn:
		l.logger.Warn("", data...)
	case log.LevelError:
		l.logger.Error("", data...)
	case log.LevelFatal:
		l.logger.Fatal("", data...)
	}
	return nil
}
