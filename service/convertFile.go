package service

import (
    "bufio"
	"fmt"
	"io"
    "log"
	"os"
	"regexp"
	//"reflect"
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
	
	flagRegex, err := regexp.MatchString("p(Report: [0-9]* quotation(s) for [0-9]* code(.*)?)ch", lines[0])
	if err != nil {
		log.Fatal(err)
	}

	version := 8
	if flagRegex {
		version = 7
		fmt.Println("Version 7")
	} else {
		fmt.Println("Version 8")
	}

	fmt.Println(version)
}

func ConvertFile(pathAtlasFile string, pathCsvFile string, pathResultFile string, separatorCsv string) bool {
	readFile(pathAtlasFile)

	return true
}