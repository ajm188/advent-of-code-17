package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type CPU struct {
	registers map[string]int
}

type Instruction struct {
}

func readInstructions(lines string) (*[]Instruction, error) {
	return []*Instruction{}, nil
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	newline := func(r rune) bool { return r == '\n' }
	instructions, err := readInstructions(strings.FieldsFunc(string(input), newline))
	if err != nil {
		panic(err)
	}
	fmt.Println(instructions)
}
