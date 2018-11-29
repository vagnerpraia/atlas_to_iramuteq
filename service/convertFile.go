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

	fmt.Println(version)

	if version == 7 {
		lines = lines[15:]
	}

	fmt.Println(lines[0])
	
	flagHead := true
	var keyInterview string
	for index, line := range lines {
		
		flagRegexHead := true
		if version == 8 {
			re := regexp.MustCompile("([0-9]*):([0-9]*) .*")
			flagRegexHead = re.MatchString(line)
		} else if version == 7 {
			flagRegexHead = strings.Constains("   (Super)", line)
		}

		if flagRegexHead && flagHead {
			keyInterview := ""
			if version == 8 {
				re := regexp.MustCompile(".1 [0-9]*: [0-9]*([a-zA-Z]1)?")
				subtringRegex := re.FindStringSubmatch(line)
			} else if version == 7 {
				re := regexp.MustCompile(".1 [0-9]*: [0-9]*([a-zA-Z]1)?")
				subtringRegex := re.FindStringSubmatch(line)
				keyInterview = subtringRegex[1 : len(subtringRegex) - 2]
			}

			regexSearch = re.search(r' - . [0-9]*: ', row, re.M|re.I)
			if regexSearch:
				keyStart = regexSearch.group()
				keyEnd = row[len(row) - 1]
			else:
				keyStart = ':'
				keyEnd = ' - '

			key_interview = find_between(row, keyStart, keyEnd).strip()
			if '.' in key_interview:
				key_interview = find_between(key_interview, '', '.').strip()
			if re.match(r'[0-9]*e', key_interview, re.M|re.I):
				key_interview = key_interview[:-1]
		}
		else:
			if row[0:6] != 'Codes:' and row[0:8] != 'No memos' and row != '':
				if interview_dict.has_key(key_interview):
					row_list = interview_dict.get(key_interview)
					row_list.append(row)
					interview_dict.update({key_interview : row_list})
				else:
					interview_dict.update({key_interview : [row]})
				flagHead = True
	*/
}

func ConvertFile(pathAtlasFile string, pathCsvFile string, pathResultFile string, separatorCsv string) bool {
	readFile(pathAtlasFile)

	return true
}