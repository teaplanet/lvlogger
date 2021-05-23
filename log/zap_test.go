package log

import (
	"testing"

	"go.uber.org/zap"
)

func TestSetLogger(t *testing.T) {
	type args struct {
		level zap.AtomicLevel
	}
	tests := []struct {
		name string
		args args
	}{{
		name: "DEBUG Level",
		args: args{
			level: zap.NewAtomicLevelAt(zap.DebugLevel),
		},
	}, {
		name: "INFO Level",
		args: args{
			level: zap.NewAtomicLevelAt(zap.InfoLevel),
		},
	}}
	for _, tt := range tests {
		conf := zap.NewDevelopmentConfig()
		conf.Level = tt.args.level
		RebuildLogger(conf)
		t.Run(tt.name, func(t *testing.T) {
			Debug("DEBUG Level")
			Info("INFO Level")
		})
	}
}

func TestSetLoggerLevel(t *testing.T) {
	type args struct {
		level zap.AtomicLevel
	}
	tests := []struct {
		name string
		args args
	}{{
		name: "DEBUG Level",
		args: args{
			level: zap.NewAtomicLevelAt(zap.DebugLevel),
		},
	}, {
		name: "INFO Level",
		args: args{
			level: zap.NewAtomicLevelAt(zap.InfoLevel),
		},
	}}
	for _, tt := range tests {
		SetLevel(tt.args.level)
		t.Run(tt.name, func(t *testing.T) {
			Debug("DEBUG Level")
			Info("INFO Level")
		})
	}
}