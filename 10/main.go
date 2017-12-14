package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

import (
	"github.com/ajm188/advent-of-code-17/util"
)

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
	hk := util.NewHashKnot()
	hk.RunRound(lengths)

	fmt.Println(hk.Ring[0] * hk.Ring[1])

	hk = util.NewHashKnot()
	lengths = append(
		util.AsASCIICodes(string(input[0:len(input)-1])),
		[]int{17, 31, 73, 47, 23}...,
	)
	i := 0
	for i < util.NUM_ROUNDS {
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
