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

func (self *CPU) Step() *CPU {
	registers := make(map[string]int, len(self.registers))
	for k, v := range self.registers {
		registers[k] = v
	}
	return &CPU{
		registers: registers,
		pc:        self.pc + 1,
		f:         self.f,
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
	Execute(*CPU) *CPU
}

type SoundInstruction struct {
	frequency string
}

func (self *SoundInstruction) Execute(cpu *CPU) *CPU {
	next := cpu.Step()
	next.f = next.Value(self.frequency)
	return next
}

type SetInstruction struct {
	register, val string
}

func (self *SetInstruction) Execute(cpu *CPU) *CPU {
	next := cpu.Step()
	next.registers[self.register] = next.Value(self.val)
	return next
}

type AddInstruction struct {
	register, val string
}

func (self *AddInstruction) Execute(cpu *CPU) *CPU {
	next := cpu.Step()
	next.registers[self.register] = next.Get(self.register) + next.Value(self.val)
	return next
}

type MultiplyInstruction struct {
	register, val string
}

func (self *MultiplyInstruction) Execute(cpu *CPU) *CPU {
	next := cpu.Step()
	next.registers[self.register] = next.Get(self.register) * next.Value(self.val)
	return next
}

type ModuloInstruction struct {
	register, val string
}

func (self *ModuloInstruction) Execute(cpu *CPU) *CPU {
	next := cpu.Step()
	next.registers[self.register] = next.Get(self.register) % next.Value(self.val)
	return next
}

type RecoverInstruction struct {
	val string
}

func (self *RecoverInstruction) Execute(cpu *CPU) *CPU {
	next := cpu.Step()
	val := next.Value(self.val)
	if val != 0 {
		fmt.Printf("RECOVERED: %d\n", next.f)
	}
	return next
}

type BranchInstruction struct {
	register, offset string
}

func (self *BranchInstruction) Execute(cpu *CPU) *CPU {
	next := cpu.Step()
	jump := next.Value(self.offset) - 1
	if next.Get(self.register) > 0 {
		next.pc += jump
	}
	return next
}

func readInstruction(line string) (Instruction, error) {
	fields := strings.Fields(line)
	switch fields[0] {
	case "snd":
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
		return &RecoverInstruction{fields[1]}, nil
	case "jgz":
		return &BranchInstruction{fields[1], fields[2]}, nil
	default:
		return nil, fmt.Errorf("Could not parse %s", line)
	}
}

func readInstructions(lines []string) ([]Instruction, error) {
	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		instr, err := readInstruction(line)
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

	instructions, err := readInstructions(
		strings.FieldsFunc(
			string(input),
			func(r rune) bool { return r == '\n' },
		),
	)
	if err != nil {
		panic(err)
	}

	cpu := NewCPU()
	for {
		if cpu.pc < 0 || cpu.pc > len(instructions) {
			break
		}
		instr := instructions[cpu.pc]
		cpu = instr.Execute(cpu)
	}
}
