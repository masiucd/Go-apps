package lib

import (
	"encoding/csv"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

// PrepareTable prepares the table with headers and column formatting
func PrepareTable() table.Table {
	headerFmt := color.New(color.FgBlue, color.Underline, color.Bold).SprintfFunc()
	columnFmt := color.New(color.FgCyan, color.BlinkRapid, color.Bold).SprintfFunc()
	tbl := table.New("Ref", "Period", "Type", "DataValue", "L_CI", "U_CI", "Units", "Indicator", "Cause", "Validation", "Population", "Age").WithPadding(2)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	return tbl
}

// removeCSVHeaders removes the headers from the CSV file
func RemoveCSVHeaders(reader *csv.Reader) {
	// remove first row of headers
	reader.Read()
}
