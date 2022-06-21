package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
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

var instructions = make(map[string]instruction)

func read_instructions() {
	csvFile, err := os.Open("8080.csv")
	if err != nil {
		fmt.Println(err)
	}
	// Close when finished
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		opcode := line[0]
		size, _ := strconv.ParseInt(line[2], 16, 64)
		mnemonic, flags, function := line[1], line[3], line[4]
		instructions[opcode[2:]] = instruction{opcode, mnemonic, size, flags, function}
	}
	// Ignore CSV header, a bit hacky, but it'll do for now
	// instructions = instructions[1:]
}

// codebuffer should hold valid 8080 assembly code.
// pc is the current offset into the code
// returns the number of bytes of the op
func disassemble_8080op(codebuffer string, pc int) int64 {
	var code string = codebuffer[pc : pc+2]
	var instruction instruction = instructions[code]
	fmt.Printf("0x%X: %q | %q \n", pc, instruction.mnemonic, instruction.function)
	return instruction.size
}

// Takes the invader.bin file
// and turns it into a string
func read_rom(path string) string {
	invaders, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	n := 0
	for _, num := range invaders {
		if num >= 48 {
			invaders[n] = num
			n++
		}
	}
	return string(invaders[:n])
}

// Takes a path to the invader.bin file
// and prints its instructions
func disassemble_rom(path string) {
	var invaders_rom string = read_rom(path)
	program_counter := 0
	for program_counter < len(invaders_rom)-1 {
		size := disassemble_8080op(invaders_rom, program_counter)
		program_counter += (int(size) * 2)
	}

}
func main() {
	read_instructions()
	disassemble_rom("invaders.bin")
}
