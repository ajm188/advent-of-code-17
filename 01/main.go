package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	input = input[0:len(input)-1]

	sum := 0
	for i, c := range input {
		num, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}
		j := (i + 1) % len(input)
		if input[j] == c {
			sum += num
		}
	}
	fmt.Println(sum)
}
