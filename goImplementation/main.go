package main

import (
	"fmt"
	"time"
)

// main is the entry point of the application
// Author: Daniel palacios Moreno & Sofia Nicolle Ariza Goenaga
func main() {
	startTime := time.Now()

	hblv := NewHostBlackListsValidator()
	n := 202
	blackListOccurrences := hblv.CheckHost("202.24.34.55", n)
	
	fmt.Println("The host was found in the following blacklists:", blackListOccurrences)

	executionTime := time.Since(startTime).Milliseconds()
	fmt.Printf("Execution time: %d milliseconds\n", executionTime)
}
