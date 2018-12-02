package convertFile

import (
	"github.com/vagnerpraia/atlas_to_iramuteq/model"
)

var Config model.Config

func Execute() bool {
	quizMap := readQuiz(Config.PathQuizFile)
	csvHeader, csvMap := readCsv(Config.PathCsvFile, Config.SeparatorCsv, Config.QuoteCsv)

	headerMap := mountHeader(csvHeader, csvMap)
	resultMap := mountResult(quizMap, headerMap)

	writeResult(resultMap, Config.PathResultFile)

	return true
}