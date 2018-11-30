package service

import (
    "bufio"
	"io"
    "log"
	"os"
	"regexp"
	"strings"
)

func readCsv(pathFile string, separator string) ([]string, map[string][]string) {
	csvFile, err := os.Open(pathFile)
	defer csvFile.Close()

	if err != nil {
        log.Fatal(err)
    }

	reader := bufio.NewReader(csvFile)

	var lines []string
    for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		lines = append(lines, line)
	}

	var header []string
	flagHeader := true
	csvMap := make(map[string][]string)
	for _, line := range lines {
		line = strings.Replace(line, "\n", "", 1)
		
		re := regexp.MustCompile("(\"[^\",]+),([^\"]+\")")
		stringsCommaRegex := re.FindStringSubmatch(line)
		for _, s := range stringsCommaRegex {
			stringAdjusted := strings.Replace(s, ",", "|||", 1)
			line = strings.Replace(line, s, stringAdjusted, 1)
		}

		re = regexp.MustCompile("(\"[^\";]+);([^\"]+\")")
		stringsSemicolonRegex := re.FindStringSubmatch(line)
		for _, s := range stringsSemicolonRegex {
			stringAdjusted := strings.Replace(s, ";", "\\\\\\", 1)
			line = strings.Replace(line, s, stringAdjusted, 1)
		}

		if flagHeader {
			header = strings.Split(line, separator)

			for index, _ := range header {
				header[index] = strings.Replace(header[index], "|||", ",", 1)
				header[index] = strings.Replace(header[index], "\\\\\\", ";", 1)
			}

			flagHeader = false
		} else {
			content := strings.Split(line, separator)

			key := strings.Replace(content[0], "e", "", 1)
			value := content[1:]

			for _, data := range value {
				data = strings.Replace(data, "|||", ",", 1)
				data = strings.Replace(data, "\\\\\\", ";", 1)

				csvMap[key] = append(csvMap[key], data)
			}
		}
	}

	return header, csvMap
}