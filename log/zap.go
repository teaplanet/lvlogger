package log

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

const (
	LVLOGGER_ENVIRONMENT = "LVLOGGER_ENV"

	ModeProduction  = "Prod"
	ModeDevelopment = "Dev"
)

var (
	Mode          = initMode()
	defaultLogger = NewLogger()
	defaultSugar  = defaultLogger.Sugar()

	Debug  = defaultSugar.Debug
	Debugf = defaultSugar.Debugf
	Debugw = defaultSugar.Debugw
	Warn   = defaultSugar.Warn
	Warnf  = defaultSugar.Warnf
	Warnw  = defaultSugar.Warnw
	Info   = defaultSugar.Info
	Infof  = defaultSugar.Infof
	Infow  = defaultSugar.Infow
	Error  = defaultSugar.Error
	Errorf = defaultSugar.Errorf
	Errorw = defaultSugar.Errorw
	Fatal  = defaultSugar.Fatal
	Fatalf = defaultSugar.Fatalf
	Fatalw = defaultSugar.Fatalw
)

func NewLogger(options ...zap.Option) *zap.Logger {
	var logger *zap.Logger
	switch Mode {
	case ModeDevelopment:
		logger, _ = zap.NewDevelopment(options...)
	case ModeProduction:
		logger, _ = zap.NewProduction(options...)
	default:
		panic("Mode is empty")
	}
	return logger
}

func SetLogger(logger *zap.Logger)      { defaultLogger = logger }
func SetSugar(sugar *zap.SugaredLogger) { defaultSugar = sugar }

func initMode() string {
	env := os.Getenv(LVLOGGER_ENVIRONMENT)
	if strings.ToLower(env) == "p" {
		return ModeProduction
	}
	return ModeDevelopment
}
