package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type CPU struct {
	registers map[string]int
	pc, f     int
}

func NewCPU() *CPU {
	return &CPU{
		registers: map[string]int{},
		pc:        0,
		f:         -1,
	}
}

func (self *CPU) Get(register string) int {
	if v, set := self.registers[register]; set {
		return v
	}
	self.registers[register] = 0
	return 0
}

func (self *CPU) Value(valueOrRegister string) int {
	val, err := strconv.Atoi(valueOrRegister)
	if err != nil {
		return self.Get(valueOrRegister)
	} else {
		return val
	}
}

type Instruction interface {
	Execute(*CPU)
}

type SoundInstruction struct {
	frequency string
}

func (self *SoundInstruction) Execute(cpu *CPU) {
	cpu.pc++
	cpu.f = cpu.Value(self.frequency)
}

type SetInstruction struct {
	register, val string
}

func (self *SetInstruction) Execute(cpu *CPU) {
	cpu.pc++
	cpu.registers[self.register] = cpu.Value(self.val)
}

type AddInstruction struct {
	register, val string
}

func (self *AddInstruction) Execute(cpu *CPU) {
	cpu.pc++
	cpu.registers[self.register] = cpu.Get(self.register) + cpu.Value(self.val)
}

type MultiplyInstruction struct {
	register, val string
}

func (self *MultiplyInstruction) Execute(cpu *CPU) {
	cpu.pc++
	cpu.registers[self.register] = cpu.Get(self.register) * cpu.Value(self.val)
}

type ModuloInstruction struct {
	register, val string
}

func (self *ModuloInstruction) Execute(cpu *CPU) {
	cpu.pc++
	cpu.registers[self.register] = cpu.Get(self.register) % cpu.Value(self.val)
}

type RecoverInstruction struct {
	val string
}

func (self *RecoverInstruction) Execute(cpu *CPU) {
	cpu.pc++
	val := cpu.Value(self.val)
	if val != 0 {
		panic(fmt.Sprintf("RECOVERED: %d\n", cpu.f))
	}
}

type BranchInstruction struct {
	register, offset string
}

func (self *BranchInstruction) Execute(cpu *CPU) {
	jump := cpu.Value(self.offset)
	if cpu.Get(self.register) > 0 {
		cpu.pc += jump
	} else {
		cpu.pc++
	}
}

type SendInstruction struct {
	val string
}

func (self *SendInstruction) Execute(cpu *CPU) {
}

type ReceiveInstruction struct {
	register string
}

func (self *ReceiveInstruction) Execute(cpu *CPU) {
}

func readInstruction(line string, version int) (Instruction, error) {
	fields := strings.Fields(line)
	switch fields[0] {
	case "snd":
		if version == 2 {
			return &SendInstruction{fields[1]}, nil
		}
		return &SoundInstruction{fields[1]}, nil
	case "set":
		return &SetInstruction{fields[1], fields[2]}, nil
	case "add":
		return &AddInstruction{fields[1], fields[2]}, nil
	case "mul":
		return &MultiplyInstruction{fields[1], fields[2]}, nil
	case "mod":
		return &ModuloInstruction{fields[1], fields[2]}, nil
	case "rcv":
		if version == 2 {
			return &ReceiveInstruction{fields[1]}, nil
		}
		return &RecoverInstruction{fields[1]}, nil
	case "jgz":
		return &BranchInstruction{fields[1], fields[2]}, nil
	default:
		return nil, fmt.Errorf("Could not parse %s", line)
	}
}

func readInstructions(lines []string, version int) ([]Instruction, error) {
	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		instr, err := readInstruction(line, version)
		if err != nil {
			return instructions, err
		}
		instructions[i] = instr
	}
	return instructions, nil
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	instructionLines := strings.FieldsFunc(string(input), func(r rune) bool { return r == '\n' })

	instructions, err := readInstructions(instructionLines, 1)
	if err != nil {
		panic(err)
	}

	cpu := NewCPU()
	for {
		if cpu.pc < 0 || cpu.pc > len(instructions) {
			break
		}
		instructions[cpu.pc].Execute(cpu)
	}
}
