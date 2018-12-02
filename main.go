package main

import (
	"os"

	"github.com/vagnerpraia/atlas_to_iramuteq/control"
	"github.com/vagnerpraia/atlas_to_iramuteq/model"
)

func main() {
	args := os.Args[1:]

	pathQuizFile := "questionario.txt"
	if len(args) > 0 {
		pathQuizFile = args[0]
	}

	pathCsvFile := "escala.csv"
	if len(args) > 1 {
		pathCsvFile = args[1]
	}

	pathResultFile := "result.txt"
	if len(args) > 2 {
		pathResultFile = args[2]
	}

	separatorCsv := ";"
	if len(args) > 3 {
		separatorCsv = args[3]
	}

	quoteCsv := "\""
	if len(args) > 4 {
		quoteCsv = args[4]
	}

	control.Config = model.Config{
		pathQuizFile,
		pathCsvFile,
		pathResultFile,
		separatorCsv,
		quoteCsv,
	}

	control.Router(pathQuizFile)
}