package service

import (
	"fmt"
)

func ConvertFile(pathAtlasFile string, pathCsvFile string, pathResultFile string, separatorCsv string) bool {
	//interviewMap := readInterview(pathAtlasFile)
	csvHeader, csvMap := readCsv(pathCsvFile, separatorCsv)
	headerMap := mountHeader(csvHeader, csvMap)

	for index, data := range headerMap {
		fmt.Println(index)
		fmt.Println(data)
		fmt.Println("")
	}

	return true
}