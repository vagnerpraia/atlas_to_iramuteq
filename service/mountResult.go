package service

func mountResult(quizMap map[string][]string, headerMap map[string]string) map[string][]string {
	resultMap := make(map[string][]string)

	for index, quiz := range quizMap {
		for _, data := range quiz {
			if value, ok := resultMap[index]; ok {
				resultMap[index] = append(value, headerMap[index] + "\n")
				resultMap[index] = append(value, data + "\n")
			}
		}
	}

	return resultMap
}