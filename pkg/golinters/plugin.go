package golinters

import (
	"github.com/Komissarich/loglinter"
	"golang.org/x/tools/go/analysis"
)

// New - это функция, которую ожидает golangci-lint для версии 1.x
func New(conf any) ([]*analysis.Analyzer, error) {
    return []*analysis.Analyzer{
        loglinter.NewAnalyzer(),
    }, nil
}

// main здесь технически не нужен, но так требует package main. Он не будет использоваться.
func main() {}