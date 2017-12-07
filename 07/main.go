package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	Name string
	Weight int
	SupportedPrograms []string
}

func NewProgramFromLine(line string) (*Program, error) {
	fields := strings.Fields(line)
	name := fields[0]
	weightField := fields[1]
	weight, err := strconv.Atoi(weightField[1:len(weightField)-1])
	if err != nil {
		return nil, err
	}

	program := &Program{
		Name: name,
		Weight: weight,
		SupportedPrograms: make([]string, 0, len(fields)),
	}
	if len(fields) > 2 {
		for _, field := range fields[3:len(fields)] {
			if strings.HasSuffix(field, ",") {
				field = field[0:len(field)-1]
			}
			program.SupportedPrograms = append(program.SupportedPrograms, field)
		}
	}

	return program, nil
}

func (p *Program) IsBaseProgram() bool {
	return len(p.SupportedPrograms) == 0
}

func (p *Program) IsBalanced(programs map[string]*Program) bool {
	if p.IsBaseProgram() {
		return true
	}

	weight := -1
	for _, support := range p.SupportedPrograms {
		program, _ := programs[support]
		supportWeight := program.TotalWeight(programs)
		if weight == -1 {
			weight = supportWeight
		} else if weight != supportWeight {
			return false
		}
	}
	return true
}

func (p *Program) TotalWeight(programs map[string]*Program) int {
	weight := p.Weight
	for _, support := range p.SupportedPrograms {
		program, _ := programs[support]
		weight += program.TotalWeight(programs)
	}
	return weight
}

func (p *Program) Balance(programs map[string]*Program) int {
	if p.IsBaseProgram() {
		return 0
	}

	programsByWeight := make(map[int][]*Program, len(p.SupportedPrograms))
	for _, support := range p.SupportedPrograms {
		program, _ := programs[support]
		weight := program.TotalWeight(programs)
		programsWithWeight, ok := programsByWeight[weight]
		if !ok {
			programsWithWeight = make([]*Program, 0, 1)
		}
		programsByWeight[weight] = append(programsWithWeight, program)
	}
	if len(programsByWeight) < 2 {
		return 0
	}

	targetWeight := 0
	var programToBalance *Program
	for weight, plist := range programsByWeight {
		if len(plist) != 1 {
			targetWeight = weight
			continue
		}

		programToBalance = plist[0]
	}
	if programToBalance.IsBalanced(programs) {
		return targetWeight - programToBalance.TotalWeight(programs) + programToBalance.Weight
	} else {
		return programToBalance.Balance(programs)
	}
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	isNewline := func(r rune) bool { return r == '\n' }
	lines := strings.FieldsFunc(string(input), isNewline)
	programs := make(map[string]*Program, len(lines))
	for _, line := range lines {
		program, err := NewProgramFromLine(line)
		if err != nil {
			panic(err)
		}
		programs[program.Name] = program
	}

	supportMap := make(map[string][]string, len(programs))
	for name, program := range programs {
		if _, ok := supportMap[name]; !ok {
			supportMap[name] = []string{}
		}
		for _, supportedProgram := range program.SupportedPrograms {
			dependencies, ok := supportMap[supportedProgram]
			if !ok {
				dependencies = make([]string, 0, 1)
			}
			supportMap[supportedProgram] = append(dependencies, name)
		}
	}

	var topProgram *Program
	for name, dependencies := range supportMap {
		if len(dependencies) == 0 {
			topProgram = programs[name]
			break
		}
	}
	fmt.Println(topProgram.Name)
	fmt.Println(topProgram.Balance(programs))
}
