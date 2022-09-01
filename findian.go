package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Printf("Enter a string: ")
	fmt.Scan(&input)

	if (strings.HasPrefix(input, "i") || strings.HasPrefix(input, "I")) && (strings.ContainsRune(input, 'a') || strings.ContainsRune(input, 'A')) && (strings.HasSuffix(input, "n") || strings.HasSuffix(input, "N")) {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
