package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Move interface {
	Execute([]string) []string
}

type SpinMove struct {
	amount int
}

func (s *SpinMove) Execute(programs []string) []string {
	result := append([]string{}, programs...)
	for i := 0; i < s.amount; i++ {
		result = append([]string{result[len(result)-1]}, result[:len(result)-1]...)
	}
	return result
}

type ExchangeMove struct {
	positionA, positionB int
}

func (e *ExchangeMove) Execute(programs []string) []string {
	result := make([]string, len(programs))
	for i, v := range programs {
		if i == e.positionA {
			result[i] = programs[e.positionB]
		} else if i == e.positionB {
			result[i] = programs[e.positionA]
		} else {
			result[i] = v
		}
	}
	return result
}

type PartnerMove struct {
	nameA, nameB string
}

func (p *PartnerMove) Execute(programs []string) []string {
	result := make([]string, len(programs))
	positionA, positionB := -1, -1
	for i, v := range programs {
		if v == p.nameA {
			positionA = i
		} else if v == p.nameB {
			positionB = i
		} else {
			result[i] = v
		}
	}
	result[positionA] = p.nameB
	result[positionB] = p.nameA
	return result
}

func readMoves(lines []string) ([]Move, error) {
	slashSplitter := func(r rune) bool {
		return r == '/'
	}

	moves := make([]Move, len(lines))
	for i, line := range lines {
		var move Move
		switch line[0] {
		case 's':
			amount, err := strconv.Atoi(line[1:])
			if err != nil {
				return moves, err
			}
			move = &SpinMove{amount}
		case 'x':
			fields := strings.FieldsFunc(line[1:], slashSplitter)
			posA, err := strconv.Atoi(fields[0])
			if err != nil {
				return moves, err
			}
			posB, err := strconv.Atoi(fields[1])
			if err != nil {
				return moves, err
			}
			move = &ExchangeMove{posA, posB}
		case 'p':
			fields := strings.FieldsFunc(line[1:], slashSplitter)
			move = &PartnerMove{fields[0], fields[1]}
		default:
			return moves, fmt.Errorf("Unknown move type %s", line)
		}
		moves[i] = move
	}

	return moves, nil
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	input = input[:len(input)-1]

	moves, err := readMoves(strings.FieldsFunc(string(input), func(r rune) bool { return r == ',' }))
	if err != nil {
		panic(err)
	}

	programs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}
	for _, move := range moves {
		programs = move.Execute(programs)
	}
	fmt.Println(programs)
}
