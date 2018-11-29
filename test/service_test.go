package service

import (
	"testing"

	service "github.com/vagnerpraia/atlas_to_iramuteq"
)

func TestConvertFile(t *testing.T) {
	pathAtlasFile := "C:/Teste/atlas_to_iramuteq/questionario.txt"
	pathCsvFile := "C:/Teste/atlas_to_iramuteq/escala.csv"
	pathResultFile := "C:/Teste/atlas_to_iramuteq/resultado.txt"
	separatorCsv := ";"

	response := service.Load(pathWork, pathFile, separatorCsv)
	if response == false {
		t.Errorf("Ocorreu um erro.")
	}
}