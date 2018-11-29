package service

import (
    "bufio"
	"fmt"
	"io"
    "log"
	"os"
	"regexp"
	//"reflect"
	"strings"
)

func readFile(pathFile string) {
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

		//fmt.Println(reflect.TypeOf(line))

		if flagHead {
			var substringRegex []string
			if version == 8 {
				re := regexp.MustCompile("[a-zA-Z] ?[0-9]*: [0-9]*e?$")
				substringRegex = re.FindStringSubmatch(line)
			} else if version == 7 {
				re := regexp.MustCompile("^a-zA-Z] ?[0-9]*: [0-9]*e?")
				substringRegex = re.FindStringSubmatch(line)
			}

			fmt.Println(substringRegex)

			//substringAdjusted := substringRegex[strings.Index(substringRegex, ":") + 2 :]

			//re = regexp.MustCompile("[0-9]*")
			//keyInterview = re.FindStringSubmatch(substringAdjusted)
		} else if len(line) > 0 {
			flagContainsCode := strings.Contains("Codes:", line)
			flagContainsNoMemos := strings.Contains("No memos", line)
			if flagContainsCode && flagContainsNoMemos {
				interviewMap[keyInterview] = append(interviewMap[keyInterview], line)
			}
		}
	}
}

func ConvertFile(pathAtlasFile string, pathCsvFile string, pathResultFile string, separatorCsv string) bool {
	readFile(pathAtlasFile)

	return true
}