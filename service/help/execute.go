package help

import (
	"fmt"
)

func Execute() bool {
	fmt.Println("Lista de parâmetros e sua respectiva ordem de utilização:")
	fmt.Println("    pathQuizFile = Caminho do arquivo com o questionário do Atlas")
	fmt.Println("    pathCsvFile = Caminho do arquivo CSV com o dicionário de labels")
	fmt.Println("    pathResultFile = Caminho do arquivo com o resultado do processamento")
	fmt.Println("    separatorCsv = Separador utilizado no CSV")
	fmt.Println("    quoteCsv = Indicador de fronteira dos dados utilizado no CSV")
	fmt.Println("")
	fmt.Println("Para mais informações acesse: github.com/vagnerpraia/atlas_to_iramuteq")

	return true
}