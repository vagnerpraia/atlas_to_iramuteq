package convertFile

import (
    "bufio"
	"io"
    "log"
	"os"
	"regexp"
	"strings"
)

func readQuiz(pathQuizFile string) map[string][]string {
    file, err := os.Open(pathQuizFile)
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
		re := regexp.MustCompile(`\r?\n`)
		line = re.ReplaceAllString(line, " ")
		line = strings.TrimSpace(line)

		flagHead := true
		if version == 8 {
			re := regexp.MustCompile("\\([0-9]*:[0-9]*\\) - [a-zA-Z] ?[0-9]*: ([0-9]*)(([^0-9])*?)$")
			flagHead = re.MatchString(line)
		} else if version == 7 {
			flagHead = strings.Contains(line, "   (Super)")
		}

		if flagHead {
			var pattern string
			if version == 8 {
				pattern = "[a-zA-Z] ?[0-9]*: ([0-9]*)(([^0-9])*?)$"
			} else if version == 7 {
				pattern = "^[a-zA-Z] ?[0-9]*: ([0-9]*)(([^0-9])*?)"
			}
			re := regexp.MustCompile(pattern)
			substringRegex := re.FindStringSubmatch(line)[0]

			re = regexp.MustCompile("([0-9]*)(([^0-9])*?)$")
			substringRegex = re.FindStringSubmatch(substringRegex)[0]

			reKey, err := regexp.Compile("[^0-9]+")
			if err != nil {
				log.Fatal(err)
			}
			key = reKey.ReplaceAllString(substringRegex, "")
		} else {
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