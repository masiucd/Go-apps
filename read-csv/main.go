package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type row struct {
	orderNumber int
	pnr         string
	airline     string
	amount      float64
}

func main() {

	file, err := os.Open("app.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var rows []row
	records = records[1:]

	for _, rec := range records {
		rows = append(rows, row{
			orderNumber: fromStringToInt(rec[0]),
			pnr:         rec[1],
			airline:     rec[2],
			amount:      fromStringToFloat(rec[3]) / 2,
		})
	}

	printRows(rows)
}

func fromStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func fromStringToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

func printRows(rows []row) {
	for _, r := range rows {
		fmt.Println("Order Number: ", r.orderNumber)
		fmt.Println("PNR: ", r.pnr)
		fmt.Println("Airline: ", r.airline)
		fmt.Println("Amount: ", r.amount)
		fmt.Println("====================================")
	}
}
