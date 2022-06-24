package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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
}

// codebuffer should hold valid 8080 assembly code.
// pc is the current offset into the code
// returns the number of bytes of the op
func disassemble_8080op(codebuffer string, pc int) (size int64) {
	var code string = codebuffer[pc : pc+2]
	var instruction instruction = instructions[code]
	var disassembly string
	size = instruction.size
	if size > 1 {
		var arguments string = codebuffer[pc+2 : int64(pc)+(instruction.size*2)]
		disassembly = replace_arguments(instruction.mnemonic, arguments)
		fmt.Printf("0x%X: %s \n", pc, disassembly)
		return
	} else {
		disassembly = fmt.Sprintf("0x%X: %s \n", pc, instruction.mnemonic)
		fmt.Printf(disassembly)
		return
	}
}

// Takes D16 and D8 or an adr from the given mnemonic and
// replaces it with an argument. Eg:
// If I have:
// JMP adr
// This will turn into
// JMP 18d4
func replace_arguments(mnemonic string, arguments string) string {
	var has_word_argument bool = strings.Contains(mnemonic, "D16")
	var has_address_argument bool = strings.Contains(mnemonic, "adr")
	var has_byte_argument bool = strings.Contains(mnemonic, "D8")
	// Remember that Intel is little endian
	if has_word_argument {
		return strings.Replace(mnemonic, "D16", string([]byte{arguments[2], arguments[3], arguments[0], arguments[1]}[:]), -1)
	}
	if has_address_argument {
		return strings.Replace(mnemonic, "adr", string([]byte{arguments[2], arguments[3], arguments[0], arguments[1]}[:]), -1)
	}
	if has_byte_argument {
		return strings.Replace(mnemonic, "D8", string([]byte{arguments[0], arguments[1]}[:]), -1)
	}
	return "ERROR REPLACING"
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
