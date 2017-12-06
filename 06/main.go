package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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
