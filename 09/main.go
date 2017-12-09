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

func cleanGarbage(text string) (string, string) {
	cleanBuf := make([]byte, 0, len(text))
	garbageBuf := make([]byte, 0, len(text))
	i := 0
	garbage := false
	for i < len(text) {
		r := text[i]
		if r == '!' {
			i++
		} else if garbage {
			if r == '>' {
				garbage = false
			} else {
				garbageBuf = append(garbageBuf, r)
			}
		} else {
			if r == '<' {
				garbage = true
			} else {
				cleanBuf = append(cleanBuf, r)
			}
		}
		i++
	}
	return string(cleanBuf), string(garbageBuf)
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	cleaned, garbage := cleanGarbage(string(input))
	fmt.Println(groupScore(cleaned))
	fmt.Println(len(garbage))
}
