package service

import (
    "bufio"
	"fmt"
	"io"
    "log"
	"os"
	"regexp"
	"reflect"
	"strings"
)

func readInterview(pathFile string) map[string][]string {
    file, err := os.Open(pathFile)
    defer file.Close()

    if err != nil {
        log.Fatal(err)
    }

    reader := bufio.NewReader(file)

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

	re := regexp.MustCompile("Report: [0-9]* quotation\\(s\\) for [0-9]* code( *)?")
	flagRegexVersion := re.MatchString(lines[0])

	version := 8
	if flagRegexVersion {
		version = 7
	}

	if version == 7 {
		lines = lines[15:]
	}

	interviewMap := make(map[string][]string)
	var keyInterview string
	for _, line := range lines {
		flagHead := true
		if version == 8 {
			re := regexp.MustCompile("\\([0-9]*:[0-9]*\\) - [a-zA-Z] ?[0-9]*: [0-9]*e?$")
			flagHead = re.MatchString(line)
		} else if version == 7 {
			flagHead = strings.Contains(line, "   (Super)")
		}

		if flagHead {
			var pattern string
			if version == 8 {
				pattern = "[a-zA-Z] ?[0-9]*: [0-9]*e?$"
			} else if version == 7 {
				pattern = "^[a-zA-Z] ?[0-9]*: [0-9]*e?"
			}
			re := regexp.MustCompile(pattern)
			substringRegex := re.FindStringSubmatch(line)[0]

			re = regexp.MustCompile("[0-9]*e?$")
			substringRegex = re.FindStringSubmatch(substringRegex)[0]

			keyInterview = strings.Replace(substringRegex, "e", "", 1)
		} else if len(line) > 0 {
			flagContainsCode := strings.Contains("Codes:", line)
			flagContainsNoMemos := strings.Contains("No memos", line)
			if flagContainsCode && flagContainsNoMemos {
				interviewMap[keyInterview] = append(interviewMap[keyInterview], line)
			}
		}
	}

	return interviewMap
}

func ConvertFile(pathAtlasFile string, pathCsvFile string, pathResultFile string, separatorCsv string) bool {
	interview := readInterview(pathAtlasFile)
	fmt.Println(reflect.TypeOf(interview))

	return true
}