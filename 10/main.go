package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	KNOT_SIZE  = 256
	NUM_ROUNDS = 64
	BLOCK_SIZE = 16
)

type HashKnot struct {
	index    int
	skipSize int
	ring     []int
}

func NewHashKnot() *HashKnot {
	numbers := make([]int, KNOT_SIZE)
	i := 0
	for i < len(numbers) {
		numbers[i] = i
		i++
	}
	return &HashKnot{
		index:    0,
		skipSize: 0,
		ring:     numbers,
	}
}

func (hk *HashKnot) Hash() []int {
	numBlocks := len(hk.ring) / BLOCK_SIZE
	blocks := make([]int, 0, numBlocks)
	i := 0
	for i < numBlocks {
		start := i * BLOCK_SIZE
		end := start + BLOCK_SIZE
		block := hk.ring[start:end]

		v := block[0]
		j := 1
		for j < BLOCK_SIZE {
			v ^= block[j]
			j++
		}
		blocks = append(blocks, v)
		i++
	}
	return blocks
}

func (hk *HashKnot) RunRound(lengths []int) {
	for _, length := range lengths {
		hk.PinchTwist(length)
	}
}

func (hk *HashKnot) PinchTwist(length int) {
	if length > len(hk.ring) {
		return
	}

	buf := make([]int, length)
	for j, _ := range buf {
		i := ((hk.index + length) - j - 1) % len(hk.ring)
		buf[j] = hk.ring[i]
	}

	for i, v := range buf {
		j := (hk.index + i) % len(hk.ring)
		hk.ring[j] = v
	}

	hk.index = (hk.index + hk.skipSize + length) % len(hk.ring)
	hk.skipSize++
}

func readLengths(str string, f func(rune) bool) ([]int, error) {
	fields := strings.FieldsFunc(str, f)
	lengths := make([]int, 0, len(fields))
	for _, f := range fields {
		length, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
		lengths = append(lengths, length)
	}
	return lengths, nil
}

func asASCIICodes(str string) []int {
	buf := make([]int, 0, len(str))
	for _, v := range str {
		buf = append(buf, int(v))
	}
	return buf
}

func isComma(r rune) bool {
	return r == ','
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	lengths, err := readLengths(string(input[0:len(input)-1]), isComma)
	if err != nil {
		panic(err)
	}
	hk := NewHashKnot()
	hk.RunRound(lengths)

	fmt.Println(hk.ring[0] * hk.ring[1])

	hk = NewHashKnot()
	lengths = append(
		asASCIICodes(string(input[0:len(input)-1])),
		[]int{17, 31, 73, 47, 23}...,
	)
	i := 0
	for i < NUM_ROUNDS {
		hk.RunRound(lengths)
		i++
	}
	blocks := hk.Hash()
	for _, block := range blocks {
		if block < 16 {
			fmt.Printf("0")
		}
		fmt.Printf("%x", block)
	}
	fmt.Printf("\n")
}
