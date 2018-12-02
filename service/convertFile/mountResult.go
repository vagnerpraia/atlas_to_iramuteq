package convertFile

func mountResult(quizMap map[string][]string, headerMap map[string]string) map[string][]string {
	resultMap := make(map[string][]string)

	for index, quiz := range quizMap {
		for _, data := range quiz {
			_, ok := headerMap[index]
			if ok {
				resultMap[index] = append(resultMap[index], headerMap[index] + "\n")
				resultMap[index] = append(resultMap[index], data + "\n")
			}
		}
	}

	return resultMap
}