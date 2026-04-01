package helper

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseFile(fileName string) (Count, error) {

	fmt.Println("Parsing File", fileName)
	file, err := os.Open(fileName)

	if err != nil {
		//fmt.Println("Unable to read file ", err)
		return Count{}, nil
	}
	defer file.Close()
	scanBuf := bufio.NewScanner(file)
	count := Count{}

	count.FileName = fileName
	for scanBuf.Scan() {

		count.NumLines++

		if strings.Contains(scanBuf.Text(), "INFO") {
			count.NumInfo++
		} else if strings.Contains(scanBuf.Text(), "DEBUG") {
			count.NumDebug++
		} else if strings.Contains(scanBuf.Text(), "ERROR") {
			count.NumErr++
		} else if strings.Contains(scanBuf.Text(), "WARN") {
			count.NumWarn++
		} else {
			count.UnknownLines++
		}
	}
	if err := scanBuf.Err(); err != nil {
		//fmt.Println("Error while reading file:", err)

		return Count{}, nil
	}

	return count, nil

}
