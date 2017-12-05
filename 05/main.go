package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Jumper struct {
	Iterations int
	position int
	offsets []int
	jumpMod func(int) int
}

func NewJumper(offsets []int, jumpMod func(int) int) *Jumper {
	return &Jumper{
		Iterations: 0,
		position: 0,
		offsets: offsets,
		jumpMod: jumpMod,
	}
}

func (j *Jumper) Next() {
	if j.Done() { return }
	j.Iterations++

	jump := j.offsets[j.position]
	j.offsets[j.position] = j.jumpMod(jump)
	j.position += jump
}

func (j *Jumper) Done() bool {
	return j.position >= len(j.offsets)
}

func readOffsets(input string) (offsets []int, err error) {
	fields := strings.FieldsFunc(input, func(r rune) bool { return r == '\n' })
	offsets = make([]int, 0, len(fields))

	for _, line := range fields {
		offset, err := strconv.Atoi(line)
		if err != nil {
			break
		}
		offsets = append(offsets, offset)
	}
	return
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	offsets, err := readOffsets(string(input))
	if err != nil {
		panic(err)
	}
	jumper := NewJumper(offsets, func(jump int) int { return jump + 1 })
	for !jumper.Done() {
		jumper.Next()
	}
	fmt.Println(jumper.Iterations)
}
