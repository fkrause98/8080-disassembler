package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"io/ioutil"
)

type instruction struct {
	opcode   string
	mnemonic string
	size     int64
	flags    string
	function string
}

var instructions = []instruction{}

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
		instructions = append(instructions, instruction{opcode, mnemonic, size, flags, function})
	}
	// Ignore CSV header, a bit hacky, but it'll do for now
	instructions = instructions[1:]
}

// codebuffer should hold valid 8080 assembly code.
// pc is the current offset into the code
// returns the number of bytes of the op
func disassemble_8080op(codebuffer []byte, pc int) (int64){
	var code byte = codebuffer[pc]
	var instruction instruction =  instructions[code]
	fmt.Printf("0x%x: %q | %q \n", pc, instruction.mnemonic, instruction.function)
	return instruction.size
}

func main() {
	read_instructions()
	invaders, err := ioutil.ReadFile("invaders.bin")
	if err != nil {
		fmt.Println(err)
	}

	for indx, num := range invaders {
		invaders[indx] = (num % 48)
	}
	for program_counter, _ := range invaders {
		size := disassemble_8080op(invaders, program_counter)
		program_counter += (int(size)*2)
	}
}
