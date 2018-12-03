package convertFile

func mountHeader(csvHeader []string, csvMap map[string][]string) map[string]string {
	headerMap := make(map[string]string)

	for index, data := range csvMap {
		header := "****"
		for k, v := range data {
			if k < len(csvHeader) {
				header += " *" + csvHeader[k] + "_" + v
			}
		}
		headerMap[index] = header
	}

	return headerMap
}