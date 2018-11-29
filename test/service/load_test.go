package service

import (
	"testing"

	service "gitlab.com/dadospublicosio/etl/siconv/service"
)

func TestLoadGeral(t *testing.T) {
	pathWork := "C:/Teste/dadospublicos.io"
	pathFile := "C:/Teste/dadospublicos.io/siconv.zip"
	separatorCsv := ";"

	response := service.Load(pathWork, pathFile, separatorCsv)
	if response == false {
		t.Errorf("Ocorreu um erro.")
	}
}

func TestLoadConsorcios(t *testing.T) {
	pathWork := "C:/Teste/dadospublicos.io"
	pathFile := "C:/Teste/dadospublicos.io/siconv_consorcios.zip"
	separatorCsv := ";"

	response := service.Load(pathWork, pathFile, separatorCsv)
	if response == false {
		t.Errorf("Ocorreu um erro.")
	}
}