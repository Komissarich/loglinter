package main

import (
	"github.com/Komissarich/loglinter"
	"golang.org/x/tools/go/analysis"
)

// AnalyzerPlugin - экспортируемая переменная для golangci-lint
var AnalyzerPlugin analyzerPlugin

type analyzerPlugin struct{}

// GetAnalyzers возвращает список анализаторов
func (analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		loglinter.NewAnalyzer(),
	}
}

// Пустой main для компиляции
func main() {}