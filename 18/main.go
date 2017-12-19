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
	Execute(*CPU) error
}

type SoundInstruction struct {
	frequency string
}

func (self *SoundInstruction) Execute(cpu *CPU) error {
	cpu.pc++
	cpu.f = cpu.Value(self.frequency)
	return nil
}

type SetInstruction struct {
	register, val string
}

func (self *SetInstruction) Execute(cpu *CPU) error {
	cpu.pc++
	cpu.registers[self.register] = cpu.Value(self.val)
	return nil
}

type AddInstruction struct {
	register, val string
}

func (self *AddInstruction) Execute(cpu *CPU) error {
	cpu.pc++
	cpu.registers[self.register] = cpu.Get(self.register) + cpu.Value(self.val)
	return nil
}

type MultiplyInstruction struct {
	register, val string
}

func (self *MultiplyInstruction) Execute(cpu *CPU) error {
	cpu.pc++
	cpu.registers[self.register] = cpu.Get(self.register) * cpu.Value(self.val)
	return nil
}

type ModuloInstruction struct {
	register, val string
}

func (self *ModuloInstruction) Execute(cpu *CPU) error {
	cpu.pc++
	cpu.registers[self.register] = cpu.Get(self.register) % cpu.Value(self.val)
	return nil
}

type RecoverInstruction struct {
	val string
}

func (self *RecoverInstruction) Execute(cpu *CPU) error {
	cpu.pc++
	val := cpu.Value(self.val)
	if val != 0 {
		return fmt.Errorf("RECOVERED: %d\n", cpu.f)
	}
	return nil
}

type BranchInstruction struct {
	register, offset string
}

func (self *BranchInstruction) Execute(cpu *CPU) error {
	jump := cpu.Value(self.offset)
	if cpu.Get(self.register) > 0 {
		cpu.pc += jump
	} else {
		cpu.pc++
	}
	return nil
}

type SendInstruction struct {
	val string
}

func (self *SendInstruction) Execute(cpu *CPU) error {
	return nil
}

type ReceiveInstruction struct {
	register string
}

func (self *ReceiveInstruction) Execute(cpu *CPU) error {
	return nil
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
		if err := instructions[cpu.pc].Execute(cpu); err != nil {
			fmt.Println(err)
			break
		}
	}
}
