package service

func mountResult(quizMap map[string][]string, headerMap map[string]string) map[string][]string {
	resultMap := make(map[string][]string)

	for index, quiz := range quizMap {
		header := headerMap[index]
		resultMap[index] = append(resultMap[index], header)

		for _, data := range quiz {
			resultMap[index] = append(resultMap[index], data)
		}		
	}

	return resultMap
}