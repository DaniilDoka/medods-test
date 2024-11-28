package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	logger zerolog.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: log.Output(zerolog.ConsoleWriter{Out: os.Stderr}),
	}
}

func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.logger.Info().Msgf(template, args...)
}

func (l *Logger) Fatal(msg string) {
	l.logger.Fatal().Msg(msg)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatal().Msgf(template, args...)
}

func (l *Logger) Panic(msg string) {
	l.logger.Panic().Msg(msg)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.logger.Panic().Msgf(template, args...)
}

func (l *Logger) PanicError(err error) {
	l.logger.Panic().Msg(err.Error())
}

func (l *Logger) Error(err error) {
	l.logger.Error().Msg(err.Error())
}

func (l *Logger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.logger.Warn().Msgf(template, args...)
}
