package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func subSort(wg *sync.WaitGroup, arr []int) {
	fmt.Print("Go Routine: Sorting Sub-Array: ", arr, "\n")

	sort.Ints(arr)
	wg.Done()
}

func main() {

	br := bufio.NewReader(os.Stdin)
	fmt.Print("Enter integers to sort that are separated by space: ")
	a, _, _ := br.ReadLine()
	stringArray := strings.Split(string(a), " ")
	var intArray []int

	for _, s := range stringArray {
		n, _ := strconv.Atoi(s)
		intArray = append(intArray, n)
	}

	segmentSize := len(intArray) / 4

	var wg sync.WaitGroup
	for c := 0; c < 4; c++ {
		wg.Add(1)
		if c != 3 {
			go subSort(&wg, intArray[c*segmentSize:(c+1)*segmentSize])
		} else {
			go subSort(&wg, intArray[c*segmentSize:])
		}
	}

	wg.Wait()

	var sorted []int
	c1 := intArray[:segmentSize]
	c2 := intArray[segmentSize : 2*segmentSize]
	c3 := intArray[2*segmentSize : 3*segmentSize]
	c4 := intArray[3*segmentSize:]

	for k := 0; k < len(intArray); k++ {
		i := 0
		j := 0

		if len(c1) != 0 {
			i = c1[0]
			j = 1
		}

		if len(c2) != 0 {
			if j == 0 {
				i = c2[0]
				j = 2
			} else {
				if c2[0] < i {
					i = c2[0]
					j = 2
				}
			}
		}

		if len(c3) != 0 {
			if j == 0 {
				i = c3[0]
				j = 3
			} else {
				if c3[0] < i {
					i = c3[0]
					j = 3
				}
			}
		}

		if len(c4) != 0 {
			if j == 0 {
				i = c4[0]
				j = 4
			} else {
				if c4[0] < i {
					i = c4[0]
					j = 4
				}
			}
		}
		sorted = append(sorted, i)
		switch j {
		case 1:
			c1 = c1[1:]
		case 2:
			c2 = c2[1:]
		case 3:
			c3 = c3[1:]
		case 4:
			c4 = c4[1:]
		}
	}

	fmt.Println("Sorted Array: ", sorted)

}
