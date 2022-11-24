package main

import (
	"example.com/flag_finish_gomock_linter"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(flag_finish_gomock_linter.Analyzer) }
