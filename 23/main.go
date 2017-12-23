package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type InstructionType int

const (
	SET InstructionType = iota
	SUB
	MUL
	JNZ
)

type CPU struct {
	registers map[string]int
	debug     map[InstructionType]int
	pc        int
}

func NewCPU() *CPU {
	return &CPU{
		registers: map[string]int{},
		debug:     map[InstructionType]int{},
		pc:        0,
	}
}

func (self *CPU) Get(value string) int {
	v, err := strconv.Atoi(value)
	if err != nil {
		v, ok := self.registers[value]
		if ok {
			return v
		} else {
			return 0
		}
	}
	return v
}

func (self *CPU) Set(register string, value int) {
	self.registers[register] = value
}

func (self *CPU) Debug(instr Instruction) {
	instrCount, ok := self.debug[instr.Type()]
	if !ok {
		instrCount = 0
	}
	self.debug[instr.Type()] = instrCount + 1
}

type Instruction interface {
	Execute(*CPU)
	Type() InstructionType
}

type SetInstruction struct {
	register, value string
}

func (self *SetInstruction) Execute(cpu *CPU) {
	cpu.Set(self.register, cpu.Get(self.value))
	cpu.pc++
}

func (self *SetInstruction) Type() InstructionType {
	return SET
}

type SubInstruction struct {
	register, value string
}

func (self *SubInstruction) Execute(cpu *CPU) {
	value := cpu.Get(self.register) - cpu.Get(self.value)
	cpu.Set(self.register, value)
	cpu.pc++
}

func (self *SubInstruction) Type() InstructionType {
	return SUB
}

type MulInstruction struct {
	register, value string
}

func (self *MulInstruction) Execute(cpu *CPU) {
	value := cpu.Get(self.register) * cpu.Get(self.value)
	cpu.Set(self.register, value)
	cpu.pc++
}

func (self *MulInstruction) Type() InstructionType {
	return MUL
}

type JnzInstruction struct {
	value, offset string
}

func (self *JnzInstruction) Execute(cpu *CPU) {
	value := cpu.Get(self.value)
	if value != 0 {
		cpu.pc += cpu.Get(self.offset)
	} else {
		cpu.pc++
	}
}

func (self *JnzInstruction) Type() InstructionType {
	return JNZ
}

func readInstruction(text string) (Instruction, error) {
	fields := strings.Fields(text)
	switch fields[0] {
	case "set":
		return &SetInstruction{fields[1], fields[2]}, nil
	case "sub":
		return &SubInstruction{fields[1], fields[2]}, nil
	case "mul":
		return &MulInstruction{fields[1], fields[2]}, nil
	case "jnz":
		return &JnzInstruction{fields[1], fields[2]}, nil
	default:
		return nil, fmt.Errorf("Unknown instruction %v", text)
	}
}

func readInstructions(lines [][]byte) ([]Instruction, error) {
	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		instruction, err := readInstruction(string(line))
		if err != nil {
			return instructions, err
		}
		instructions[i] = instruction
	}
	return instructions, nil
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines := bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' })
	instructions, err := readInstructions(lines)
	if err != nil {
		panic(err)
	}

	cpu := NewCPU()
	for cpu.pc >= 0 && cpu.pc < len(instructions) {
		instruction := instructions[cpu.pc]
		instruction.Execute(cpu)
		cpu.Debug(instruction)
	}

	fmt.Println(cpu.debug[MUL])
}
