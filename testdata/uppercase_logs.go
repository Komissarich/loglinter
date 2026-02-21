package logtest

import (
	"log"
	"log/slog"

	"go.uber.org/zap"
)


func TestUppercase() {
	log.Println("Uppercase log") // want `log message 'Uppercase log' should be named 'uppercase log'`
	slog.Info("Uppercase log") // want `log message 'Uppercase log' should be named 'uppercase log'`
	logger, _ := zap.NewProduction()
	logger.Info("Uppercase log") // want `log message 'Uppercase log' should be named 'uppercase log'`
}