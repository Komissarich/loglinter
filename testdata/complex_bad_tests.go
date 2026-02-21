package logtest

import (
	"log"
	"log/slog"

	"go.uber.org/zap"
)


func TestComplexBad() {
	password := "qwerty"
	token := "token"
	log.Println("User инфо") // want `log message 'User инфо' should be named 'user инфо';log message 'User инфо' should not use cyrillic characters`
	log.Println("привет" + password) // want `log message should not contain critical information like password;log message 'привет' should not use cyrillic characters`
	slog.Info("complex" + token + password) // want `log message should not contain critical information like password;log message should not contain critical information like token`
	logger, _ := zap.NewProduction()
	logger.Info("Very" + "сложный" + token + "log!") // want `log message should not contain critical information like token;log message should not use special symbols;log message 'сложный' should not use cyrillic characters;log message 'Very' should be named 'very'`

}
