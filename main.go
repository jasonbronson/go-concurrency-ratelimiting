package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	f, err := os.Open("./import.dat")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	var line string
	var lineCount int64
	var wg sync.WaitGroup
	guard := make(chan struct{}, 30)
	for {
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if lineCount == 0 {
			lineCount++
			continue
		}
		lineCount++
		r := csv.NewReader(strings.NewReader(string(line)))
		r.Comma = '|'
		records, err := r.ReadAll()
		if err != nil {
			panic(err)
		}
		for _, row := range records {
			//fmt.Println(row)
			guard <- struct{}{}
			wg.Add(1)
			go func(row []string, wg *sync.WaitGroup) {
				//inside here
				
				defer wg.Done()
				data := &Testing{
					FirstName:   row[0],
					LastName:    row[1],
					Phone:       row[2],
					LastUpdated: formatDate(row[3]),
				}
				insertRow(data)
				<-guard

			}(row, &wg)

		}
	}
	fmt.Println("Finished")
	wg.Wait()
}

func formatDate(dateTime string) time.Time {
	if len(dateTime) == 0 {
		return time.Time{}
	}
	dateTime = truncate(dateTime, " ")
	t, err := time.Parse("2006-01-02", dateTime)
	if err != nil {
		panic(err)
	}

	return t
}

func truncate(data string, field string) string {
	return strings.Trim(data, field)
}
