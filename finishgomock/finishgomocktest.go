package flag_finish_gomock_linter_test

import (
	"testing"

	"example.com/flag_finish_gomock_linter"
	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, flag_finish_gomock_linter.Analyzer, "a")
}
