package service

import (
	"fmt"
)

func ConvertFile(pathQuizFile string, pathCsvFile string, pathResultFile string, separatorCsv string) bool {
	//quizMap := readQuiz(pathQuizFile)
	csvHeader, csvMap := readCsv(pathCsvFile, separatorCsv)
	headerMap := mountHeader(csvHeader, csvMap)
	//resultMap := mountResult(quizMap, headerMap)

	for index, data := range headerMap {
		fmt.Println(index)
		fmt.Println(data)
		fmt.Println("")
	}

	return true
}