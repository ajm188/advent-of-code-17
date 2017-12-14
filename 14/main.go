package main

import (
	"fmt"
)

import (
	"github.com/ajm188/advent-of-code-17/util"
)

const (
	INPUT = "ljoxqyyw"
)

func main() {
	for i := 0; i < 128; i++ {
		codes := util.AsASCIICodes(fmt.Sprintf("%s-%d", INPUT, i))
		hk := util.NewHashKnot()
		hk.RunRound(codes)
		hash := hk.Hash()
		// TODO: convert to binary, count 1s
		fmt.Println(hash)
	}
}
