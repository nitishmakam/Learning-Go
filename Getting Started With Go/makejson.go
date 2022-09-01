package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	var address string
	fmt.Print("Enter your address: ")
	scanner.Scan()
	address = scanner.Text()

	m["name"] = name
	m["address"] = address

	jsonString, _ := json.Marshal(m)
	fmt.Println(string(jsonString))
}
