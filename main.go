package main

import (
	"os"

	service "github.com/vagnerpraia/atlas_to_iramuteq/service"
)

func main() {
	args := os.Args[1:]

	pathQuizFile := args[0]
	pathCsvFile := args[1]

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

	service.ConvertFile(pathQuizFile, pathCsvFile, pathResultFile, separatorCsv, quoteCsv)
}