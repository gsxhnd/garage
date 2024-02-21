package utils

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debugf(template string, args ...interface{})
	Debugw(template string, keysAndValues ...interface{})
	Infof(template string, args ...interface{})
	Infow(template string, keysAndValues ...interface{})
	Warnf(template string, args ...interface{})
	Warnw(template string, keysAndValues ...interface{})
	Errorf(template string, args ...interface{})
	Errorw(template string, keysAndValues ...interface{})
}

type logger struct {
	// Logger *zap.Logger
	Suger *zap.SugaredLogger
}

var (
	devEncoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "time",
		CallerKey:   "caller",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000000-07:00"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
	prodEncoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "time",
		CallerKey:   "caller",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000000-07:00"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
)

func NewLogger(cfg *Config) Logger {
	var encode zapcore.Encoder
	var level zapcore.Level

	if cfg.Dev {
		encode = devEncoder
	} else {
		encode = prodEncoder
	}
	if cfg.LogLevel == "debug" {
		level = zap.DebugLevel
	} else if cfg.LogLevel == "info" {
		level = zap.InfoLevel
	} else {
		level = zap.WarnLevel
	}

	var core = zapcore.NewTee(
		zapcore.NewCore(encode, os.Stdout, level),
	)
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &logger{
		// Logger: zapLogger,
		Suger: zapLogger.Sugar(),
	}
}

func (l *logger) Debugf(template string, args ...interface{}) {
	defer l.Suger.Sync()
	l.Suger.Debugf(template, args...)
}
func (l *logger) Debugw(template string, keysAndValues ...interface{}) {
	defer l.Suger.Sync()
	l.Suger.Debugw(template, keysAndValues...)
}

func (l *logger) Infof(template string, args ...interface{}) {
	defer l.Suger.Sync()
	l.Suger.Infof(template, args...)
}
func (l *logger) Infow(template string, keysAndValues ...interface{}) {
	defer l.Suger.Sync()
	l.Suger.Infow(template, keysAndValues...)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	defer l.Suger.Sync()
	l.Suger.Warnf(template, args...)
}
func (l *logger) Warnw(template string, keysAndValues ...interface{}) {
	defer l.Suger.Sync()
	l.Suger.Warnw(template, keysAndValues...)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	defer l.Suger.Sync()
	l.Suger.Errorf(template, args...)
}
func (l *logger) Errorw(template string, keysAndValues ...interface{}) {
	defer l.Suger.Sync()
	l.Suger.Errorw(template, keysAndValues...)
}

func GetLogger() *zap.Logger {
	debugEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "ts",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000000-07:00"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
	var core = zapcore.NewTee(
		zapcore.NewCore(debugEncoder, os.Stdout, zap.DebugLevel),
	)
	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}
