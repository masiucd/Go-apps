package main

import (
	"encoding/csv"
	"log"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func main() {

	file, err := os.Open("file-1.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Accept variable number of fields per line

	removeCSVHeaders(reader)

	rowChan := make(chan []string) // Channel to send rows to
	start := time.Now()
	// Read the file in a separate goroutine
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
	var wg sync.WaitGroup
	for row := range rowChan {
		wg.Add(1)
		go func(row []string) {
			defer wg.Done()
			tbl.AddRow(row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11])

		}(row)
	}

	end := time.Since(start)
	wg.Wait()
	tbl.Print()
	log.Printf("Time taken: %v", end)

}

func prepareTable() table.Table {
	headerFmt := color.New(color.FgBlue, color.Underline, color.Bold).SprintfFunc()
	columnFmt := color.New(color.FgCyan, color.BlinkRapid, color.Bold).SprintfFunc()
	tbl := table.New("Ref", "Period", "Type", "DataValue", "L_CI", "U_CI", "Units", "Indicator", "Cause", "Validation", "Population", "Age").WithPadding(2)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	return tbl
}

// removeCSVHeaders removes the headers from the CSV file
func removeCSVHeaders(reader *csv.Reader) {
	// remove first row of headers
	reader.Read()
}
