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
	totalCount := &helper.TotalCount{}
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
			helper.PrintSummary(result.Count)
			helper.UpdateTotalCount(totalCount, result.Count)
		}

	}

	helper.PrintTotalSummary(totalCount)

}
