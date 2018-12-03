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
		if len(args[0]) > 0 {
			pathQuizFile = args[0]
		}
	}

	pathCsvFile := "escala.csv"
	if len(args) > 1 {
		if len(args[1]) > 0 {
			pathCsvFile = args[1]
		}
	}

	pathResultFile := "resultado.txt"
	if len(args) > 2 {
		if len(args[2]) > 0 {
			pathResultFile = args[2]
		}
	}

	separatorCsv := ";"
	if len(args) > 3 {
		if len(args[3]) > 0 {
			separatorCsv = args[3]
		}
	}

	quoteCsv := "\""
	if len(args) > 4 {
		if len(args[4]) > 0 {
			quoteCsv = args[4]
		}
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