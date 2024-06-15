package main

import (
	"encoding/csv"
	"log"
	"os"
	"sync"
	"time"

	"go-apps.com/injury-statistics/lib"
)

func main() {

	file, err := os.Open("file-1.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Accept variable number of fields per line

	lib.RemoveCSVHeaders(reader)

	rowChan := make(chan []string) // Channel to send rows to
	start := time.Now()

	go readCSVRows(reader, rowChan)

	tbl := lib.PrepareTable()
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

// readCSVRows reads rows from a CSV reader and sends them to a channel.
// It continues reading until an error occurs or the end of the file is reached.
// Upon encountering an error or EOF, it closes the channel to signal completion.
func readCSVRows(reader *csv.Reader, rowChan chan []string) {
	for {
		row, err := reader.Read() // Attempt to read the next row from the CSV.
		if err != nil {
			close(rowChan) // Close the channel on error or EOF (End of line) to signal no more rows.
			break
		}
		rowChan <- row // Send the successfully read row to the channel.
	}
}
