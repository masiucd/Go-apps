package printfile

import (
	"fmt"
	"strings"
)

const NEW_LINE = 10

// SimplePrint prints the data to the console
func SimplePrint(data []byte) {
	var output string
	for _, v := range data {
		if v == NEW_LINE {
			// Print a NEW_LINE character
			output += "\n"
		} else {
			// Print the character
			output += string(v)
		}
	}
	fmt.Println(output)
}

// PrintEven prints the even lines of the data to the console
func PrintEven(data []byte) {
	var builder strings.Builder
	lineCount := 0
	for _, v := range data {
		if v == '\n' {
			// If it's an even line, print it
			if lineCount%2 == 0 {
				fmt.Println(builder.String())
			}
			// Reset the builder and increase the line count
			builder.Reset()
			lineCount++
		} else {
			// Add the character to the current line
			builder.WriteByte(v)
		}
	}
	// Print the last line if it's an even line and the data doesn't end with a NEW_LINE
	if lineCount%2 == 0 {
		fmt.Println(builder.String())
	}
}
