package golinters

import (
	"github.com/Komissarich/loglinter"
	"github.com/golangci/golangci-lint/pkg/goanalysis"
	"golang.org/x/tools/go/analysis"
)

func NewLogLinter() *goanalysis.Linter {  // ← маленькая l после Log
    return goanalysis.NewLinter(
        "loglinter",
        "Проверяет логи на английский, отсутствие кириллицы, эмодзи, критических данных",
        []*analysis.Analyzer{loglinter.NewAnalyzer()},
        nil,
    ).WithLoadMode(goanalysis.LoadModeSyntax)
}