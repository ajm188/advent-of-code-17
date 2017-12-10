package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	KNOT_SIZE = 256
)

type HashKnot struct {
	index int
	skipSize int
	ring []int
}

func NewHashKnot() *HashKnot {
	numbers := make([]int, KNOT_SIZE)
	i := 0
	for i < len(numbers) {
		numbers[i] = i
		i++
	}
	return &HashKnot{
		index: 0,
		skipSize: 0,
		ring: numbers,
	}
}

func (hk *HashKnot) PinchTwist(length int) {
	if length > len(hk.ring) { return }

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

func readLengths(fields []string) ([]int, error) {
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

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	lengthFields := strings.FieldsFunc(
		string(input),
		func(r rune) bool { return r == ',' || r == '\n' },
	)
	lengths, err := readLengths(lengthFields)
	if err != nil {
		panic(err)
	}
	hk := NewHashKnot()
	for _, length := range lengths {
		hk.PinchTwist(length)
	}

	fmt.Println(hk.ring[0] * hk.ring[1])
}
