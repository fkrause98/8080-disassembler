package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type instruction struct {
	opcode   string
	mnemonic string
	size     int64
	flags    string
	function string
}

func read_instructions() {
	csvFile, err := os.Open("8080.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("CSV File Opened")
	// Close when finished
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	instructions := []instruction{}

	for _, line := range csvLines {
		opcode := line[0]
		size, _ := strconv.ParseInt(line[2], 16, 64)
		mnemonic, flags, function := line[1], line[3], line[4]
		instructions = append(instructions, instruction{opcode, mnemonic, size, flags, function})
	}
	for _, instruction := range instructions {
		fmt.Println(instruction)
	}
}
func main() {
	read_instructions()
}
