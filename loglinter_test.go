package loglinter

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)


func TestLinter(t *testing.T) {
	
	analysistest.Run(t, analysistest.TestData(), NewAnalyzer())
}