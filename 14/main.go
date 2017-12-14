package main

import (
	"fmt"
	"strconv"
	"strings"
)

import (
	"github.com/ajm188/advent-of-code-17/util"
)

const (
	INPUT = "ljoxqyyw"
)

func main() {
	ones := 0
	for i := 0; i < 128; i++ {
		codes := util.AsASCIICodes(fmt.Sprintf("%s-%d", INPUT, i))
		hk := util.NewHashKnot()
		codes = append(codes, []int{17, 31, 73, 47, 23}...)
		for j := 0; j < util.NUM_ROUNDS; j++ {
			hk.RunRound(codes)
		}
		hash := hk.Hash()
		row := ""
		hexRow := ""
		for _, i := range hash {
			hex := fmt.Sprintf("%.2x", i)
			hexRow += hex
			for _, d := range hex {
				var dec int
				switch d {
				case 'a', 'b', 'c', 'd', 'e', 'f':
					dec = int(d-'a') + 10
				default:
					dec, _ = strconv.Atoi(string(d))
				}
				bin := fmt.Sprintf("%.4b", dec)
				row += bin
			}
		}
		ones += strings.Count(row, "1")
	}
	fmt.Println(ones)
}
