package main

import (
	loglinter "github.com/Komissarich/loglinter/pkg"
	"golang.org/x/tools/go/analysis/singlechecker"
)


func main() {
	singlechecker.Main(loglinter.NewAnalyzer().Analyzer)
}