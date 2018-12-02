package help

import (
	"fmt"
)

func Execute() bool {
	fmt.Println("Lista de parâmetros e sua respectiva ordem de utilização:")
	fmt.Println("    pathQuizFile: Caminho do arquivo com o questionário do Atlas (default: questionario.txt)")
	fmt.Println("    pathCsvFile: Caminho do arquivo CSV com o dicionário de labels (default: escala.csv)")
	fmt.Println("    pathResultFile: Caminho do arquivo com o resultado do processamento (default: resultado.txt)")
	fmt.Println("    separatorCsv: Separador utilizado no CSV (default: ;)")
	fmt.Println("    quoteCsv: Indicador de fronteira dos dados utilizado no CSV (default: \")")
	fmt.Println("")
	fmt.Println("Para mais informações acesse: github.com/vagnerpraia/atlas_to_iramuteq")

	return true
}