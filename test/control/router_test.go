package control

import (
	"testing"

	"github.com/vagnerpraia/atlas_to_iramuteq/model"
	"github.com/vagnerpraia/atlas_to_iramuteq/control"
)

func TestRouterVersion7(t *testing.T) {
	pathQuizFile := "C:/Teste/atlas_to_iramuteq/questionario_version7.txt"
	pathCsvFile := "C:/Teste/atlas_to_iramuteq/escala_version7.csv"
	pathResultFile := "C:/Teste/atlas_to_iramuteq/resultado_version7.txt"
	separatorCsv := ","
	quoteCsv := "\""

	control.Config = model.Config{
		pathQuizFile,
		pathCsvFile,
		pathResultFile,
		separatorCsv,
		quoteCsv,
	}

	response := control.Router(pathQuizFile)

	if response == false {
		t.Errorf("Ocorreu um erro.")
	}
}

func TestRouterVersion8(t *testing.T) {
	pathQuizFile := "C:/Teste/atlas_to_iramuteq/questionario_version8.txt"
	pathCsvFile := "C:/Teste/atlas_to_iramuteq/escala_version8.csv"
	pathResultFile := "C:/Teste/atlas_to_iramuteq/resultado_version8.txt"
	separatorCsv := ","
	quoteCsv := "\""

	control.Config = model.Config{
		pathQuizFile,
		pathCsvFile,
		pathResultFile,
		separatorCsv,
		quoteCsv,
	}

	response := control.Router(pathQuizFile)

	if response == false {
		t.Errorf("Ocorreu um erro.")
	}
}