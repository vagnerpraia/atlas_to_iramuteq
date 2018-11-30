package service

import (
	"fmt"
)

func ConvertFile(pathAtlasFile string, pathCsvFile string, pathResultFile string, separatorCsv string) bool {
	//interviewMap := readInterview(pathAtlasFile)
	//csvHeader, csvMap := readCsv(pathCsvFile, separatorCsv)
	_, csvMap := readCsv(pathCsvFile, separatorCsv)

	for index, data := range csvMap {
		if len(data) != 73 {
			fmt.Println(index)
			fmt.Println(data)
			fmt.Println(len(data))
			fmt.Println("")
		}
	}

	return true
}