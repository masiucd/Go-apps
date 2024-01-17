package main

import (
	"fmt"
	"log"
)

func main() {
	run()
}

func run() {
	inputAmount := getUserInput[float64]("Enter the amount of money you want to convert: ")

	fmt.Println("1. USD")
	fmt.Println("2. EUR")
	fmt.Println("3. PLN")
	currencyFrom := getUserInput[int]("Enter the currency you want to convert from: ")

	fmt.Println("1. USD")
	fmt.Println("2. EUR")
	fmt.Println("3. PLN")
	currencyTo := getUserInput[int]("Enter the currency you want to convert to: ")

	result := convert(inputAmount, currencyFrom, currencyTo)
	prettyPrint(result, inputAmount, currencyFrom, currencyTo)
}

var currencyMap = map[int]string{
	1: "USD",
	2: "EUR",
	3: "PLN",
}

func prettyPrint(result, inputAmount float64, currencyFrom, currencyTo int) {
	fmt.Printf("%v %v is %v %v \n", inputAmount, currencyMap[currencyFrom], result, currencyMap[currencyTo])
}

const (
	USD = iota + 1
	EUR
	PLN
)

func convert(amount float64, currencyFrom, currencyTo int) float64 {
	switch {
	case currencyFrom == USD && currencyTo == EUR:
		return fromUsdToEur(amount)
	case currencyFrom == EUR && currencyTo == USD:
		return fromEurToUsd(amount)
	case currencyFrom == PLN && currencyTo == EUR:
		return fromPlnToEur(amount)
	case currencyFrom == EUR && currencyTo == PLN:
		return fromEurToPln(amount)
	case currencyFrom == PLN && currencyTo == USD:
		return fromPlnToUsd(amount)
	case currencyFrom == USD && currencyTo == PLN:
		return fromUsdToPln(amount)
	}
	return amount

}

func fromUsdToEur(amount float64) float64 {
	return amount * 0.83
}

func fromEurToUsd(amount float64) float64 {
	return amount * 1.20
}

func fromPlnToEur(amount float64) float64 {
	return amount * 0.22
}

func fromEurToPln(amount float64) float64 {
	return amount * 4.53
}

func fromPlnToUsd(amount float64) float64 {
	return amount * 0.27
}

func fromUsdToPln(amount float64) float64 {
	return amount * 3.71
}

func getUserInput[T any](prompt string) T {
	fmt.Println(prompt)
	var input T
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}
	return input
}
