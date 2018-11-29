package main

func main() {
	pathAtlasFile := "C:/Teste/atlas_to_iramuteq/questionario.txt"
	pathCsvFile := "C:/Teste/atlas_to_iramuteq/escala.csv"
	pathResultFile := "C:/Teste/atlas_to_iramuteq/resultado.txt"
	separatorCsv := ";"

	ConvertFile(pathAtlasFile, pathCsvFile, pathResultFile, separatorCsv)
}