package service

import (
	"testing"

	service "github.com/vagnerpraia/atlas_to_iramuteq/service"
)

func TestConvertFileVersion7(t *testing.T) {
	pathQuizFile := "C:/Teste/atlas_to_iramuteq/questionario_version7.txt"
	pathCsvFile := "C:/Teste/atlas_to_iramuteq/escala.csv"
	pathResultFile := "C:/Teste/atlas_to_iramuteq/resultado_version7.txt"
	separatorCsv := ","
	quoteCsv := "\""

	response := service.ConvertFile(pathQuizFile, pathCsvFile, pathResultFile, separatorCsv, quoteCsv)
	if response == false {
		t.Errorf("Ocorreu um erro.")
	}
}

func TestConvertFileVersion8(t *testing.T) {
	pathQuizFile := "C:/Teste/atlas_to_iramuteq/questionario_version8.txt"
	pathCsvFile := "C:/Teste/atlas_to_iramuteq/escala.csv"
	pathResultFile := "C:/Teste/atlas_to_iramuteq/resultado_version8.txt"
	separatorCsv := ";"
	quoteCsv := "\""

	response := service.ConvertFile(pathQuizFile, pathCsvFile, pathResultFile, separatorCsv, quoteCsv)
	if response == false {
		t.Errorf("Ocorreu um erro.")
	}
}