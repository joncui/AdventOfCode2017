package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	rawData, err := ioutil.ReadFile("data")
	check(err)
	countPt1 := 0
	countPt2 := 0

	rows := strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n")
	for _, row := range rows {
		words := strings.Split(row, " ")
		if !HasDuplicate(words) {
			countPt1++
		}

		if !hasDuplicatePt2(words) {
			countPt2++
		}
	}

	fmt.Printf("Part One: %d passphrases are valid\n", countPt1)
	fmt.Printf("Part Two: %d passphrases are valid\n", countPt2)
}

func HasDuplicate(words []string) (hasDuplicate bool) {
	hasDuplicate = false

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if words[i] == words[j] {
				hasDuplicate = true
				return
			}
		}
	}

	return
}

func hasDuplicatePt2(words []string) (hasDuplicate bool) {
	var sortedWords []string
	for _, word := range words {
		sortedWords = append(sortedWords, SortString(word))
	}

	return HasDuplicate(sortedWords)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
