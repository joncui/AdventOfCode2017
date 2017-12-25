package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
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

	rows := toIntArray(strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n"))

	fmt.Printf("Part One: It took %d steps.\n", partOne(rows))
	// fmt.Printf("Part Two: It took %d steps.\n", partTwo(rows))
}

func partOne(rows []int) (steps int) {
	curr := 0
	for i := 0; i < len(rows); steps++ {
		curr = rows[i]
		rows[i]++
		i += curr
	}

	return
}

func partTwo(rows []int) (steps int) {
	curr := 0
	for i := 0; i < len(rows); steps++ {
		curr = rows[i]
		if curr >= 3 {
			rows[i]--
		} else {
			rows[i]++
		}
		i += curr
	}

	return
}

func convertStringToInt(str string) int {
	value, err := strconv.Atoi(str)
	check(err)

	return value
}

func toIntArray(strArray []string) (intArray []int) {
	for _, e := range strArray {
		intArray = append(intArray, convertStringToInt(e))
	}
	return
}
