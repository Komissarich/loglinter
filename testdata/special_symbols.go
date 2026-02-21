package logtest

import (
	"log"
	"log/slog"

	"go.uber.org/zap"
)



func TestSpecial() {
	log.Println("server started!")	// want `log message should not use special symbols`
	slog.Info("server started...") // want `log message should not use special symbols`
	logger, _ := zap.NewProduction()
	logger.Info("server started on port [8080]") // want `log message should not use special symbols`
}