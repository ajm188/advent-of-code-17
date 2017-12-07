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
		for _, field := range fields[2:len(fields)] {
			if strings.HasSuffix(field, ",") {
				field = field[0:len(field)-1]
			}
			program.SupportedPrograms = append(program.SupportedPrograms, field)
		}
	}

	return program, nil
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	isNewline := func(r rune) bool { return r == '\n' }
	lines := strings.FieldsFunc(string(input), isNewline)
	programs := make([]*Program, 0, len(lines))
	for _, line := range lines {
		program, err := NewProgramFromLine(line)
		if err != nil {
			panic(err)
		}
		programs = append(programs, program)
	}

	supportMap := make(map[string][]string, len(programs))
	for _, program := range programs {
		if _, ok := supportMap[program.Name]; !ok {
			supportMap[program.Name] = []string{}
		}
		for _, supportedProgram := range program.SupportedPrograms {
			dependencies, ok := supportMap[supportedProgram]
			if !ok {
				dependencies = make([]string, 0, 1)
			}
			supportMap[supportedProgram] = append(dependencies, program.Name)
		}
	}

	for name, dependencies := range supportMap {
		if len(dependencies) == 0 {
			fmt.Println(name)
			break
		}
	}
}
