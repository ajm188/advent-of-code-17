package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func groupScore(text string) int {
	score := 0
	depth := 0
	for _, r := range text {
		if r == '{' {
			depth++
			score += depth
		} else if r == '}' && depth > 0 {
			depth--
		}
	}
	return score
}

func cleanGarbage(text string) string {
	buf := make([]byte, 0, len(text))
	i := 0
	garbage := false
	for i < len(text) {
		r := text[i]
		if r == '!' {
			i++
		} else if garbage {
			if r == '>' {
				garbage = false
			}
		} else {
			if r == '<' {
				garbage = true
			} else {
				buf = append(buf, r)
			}
		}
		i++
	}
	return string(buf)
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	cleaned := cleanGarbage(string(input))
	fmt.Println(groupScore(cleaned))
}
