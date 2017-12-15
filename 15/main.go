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

func step(val, factor int) int {
	return (val * factor) % DIVISOR
}

func isLower16Match(a, b int) bool {
	aBits := fmt.Sprintf("%.16b", a)
	bBits := fmt.Sprintf("%.16b", b)
	aLowerSixteen := aBits[len(aBits)-16:]
	bLowerSixteen := bBits[len(bBits)-16:]
	return aLowerSixteen == bLowerSixteen
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
		genA = step(genA, A_FACTOR)
		genB = step(genB, B_FACTOR)
		if isLower16Match(genA, genB) {
			matches++
		}
	}
	fmt.Println(matches)
}
