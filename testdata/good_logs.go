package logtest

import (
	"log"
	"log/slog"

	"go.uber.org/zap"
)


func TestGood() {
	log.Println("goodlog")
	slog.Debug("goodLog123")
	logger, _ := zap.NewProduction()
	logger.Info("veryGoodLog1111") 
}