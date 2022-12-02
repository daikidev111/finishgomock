package main

import (
	"github.com/daikidev111/finishgomock"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(finishgomock.Analyzer) }
