package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type CPU struct {
	registers map[string]int
}

func NewCPU() *CPU {
	return &CPU{
		registers: map[string]int{},
	}
}

func (cpu *CPU) Get(reg string) int {
	if v, ok := cpu.registers[reg]; ok {
		return v
	} else {
		cpu.registers[reg] = 0
		return 0
	}
}

func (cpu *CPU) Set(reg string, val int) {
	cpu.registers[reg] = val
}

func (cpu *CPU) LargestValue() int {
	max := math.Inf(-1)
	for _, v := range cpu.registers {
		max = math.Max(max, float64(v))
	}
	return int(max)
}

type Condition func(*CPU) bool
type Operation func(*CPU)

type Instruction struct {
	Condition
	Operation
}

func createOperation(operationName string) (op func(int, int) int) {
	switch operationName {
	case "inc":
		op = func(a, b int) int { return a + b }
	case "dec":
		op = func(a, b int) int { return a - b }
	}
	return
}

func createCondition(conditionName string) (cond func(int, int) bool) {
	switch conditionName {
	case "==":
		cond = func(a, b int) bool { return a == b }
	case "!=":
		cond = func(a, b int) bool { return a != b }
	case ">":
		cond = func(a, b int) bool { return a > b }
	case ">=":
		cond = func(a, b int) bool { return a >= b }
	case "<":
		cond = func(a, b int) bool { return a < b }
	case "<=":
		cond = func(a, b int) bool { return a <= b }
	}
	return
}

func NewInstructionFromFields(fields []string) (*Instruction, error) {
	opFields := fields[0:3]
	condFields := fields[4:len(fields)]
	opRegister := opFields[0]
	opValue, err := strconv.Atoi(opFields[2])
	if err != nil {
		return nil, err
	}
	condRegister := condFields[0]
	condValue, err := strconv.Atoi(condFields[2])
	if err != nil {
		return nil, err
	}

	op := createOperation(opFields[1])
	operation := func(cpu *CPU) {
		reg := cpu.Get(opRegister)
		cpu.Set(opRegister, op(reg, opValue))
	}

	cond := createCondition(condFields[1])
	condition := func(cpu *CPU) bool {
		return cond(cpu.Get(condRegister), condValue)
	}

	inst := &Instruction{
		Condition: condition,
		Operation: operation,
	}
	return inst, nil
}

func readInstructions(lines []string) ([]*Instruction, error) {
	instructions := make([]*Instruction, 0, len(lines))
	for _, line := range lines {
		inst, err := NewInstructionFromFields(strings.Fields(line))
		if err != nil {
			return instructions, err
		}
		instructions = append(instructions, inst)
	}
	return instructions, nil
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
	cpu := NewCPU()
	for _, inst := range instructions {
		if inst.Condition(cpu) {
			inst.Operation(cpu)
		}
	}
	fmt.Println(cpu.LargestValue())
}
