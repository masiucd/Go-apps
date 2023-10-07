package main

import "fmt"

func main() {
	var operator string
	var num1, num2 int
	fmt.Println("Enter the operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&num2)

	fmt.Println(operator, num1, num2)

	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Invalid operator")
	}
}
