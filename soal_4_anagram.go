package main

import (
	"fmt"
	"sort"
)

func main() {
	anagram := []string{
		"kita", "atik", "tika", "aku", "kia", "makan", "kua",
	}
	a := Anagram(anagram)
	fmt.Println(a)

}

func Anagram(anagram []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, v := range anagram {
		word := []byte(v)
		sort.SliceStable(word, func(i, j int) bool {
			return word[i] < word[j]
		})
		s := string(word)
		anagramMap[s] = append(anagramMap[s], v)
	}
	var group [][]string
	for i := range anagramMap {
		group = append(group, anagramMap[i])
	}

	return group
}
