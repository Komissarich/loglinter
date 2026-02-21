package logtest

import (
	"log"
	"log/slog"
)


func TestUppercase() {
	log.Println("Uppercase log") // want `log message 'Uppercase log' should be named 'uppercase log'`
	slog.Info("Uppercase log") // want `log message 'Uppercase log' should be named 'uppercase log'`
}