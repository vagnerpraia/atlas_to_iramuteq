package service

func ConvertFile(pathQuizFile string, pathCsvFile string, pathResultFile string, separatorCsv string) bool {
	quizMap := readQuiz(pathQuizFile)
	csvHeader, csvMap := readCsv(pathCsvFile, separatorCsv)
	headerMap := mountHeader(csvHeader, csvMap)
	resultMap := mountResult(quizMap, headerMap)
	writeResult(resultMap, pathResultFile)

	return true
}