package logtest

import (
	"log"
	"log/slog"

	"go.uber.org/zap"
)


func TestComplexGood() {
	log.Default().Println("complex good log")
	slog.With().Info("complex good log")
	logger, _ := zap.NewProduction()
	logger.Sugar().Info("complex good log")

}
