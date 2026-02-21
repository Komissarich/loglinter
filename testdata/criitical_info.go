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
	logger.Info("user Info" + password) // want `log message should not contain critical information like password`

																				   
}