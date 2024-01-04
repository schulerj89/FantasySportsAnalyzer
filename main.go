package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	filePath := flag.String("file", "", "Path to the DraftKings export file")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please specify a file path using the -file flag.")
		os.Exit(1)
	}

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	players, err := ParseCSV(reader)
	if err != nil {
		fmt.Printf("Error parsing CSV file: %v\n", err)
		os.Exit(1)
	}

	ProcessData(players)
}
