package logtest

import (
	"log"
	"log/slog"
)


func TestGood() {
	log.Println("goodlog")
	slog.Debug("goodlog")
	slog.Info("goodlog")
	slog.Error("good log")
}