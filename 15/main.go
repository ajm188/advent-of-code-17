package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	A_FACTOR = 16807
	B_FACTOR = 48271
	DIVISOR  = 2147483647
	PAIRS    = 40000000
)

func readInitialValue(line string) (int, error) {
	fields := strings.Fields(line)
	return strconv.Atoi(fields[len(fields)-1])
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	lines := strings.FieldsFunc(string(input), func(r rune) bool { return r == '\n' })
	genA, err := readInitialValue(lines[0])
	if err != nil {
		panic(err)
	}
	genB, err := readInitialValue(lines[1])
	if err != nil {
		panic(err)
	}

	matches := 0
	for i := 0; i < PAIRS; i++ {
		genA = (genA * A_FACTOR) % DIVISOR
		genB = (genB * B_FACTOR) % DIVISOR

		aBits := fmt.Sprintf("%.16b", genA)
		bBits := fmt.Sprintf("%.16b", genB)
		aLowerSixteen := aBits[len(aBits)-16:]
		bLowerSixteen := bBits[len(bBits)-16:]

		if aLowerSixteen == bLowerSixteen {
			matches++
		}
	}
	fmt.Println(matches)
}
