// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"net/url"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger is a simple logging interface
type Logger interface {
	Errorf(template string, args ...interface{})
	Infof(template string, args ...interface{})
}

type rollingLogger struct {
	*lumberjack.Logger
}

// Sync is a noop to implement zapcore.WriteSyncer interface
func (*rollingLogger) Sync() error {
	// noop
	return nil
}

func newLogger(appData string) (Logger, error) {
	rl := &rollingLogger{
		&lumberjack.Logger{
			Filename: filepath.Join(appData, "app.log"),
			MaxSize:  50,
			MaxAge:   3,
		},
	}

	zap.RegisterSink("rolling", func(*url.URL) (zap.Sink, error) {
		return rl, nil
	})

	output := [1]string{"stdout"}

	if !isDebug {
		output[0] = "rolling://"
	}

	ecfg := zap.NewProductionEncoderConfig()
	ecfg.EncodeTime = zapcore.ISO8601TimeEncoder

	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:         "console",
		OutputPaths:      output[:],
		ErrorOutputPaths: output[:],
		EncoderConfig:    ecfg,
	}

	zl, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	_ = zl

	return zl.Sugar(), nil
}
