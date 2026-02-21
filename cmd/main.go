package main

import (
	linter "github.com/Komissarich/loglinter"
	"golang.org/x/tools/go/analysis/singlechecker"
)


func main() {
	singlechecker.Main(linter.NewAnalyzer())
}