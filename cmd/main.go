package main

import (
	"github.com/Komissarich/loglinter"
	"golang.org/x/tools/go/analysis/singlechecker"
)


func main() {
	singlechecker.Main(loglinter.NewAnalyzer().Analyzer)
}