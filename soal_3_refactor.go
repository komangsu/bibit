package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "bibit(stockbit)"

	find := findStringInBracket(a)
	fmt.Println("find:", find)

}

// sebelum di refactor
func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1 : indexClosingBracketFound-1])
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}
}

// setelah di refactor
func findStringInBracket(str string) string {
	i := strings.Index(str, "(")
	if i >= 0 {
		j := strings.Index(str, ")")
		if j >= 0 {
			return str[i+1 : j]
		}
	}
	return ""
}
