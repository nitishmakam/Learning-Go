package main

import (
	"fmt"
	"os"
)

func main() {
	var length int
	fmt.Print("Enter the number of integers: ")
	fmt.Scan(&length)

	if length > 10 {
		fmt.Println("Failed! A maximum of 10 integers are accepted.")
		os.Exit(0)
	}

	slice := make([]int, 0, length)
	for i := 0; i < length; i++ {
		fmt.Printf("Enter integer for index %d: ", i)
		var input int
		fmt.Scan(&input)
		slice = append(slice, input)
	}

	BubbleSort(slice)

	fmt.Print("Sorted Order: \n")
	fmt.Println(slice)

}

func BubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func Swap(slice []int, index int) {
	orig := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = orig
}
