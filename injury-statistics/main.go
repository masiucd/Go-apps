package main

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

var testData = [][]string{
	{"1", "Alice", "100", "2018-01-01"},
	{"2", "Bob", "95", "2018-01-02"},
	{"3", "Charlie", "90", "2018-01-03"},
	{"4", "David", "85", "2018-01-04"},
}

func main() {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Name", "Score", "Added")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, row := range testData {
		tbl.AddRow(row[0], row[1], row[2], row[3])
	}
	tbl.Print()

	// file, err := os.Open("file-1.csv")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// defer file.Close()

	// reader := csv.NewReader(file)
	// reader.FieldsPerRecord = -1 // Accept variable number of fields per line
	// data, err := reader.ReadAll()
	// if err != nil {
	// 	panic(err)
	// }

	// var titles []string
	// for _, title := range data[0] {
	// 	titles = append(titles, title)
	// }

	// var result []map[string]string
	// fmt.Println(len(titles))
	// fmt.Println(len(data[2]))
	// for i := range data[1:] {
	// 	// item := make(map[string]string)
	// 	for _, col := range titles {
	// 		fmt.Println(titles[i], col)
	// 	}
	// 	// result = append(result, item)
	// }

	// fmt.Println(result)
	// var xs []string
	// for _, row := range data {

	// 	for _, col := range row {
	// 		fmt.Println(col)
	// 	}

	// }

	// fmt.Println(injuryItems)
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
