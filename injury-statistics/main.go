package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func getFile(f string) (*os.File, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	return file, nil
}

type FileType int

const (
	FileOne FileType = 1
	FileTwo FileType = 2
)

func (f FileType) String() string {
	switch f {
	case FileOne:
		return "file-1.csv"
	case FileTwo:
		return "file-2.csv"
	default:
		return "Unknown"
	}
}

func main() {

	var fileNum FileType
	fmt.Println("You can choose file one or file two")
	fmt.Printf("\n")
	fmt.Println("Select 1 for file one and 2 for file two")
	fmt.Scan(&fileNum)
	fmt.Printf("\n")
	fmt.Println("You selected file ", fileNum)

	var file *os.File
	var err error

	switch fileNum {
	case FileOne:
		fmt.Println("You selected file one")
		file, err = getFile("file-1.csv")
	case FileTwo:
		fmt.Println("You selected file two")
		file, err = getFile("file-2.csv")
	default:
		fmt.Println("You selected an unknown file")
	}

	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Accept variable number of fields per line

	removeCSVHeaders(reader)

	for {
		row, err := reader.Read()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(row)
	}

	// tbl, err := prepareTable(fileNum)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	// rowChan := make(chan []string) // Channel to send rows to

	// go func() {
	// 	for {
	// 		row, err := reader.Read()
	// 		if err != nil {
	// 			close(rowChan) // Close the channel when we reach the end of the file
	// 			break
	// 		}
	// 		rowChan <- row // Send the row to the channel
	// 	}
	// }()

	// tbl := prepareTable()
	// var wg sync.WaitGroup
	// for row := range rowChan {
	// 	wg.Add(1)
	// 	go func(row []string) {
	// 		defer wg.Done()
	// 		tbl.AddRow(row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11])
	// 	}(row)
	// }

	// wg.Wait()
	// tbl.Print()

}

func prepareTable(file FileType) (*table.Table, error) {
	headerFmt := color.New(color.FgBlue, color.Underline, color.Bold).SprintfFunc()
	columnFmt := color.New(color.FgHiRed).SprintfFunc()

	switch file {
	case FileOne:
		headersForFileOne := []string{
			"Reference",
			"Period",
			"Type",
			"DataValue",
			"Lower_CI",
			"Upper_CI",
			"Units",
			"Indicator",
			"Cause",
			"Validation",
			"Population",
			"Age",
			"Severity",
		}
		var tbl table.Table
		for _, header := range headersForFileOne {
			tbl.AddRow(header)
		}
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(2)
		return &tbl, nil
	case FileTwo:
		headersFileTwo := []string{
			"Year",
			"Sex",
			"Age group (years) at date of injury",
			"Geographic region where injury occurred",
			"Employment status",
			"Occupation",
			"Injury/illness/disease group",
			"Type of injury/illness/disease",
			"Industry",
			"Industry subgroup",
			"Value",
			"Measure",
			"Status",
		}
		var tbl table.Table
		for _, header := range headersFileTwo {
			tbl.AddRow(header)
		}
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(2)
		return &tbl, nil
	default:
		return nil, fmt.Errorf("unknown file")

	}

}

// removeCSVHeaders removes the headers from the CSV file
func removeCSVHeaders(reader *csv.Reader) {
	// remove first row of headers
	reader.Read()
}
