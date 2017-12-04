package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func isValidPassphrase(passphrase string) bool {
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

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	lines := strings.FieldsFunc(string(input), func(r rune) bool { return r == '\n' })
	valid := 0
	for _, line := range lines {
		if isValidPassphrase(line) {
			valid++
		}
	}
	fmt.Println(valid)
}
