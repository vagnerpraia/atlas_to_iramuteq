package service

import (
	"fmt"
)

func ConvertFile(pathAtlasFile string, pathCsvFile string, pathResultFile string, separatorCsv string) bool {
	//interviewMap := readInterview(pathAtlasFile)
	//csvHeader, csvMap := readCsv(pathCsvFile, separatorCsv)
	_, csvMap := readCsv(pathCsvFile, separatorCsv)
	fmt.Println(len(csvMap))

	for index, question := range csvMap {
		fmt.Println(index)
		fmt.Println(question)
		fmt.Println("")
	}

	return true
}