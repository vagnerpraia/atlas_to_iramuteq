package service

func mountResult(quizMap map[string][]string, headerMap map[string][]string) map[string][]string {
	resultMap := make(map[string][]string)
	/*
	for index, data := range quizMap {
		header := "****"
		for k, v := range data {
			header += " *" + csvHeader[k] + "_" + v
		}
		headerMap[index] = append(headerMap[index], header)
	}
	*/
	return resultMap
}