package log

import (
	"os"
	"strings"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

const (
	LVLOGGER_ENVIRONMENT = "LVLOGGER_ENV"

	ModeProduction  = "Prod"
	ModeDevelopment = "Dev"
	ModeGCP         = "GCP"
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
	case ModeGCP:
		logger, _ = zapdriver.NewProductionWithCore(
			zapdriver.WrapCore(zapdriver.ReportAllErrors(true)), options...)
	default:
		panic("Mode is empty")
	}
	return logger
}

func SetLogger(logger *zap.Logger) {
	updateLoggerFuncs(logger)
}

func SetLevel(level zap.AtomicLevel) {
	var conf zap.Config
	switch Mode {
	case ModeDevelopment:
		conf = zap.NewDevelopmentConfig()
	case ModeProduction:
		conf = zap.NewProductionConfig()
	default:
		panic("Mode is empty")
	}
	conf.Level = level
	err := RebuildLogger(conf)
	if err != nil {
		panic(err)
	}
}

func RebuildLogger(conf zap.Config, opts ...zap.Option) error {
	logger, err := conf.Build(opts...)
	if err != nil {
		return err
	}
	updateLoggerFuncs(logger)
	return nil
}

func initMode() string {
	env := strings.ToLower(os.Getenv(LVLOGGER_ENVIRONMENT))
	if env == "" {
		return ModeDevelopment
	}

	switch env[0:1] {
	case "p":
		return ModeProduction
	case "g":
		return ModeGCP
	default:
		return ModeDevelopment
	}
}

func updateLoggerFuncs(logger *zap.Logger) {
	defaultLogger = logger
	defaultSugar = defaultLogger.Sugar()
	Debug = defaultSugar.Debug
	Debugf = defaultSugar.Debugf
	Debugw = defaultSugar.Debugw
	Warn = defaultSugar.Warn
	Warnf = defaultSugar.Warnf
	Warnw = defaultSugar.Warnw
	Info = defaultSugar.Info
	Infof = defaultSugar.Infof
	Infow = defaultSugar.Infow
	Error = defaultSugar.Error
	Errorf = defaultSugar.Errorf
	Errorw = defaultSugar.Errorw
	Fatal = defaultSugar.Fatal
	Fatalf = defaultSugar.Fatalf
	Fatalw = defaultSugar.Fatalw
}
