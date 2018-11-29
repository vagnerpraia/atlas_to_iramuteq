package main

func main() {
	args := os.Args[1:]

	pathAtlasFile := args[0]
	pathCsvFile := args[1]

	pathResultFile := "result.txt"
	if len(args) > 2 {
		pathResultFile = args[2]
	}

	separatorCsv := ";"
	if len(args) > 3 {
		separatorCsv = args[3]
	}

	ConvertFile(pathAtlasFile, pathCsvFile, pathResultFile, separatorCsv)
}