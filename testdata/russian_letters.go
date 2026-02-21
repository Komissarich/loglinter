package logtest

import (
	"log"
	"log/slog"
)



func TestRussian() {
	log.Println("сtart server")	// want `log message 'сtart server' should not use cyrillic characters`
	slog.Info("старт сервера") // want `log message 'старт сервера' should not use cyrillic characters`
}