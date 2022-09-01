package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filename string
	fmt.Print("Enter filename: ")
	fmt.Scan(&filename)

	fd, err := os.Open(filename)
	if err != nil {
		fmt.Print("Error occurred while opening file!")
		os.Exit(0)
	}

	type name struct {
		fname string
		lname string
	}

	names := make([]name, 0, 100)

	fileScanner := bufio.NewScanner(fd)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		flname := strings.Split(line, " ")
		p := name{flname[0], flname[1]}
		names = append(names, p)
	}

	for _, entry := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", entry.fname, entry.lname)
	}

	fd.Close()

}
