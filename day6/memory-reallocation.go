package main

import (
	"fmt"
	"strconv"
)

func main() {
	pt1 := []int{10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6}
	pt2 := []int{1, 1, 0, 15, 14, 13, 12, 10, 10, 9, 8, 7, 6, 4, 3, 5}
	// test := []int{0, 2, 7, 0}
	pt1Cycles, pt1Data := findCycles(pt1)
	fmt.Printf("Part One: %d -> %v\n", pt1Cycles, pt1Data)

	pt2Cycles, pt2Data := findCycles(pt2)
	fmt.Printf("Part Two: %d -> %v\n", pt2Cycles, pt2Data)
}

func findCycles(data []int) (int, []int) {
	maxIndex := len(data)
	var index, value int
	currentState := intArrayToString(data)

	seen := make(map[string]bool)
	seen[currentState] = true

	for true {
		index, value = findLargest(data)
		data[index] = 0
		pointer := index + 1
		for i := 0; i < value; i++ {
			if pointer > maxIndex-1 {
				pointer %= maxIndex
			}
			data[pointer]++
			pointer++
		}

		currentState = intArrayToString(data)
		if seen[currentState] {
			break
		} else {
			seen[currentState] = true
		}
	}

	return len(seen), data
}

func intArrayToString(intArr []int) (str string) {
	for _, v := range intArr {
		str += strconv.Itoa(v)
	}

	return
}

func findLargest(intArr []int) (index, value int) {
	for i, v := range intArr {
		if v > value {
			value = v
			index = i
		}
	}

	return
}
