package convertFile

import (
    "log"
    "os"
)

func writeResult(resultMap map[string][]string, pathResultFile string) {
    f, err := os.OpenFile(pathResultFile, os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
	}

	for _, result := range resultMap {
		for _, data := range result {
			if _, err := f.Write([]byte(data))
			err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := f.Close()
	err != nil {
        log.Fatal(err)
    }
}