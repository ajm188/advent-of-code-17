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

	sum1 := 0
	sum2 := 0
	n := len(input)
	for i, c := range input {
		num, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}
		j := (i + 1) % n
		k := (i + (n / 2)) % n
		if input[j] == c {
			sum1 += num
		}
		if input[k] == c {
			sum2 += num
		}
	}
	fmt.Println(sum1, sum2)
}
