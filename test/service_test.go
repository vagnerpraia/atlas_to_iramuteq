package service

import (
	"testing"

	service "github.com/vagnerpraia/atlas_to_iramuteq/service"
)

func TestConvertFileVersion7(t *testing.T) {
	pathAtlasFile := "C:/Teste/atlas_to_iramuteq/questionario_version7.txt"
	pathCsvFile := "C:/Teste/atlas_to_iramuteq/escala.csv"
	pathResultFile := "C:/Teste/atlas_to_iramuteq/resultado_version7.txt"
	separatorCsv := ","

	response := service.ConvertFile(pathAtlasFile, pathCsvFile, pathResultFile, separatorCsv)
	if response == false {
		t.Errorf("Ocorreu um erro.")
	}
}

func TestConvertFileVersion8(t *testing.T) {
	pathAtlasFile := "C:/Teste/atlas_to_iramuteq/questionario_version8.txt"
	pathCsvFile := "C:/Teste/atlas_to_iramuteq/escala.csv"
	pathResultFile := "C:/Teste/atlas_to_iramuteq/resultado_version8.txt"
	separatorCsv := ";"

	response := service.ConvertFile(pathAtlasFile, pathCsvFile, pathResultFile, separatorCsv)
	if response == false {
		t.Errorf("Ocorreu um erro.")
	}
}