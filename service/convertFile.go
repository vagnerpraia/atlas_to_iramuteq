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
	//fmt.Println(reflect.TypeOf(reader))

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
	for index, line := range lines {
		
		flagRegexHead := true
		if version == 8 {
			re := regexp.MustCompile("([0-9]*):([0-9]*) .*")
			flagRegexHead = re.MatchString(line)
		} else if version == 7 {
			flagRegexHead = strings.Constains("   (Super)", line)
		}

		if flagRegexHead {
			re := regexp.MustCompile(".1 [0-9]*: [0-9]*([a-zA-Z]1)?")
			substringRegex := re.FindStringSubmatch(line)
			substringAdjusted := substringRegex[strings.Index(":") + 2 :]

			re = regexp.MustCompile("[0-9]*")
			keyInterview := re.FindStringSubmatch(substringAdjusted)
		} else if len(line) > 0 && line[0:6] != "Codes:" && line[0:8] != "No memos" {
			interviewMap[key_interview] = append(interviewMap[key_interview], line)
		}
	}
}

func ConvertFile(pathAtlasFile string, pathCsvFile string, pathResultFile string, separatorCsv string) bool {
	readFile(pathAtlasFile)

	return true
}