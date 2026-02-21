package logtest

import (
	"log"
	"log/slog"

	"go.uber.org/zap"
)


func TestCriticalInfo() {
	password := "qwerty"
	log.Println("user password" + password) // want `log message should not contain critical information like password`
	slog.Info("user password" + password) // want `log message should not contain critical information like password`
	logger, _ := zap.NewProduction()
	logger.Info("user password" + password) // want `log message should not contain critical information like password`
	token := "api_token"
	log.Println("api token" + token) // want `log message should not contain critical information like token`
	slog.Info("api token" + token) // want `log message should not contain critical information like token`
	logger.Info("api token" + token) // want `log message should not contain critical information like token`
}