package logtest

import (
	"log"
	"log/slog"

	"go.uber.org/zap"
)


func TestRussian() {
	log.Println("сtart server")	// want `log message 'сtart server' should not use cyrillic characters`
	slog.Info("старт сервера") // want `lыog message 'старт сервера' should not use cyrillic characters`
	logger, _ := zap.NewProduction()
	logger.Info("start sеrvеr") // want `log message 'start sеrvеr' should not use cyrillic characters`
}