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
	V2_PAIRS = 5000000
)

type V2Generator struct {
	values        []int
	factor        int
	discriminator func(int) bool
}

func NewV2Generator(initialValue, factor int, discriminator func(int) bool) *V2Generator {
	return &V2Generator{
		values:        []int{initialValue},
		factor:        factor,
		discriminator: discriminator,
	}
}

func (g *V2Generator) Last() int {
	return g.values[len(g.values)-1]
}

func (g *V2Generator) Next() {
	previousValue := g.Last()
	next := step(previousValue, g.factor)
	for !g.discriminator(next) {
		next = step(next, g.factor)
	}
	g.values = append(g.values, next)
}

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

	v2GenA := NewV2Generator(genA, A_FACTOR, func(i int) bool { return i%4 == 0 })
	v2GenB := NewV2Generator(genB, B_FACTOR, func(i int) bool { return i%8 == 0 })

	matches := 0
	for i := 0; i < PAIRS; i++ {
		genA = step(genA, A_FACTOR)
		genB = step(genB, B_FACTOR)
		if isLower16Match(genA, genB) {
			matches++
		}
	}
	fmt.Println(matches)

	v2Matches := 0
	for i := 0; i < V2_PAIRS; i++ {
		v2GenA.Next()
		v2GenB.Next()

		if isLower16Match(v2GenA.Last(), v2GenB.Last()) {
			v2Matches++
		}
	}
	fmt.Println(v2Matches)
}
