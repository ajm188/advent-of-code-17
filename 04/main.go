package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type LetterCount map[rune]int

func (lc LetterCount) isEqual(other LetterCount) bool {
	if lc == nil && other == nil {
		return true
	}
	if (lc != nil) != (other != nil) {
		return false
	}
	if len(lc) != len(other) {
		return false
	}

	for k, v := range lc {
		if v2, ok := other[k]; !ok || v != v2 {
			return false
		}
	}
	return true
}

func count(word string) LetterCount {
	counter := make(map[rune]int, len(word))
	for _, c := range word {
		counter[c]++
	}
	return counter
}

func isValidPassphrase1(passphrase string) bool {
	words := strings.Fields(passphrase)
	cache := make(map[string]bool, len(words))
	for _, word := range words {
		if _, exists := cache[word]; exists {
			return false
		}
		cache[word] = true
	}
	return true
}

func isValidPassphrase2(passphrase string) bool {
	words := strings.Fields(passphrase)
	for i, word := range words {
		counter := count(word)
		for j, otherWord := range words {
			if i == j {
				continue
			}
			if word == otherWord || counter.isEqual(count(otherWord)) {
				return false
			}
		}
	}
	return true
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	lines := strings.FieldsFunc(string(input), func(r rune) bool { return r == '\n' })
	valid1, valid2 := 0, 0
	for _, line := range lines {
		if isValidPassphrase1(line) {
			valid1++
		}
		if isValidPassphrase2(line) {
			valid2++
		}
	}
	fmt.Println(valid1)
	fmt.Println(valid2)
}
