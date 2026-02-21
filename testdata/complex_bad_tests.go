package logtest

import (
	"log"
	"log/slog"
)


func TestComplexBad() {
	password := "qwerty"
	token := "token"
	log.Println("User инфо") // want `log message 'User инфо' should be named 'user инфо';log message 'User инфо' should not use cyrillic characters`
	log.Println("hello" + password) // want `log message 'привет' should not use cyrillic characters;log message should not contain critical information like password`
	slog.Info("complex" + token + password) // want `log message should not contain critical information like password;log message should not contain critical information like token`
	// logger, _ := zap.NewProduction()
	// logger.Info("Very", zap.String("complex", "комплексный"), zap.String("Log", "token" + token)) 

}
