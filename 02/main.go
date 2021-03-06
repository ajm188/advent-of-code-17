package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func toMatrix(input []byte) ([][]int, error) {
	fields := strings.FieldsFunc(
		string(input),
		func(c rune) bool {
			return c == '\n'
		},
	)
	matrix := make([][]int, 0, len(fields))
	for _, field := range fields {
		row := make([]int, 0, len(field))
		for _, c := range strings.Fields(field) {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				return matrix, err
			}
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}
	return matrix, nil
}

func minMax(row []int) (min, max int) {
	min = row[0]
	max = row[0]
	for _, n := range row[1:len(row)] {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return
}

func checksum1(matrix [][]int) int {
	sum := 0
	for _, row := range matrix {
		min, max := minMax(row)
		sum += (max - min)
	}
	return sum
}

func findDivisorPair(row []int) (int, int) {
	for i, x := range row {
		for j, y := range row {
			if i == j || x < y {
				continue
			}
			if (x % y) == 0 {
				return x, y
			}
		}
	}
	return 1, 1
}

func checksum2(matrix [][]int) int {
	sum := 0
	for _, row := range matrix {
		divisor, dividend := findDivisorPair(row)
		sum += (divisor / dividend)
	}
	return sum
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	mat, err := toMatrix(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(checksum1(mat), checksum2(mat))
}
