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
}
