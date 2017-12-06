package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type MemoryBank struct {
	banks      []int
	Iterations int
}

func NewMemoryBank(banks []int) *MemoryBank {
	return &MemoryBank{
		banks:      banks,
		Iterations: 0,
	}
}

func (m *MemoryBank) Next() *MemoryBank {
	i := findMax(m.banks)
	redistributionAmount := m.banks[i]
	newBanks := append([]int(nil), m.banks...)
	newBanks[i] = 0
	for redistributionAmount > 0 {
		i = (i + 1) % len(newBanks)
		newBanks[i]++
		redistributionAmount--
	}
	return &MemoryBank{
		banks:      newBanks,
		Iterations: m.Iterations + 1,
	}
}

func findMax(arr []int) (index int) {
	if len(arr) == 0 {
		return -1
	}
	index = 0
	max := arr[index]
	for i, v := range arr {
		if v > max {
			max = v
			index = i
		}
	}
	return
}

func readBanks(input string) (banks []int, err error) {
	fields := strings.Fields(input)
	banks = make([]int, 0, len(fields))
	for _, field := range fields {
		bank, err := strconv.Atoi(field)
		if err != nil {
			break
		}
		banks = append(banks, bank)
	}
	return
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	input = input[0 : len(input)-1]
	banks, err := readBanks(string(input))
	if err != nil {
		panic(err)
	}

	fmt.Println(banks)
}
