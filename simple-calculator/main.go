package main

import "fmt"

func main() {
	var operator string = getOperator()
	var num1, num2 int = getNumbers()
	var version int = getCalcVersion()
	switch version {
	case 1:
		calculateV1(operator, num1, num2)
	case 2:
		calculateV2(operator, num1, num2)
	default:
		fmt.Println("Invalid version")
	}
}

func getCalcVersion() int {
	var version int
	fmt.Println("select calculate version either 1 or 2")
	fmt.Scanln(&version)
	return version
}

func getOperator() string {
	var operator string
	fmt.Println("Enter the operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	return operator
}

func getNumbers() (int, int) {
	var num1, num2 int
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&num2)
	return num1, num2
}

var (
	add = func(a, b int) int {
		return a + b
	}
	subtract = func(a, b int) int {
		return a - b
	}
	multiply = func(a, b int) int {
		return a * b
	}
	divide = func(a, b int) int {
		if b == 0 {
			return 0
		}
		return a / b
	}
)

type calcFunc func(int, int) int

var calcMap = map[string]calcFunc{
	"+": add,
	"-": subtract,
	"*": multiply,
	"/": divide,
}

func calculateV2(operator string, num1 int, num2 int) {
	calcFunc, ok := calcMap[operator]
	if !ok {
		fmt.Println("Invalid operator")
		return
	}
	fmt.Println(calcFunc(num1, num2))
}

func calculateV1(operator string, num1 int, num2 int) {
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
