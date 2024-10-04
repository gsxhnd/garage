package utils

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
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
	Panicf(template string, args ...interface{})
	Panicw(template string, keysAndValues ...interface{})
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
	var (
		level zapcore.Level
		core  zapcore.Core
	)

	level = zap.InfoLevel
	switch cfg.Log.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	default:
		level = zap.InfoLevel
	}

	if !(cfg.Mode == "dev") {
		core = zapcore.NewCore(prodEncoder, zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.Log.FileName,
			MaxBackups: cfg.Log.MaxBackups,
			MaxAge:     cfg.Log.MaxAge,
		},
		), level)
	} else {
		core = zapcore.NewCore(devEncoder, os.Stdout, level)
	}

	return &logger{
		Suger: zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar(),
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

func (l *logger) Panicf(template string, args ...interface{}) {
	defer l.Suger.Sync()
	l.Suger.Panicf(template, args...)
}
func (l *logger) Panicw(template string, keysAndValues ...interface{}) {
	defer l.Suger.Sync()
	l.Suger.Panicw(template, keysAndValues...)
}
