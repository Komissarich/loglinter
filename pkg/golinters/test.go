package golinters

import "github.com/golangci/golangci-lint/pkg/goanalysis"

func NewTestlinter() *goanalysis.Linter {
	return goanalysis.NewLinter("testlinter", "Test", nil, nil)
}