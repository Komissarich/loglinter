package loglinter

import (
	"testing"

	loglinter "github.com/Komissarich/loglinter/pkg"
	"golang.org/x/tools/go/analysis/analysistest"
)


func TestLinter(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), loglinter.NewAnalyzer().Analyzer)
}