package service

func ConvertFile(pathQuizFile string, pathCsvFile string, pathResultFile string, separatorCsv string, quoteCsv string) bool {
	quizMap := readQuiz(pathQuizFile)
	csvHeader, csvMap := readCsv(pathCsvFile, separatorCsv, quoteCsv)
	headerMap := mountHeader(csvHeader, csvMap)
	resultMap := mountResult(quizMap, headerMap)
	writeResult(resultMap, pathResultFile)

	return true
}