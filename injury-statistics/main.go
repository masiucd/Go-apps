package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func main() {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("SeriesReference", "Period", "Type", "DataValue", "Lower_CI", "Units", "Indicator", "Cause", "Validation", "Population", "Age", "Severity")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	file, err := os.Open("file-1.csv")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Accept variable number of fields per line

	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, row := range data[1:] {

		tbl.AddRow(row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11])
	}
	tbl.Print()

}

// type InjuryItem struct {
// 	SeriesReference string `csv:Series_reference`
// 	Period          string `csv:Period`
// 	Type            string `csv:Type`
// 	DataValue       string `csv:Data_value`
// 	Lower_CI        string `csv:Lower_CI`
// 	Units           string `csv:Units`
// 	Indicator       string `csv:Indicator`
// 	Cause           string `csv:Cause`
// 	Validation      string `csv:Validation`
// 	Population      string `csv:Population`
// 	Age             string `csv:Age`
// 	Severity        string `csv:Severity`
// }
