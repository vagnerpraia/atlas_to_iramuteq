package service

import (
    "bufio"
	"io"
    "log"
	"os"
	"regexp"
	"strings"
)

func readQuiz(pathFile string) map[string][]string {
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

	quizMap := make(map[string][]string)
	var key string
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

			key = strings.Replace(substringRegex, "e", "", 1)
		} else {
			line = strings.Replace(line, "\n", "", 1)
			flagContainsContent := false
			if len(line) > 0 {
				flagContainsContent = true
			}

			flagContainsCode := strings.Contains(line, "Codes:")
			flagContainsNoMemos := strings.Contains(line, "No memos")
			if flagContainsContent && !flagContainsCode && !flagContainsNoMemos {
				quizMap[key] = append(quizMap[key], line)
			}
		}
	}

	return quizMap
}