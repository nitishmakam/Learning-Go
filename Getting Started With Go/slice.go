package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var x = make([]int, 0, 3)

	for {
		fmt.Print("Enter a number: ")
		var input string
		fmt.Scan(&input)

		if input == "X" {
			break
		}
		intVar, _ := strconv.Atoi(input)
		x = append(x, intVar)

		sort.Ints(x)
		fmt.Println(x)

	}
}
