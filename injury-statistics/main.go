package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func main() {

	file, err := os.Open("file-1.csv")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Accept variable number of fields per line

	// remove first row
	reader.Read()

	var wg sync.WaitGroup
	rowChan := make(chan []string) // Channel to send rows to

	go func() {
		for {
			row, err := reader.Read()
			if err != nil {
				close(rowChan) // Close the channel when we reach the end of the file
				break
			}
			rowChan <- row // Send the row to the channel
		}
	}()
	tbl := prepareTable()
	for row := range rowChan {
		wg.Add(1)
		go func(row []string) {
			defer wg.Done()
			tbl.AddRow(row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11])
		}(row)
	}

	wg.Wait()
	tbl.Print()

}

func prepareTable() table.Table {
	headerFmt := color.New(color.FgBlue, color.Underline, color.Bold).SprintfFunc()
	columnFmt := color.New(color.FgHiRed).SprintfFunc()

	tbl := table.New("Reference", "Period", "Type", "DataValue", "Lower_CI", "Units", "Indicator", "Cause", "Validation", "Population", "Age", "Severity")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(2)

	return tbl
}
