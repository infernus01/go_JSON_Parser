package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	var input string

	fmt.Println("Enter a JSON object:")
	fmt.Scan(&input)

	var parsedData interface{}
	err := json.Unmarshal([]byte(input), &parsedData)
	if err != nil {
	    log.Println();
		fmt.Println("Invalid JSON:", err)
		os.Exit(1)
	}

	fmt.Println("Valid JSON")
	os.Exit(0)
}
