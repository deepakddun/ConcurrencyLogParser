package main

import (
	"fmt"
	"os"
	"sync"

	"example.com/LogAnalyser/internal/helper"
)

func main() {

	var wg sync.WaitGroup
	args := os.Args

	if len(args) < 2 {
		fmt.Println(" Not sufficient arguments ")
		os.Exit(1)
	}
	//countMap := map[string]helper.Count{}
	totalCount := helper.TotalCount{}
	ch := make(chan helper.Result, len(args[1:]))
	for _, file := range args[1:] {
		f := file
		wg.Go(func() {
			count, err := helper.ParseFile(f)

			ch <- helper.Result{
				Count: count,
				Err:   err,
			}
		})
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {

		if result.Err != nil {
			fmt.Println(result.Err)
		} else {
			fmt.Println("Filename ", result.Count.FileName)
			fmt.Println("Number of Lines ", result.Count.NumLines)
			fmt.Println("Info ", result.Count.NumInfo)
			fmt.Println("Warn ", result.Count.NumWarn)
			fmt.Println("Error ", result.Count.NumErr)
			fmt.Println("Debug ", result.Count.NumDebug)
			fmt.Println("Unknown ", result.Count.UnknownLines)

			totalCount.TotalNumInfo += result.Count.NumInfo
			totalCount.TotalNumWarn += result.Count.NumWarn
			totalCount.TotalNumErr += result.Count.NumErr
			totalCount.TotalNumDebug += result.Count.NumDebug
			totalCount.TotalNumLines += result.Count.NumLines
			totalCount.TotalUnknownLines += result.Count.UnknownLines
		}

	}

	// 	if err == nil {
	// 		fmt.Println("Filename ", count.FileName)
	// 		fmt.Println("Number of Lines ", count.NumLines)
	// 		fmt.Println("Info ", count.NumInfo)
	// 		fmt.Println("Warn ", count.NumWarn)
	// 		fmt.Println("Error ", count.NumErr)
	// 		fmt.Println("Debug ", count.NumDebug)
	// 		fmt.Println("Unknown ", count.UnknownLines)

	// 		totalCount.TotalNumInfo += count.NumInfo
	// 		totalCount.TotalNumWarn += count.NumWarn
	// 		totalCount.TotalNumErr += count.NumErr
	// 		totalCount.TotalNumDebug += count.NumDebug
	// 		totalCount.TotalNumLines += count.NumLines
	// 		totalCount.TotalUnknownLines += count.UnknownLines

	// 	} else {
	// 		fmt.Println("Something Happened ", err)
	// 	}
	// }

	// //fmt.Println("Filename ", fileName)
	fmt.Println("Total Number Of  Lines ", totalCount.TotalNumLines)
	fmt.Println("Total Info Count", totalCount.TotalNumInfo)
	fmt.Println("Total Warn Count", totalCount.TotalNumWarn)
	fmt.Println("Total Error Count", totalCount.TotalNumErr)
	fmt.Println("Total Debug Count", totalCount.TotalNumDebug)
	fmt.Println("Total Unknown Count", totalCount.TotalUnknownLines)
	//wg.Wait()
}
