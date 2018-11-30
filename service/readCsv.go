package service

import (
    "bufio"
	"io"
    "log"
	"os"
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

		if flagHeader {
			header = strings.Split(line, separator)
			flagHeader = false
		} else {
			content := strings.Split(line, separator)

			key := strings.Replace(content[0], "e", "", 1)
			value := content[1:]

			for _, v := range value {
				csvMap[key] = append(csvMap[key], v)
			}
		}
	}

	return header, csvMap
}