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

func PrintSummary(count Count) {

	fmt.Println("Filename ", count.FileName)
	fmt.Println("Number of Lines ", count.NumLines)
	fmt.Println("Info ", count.NumInfo)
	fmt.Println("Warn ", count.NumWarn)
	fmt.Println("Error ", count.NumErr)
	fmt.Println("Debug ", count.NumDebug)
	fmt.Println("Unknown ", count.UnknownLines)

}

func UpdateTotalCount(totalCount *TotalCount, count Count) {
	totalCount.TotalNumInfo += count.NumInfo
	totalCount.TotalNumWarn += count.NumWarn
	totalCount.TotalNumErr += count.NumErr
	totalCount.TotalNumDebug += count.NumDebug
	totalCount.TotalNumLines += count.NumLines
	totalCount.TotalUnknownLines += count.UnknownLines
}

func PrintTotalSummary(totalCount *TotalCount) {

	fmt.Println("Total Number Of  Lines ", totalCount.TotalNumLines)
	fmt.Println("Total Info Count", totalCount.TotalNumInfo)
	fmt.Println("Total Warn Count", totalCount.TotalNumWarn)
	fmt.Println("Total Error Count", totalCount.TotalNumErr)
	fmt.Println("Total Debug Count", totalCount.TotalNumDebug)
	fmt.Println("Total Unknown Count", totalCount.TotalUnknownLines)

}
