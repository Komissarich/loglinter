package golinters

import (
	"github.com/Komissarich/loglinter"
	"github.com/golangci/golangci-lint/pkg/goanalysis"

	"golang.org/x/tools/go/analysis"
)

func NewLogLinter() *goanalysis.Linter {
	return goanalysis.NewLinter(
		"loglinter",
		"Checks that all logs use english letters, start with lowercase letter, don't have special symbols or emoji, don't contain critical info",
		[]*analysis.Analyzer{loglinter.NewAnalyzer()},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
