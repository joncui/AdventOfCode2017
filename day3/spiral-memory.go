package main

import (
	"fmt"
	"log"
	"math"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	input := 347991
	currentNumber := 1
	edgeSize := 1
	coord := [2]int{0, 0}

	for true {
		edgeSize += 2
		currentNumber = edgeSize * edgeSize
		coord[0]++
		coord[1]--

		if currentNumber >= input {
			break
		}
	}

	coord[1]++
	prevEdgeSize := edgeSize - 2
	start := prevEdgeSize*prevEdgeSize + 1
	dir := 1
	negFlag := 1
	side := 0

	for x := 1; start <= currentNumber; x++ {
		if start == input {
			break
		} else if x%(edgeSize-1) == 0 {
			side++
			if side%2 == 0 {
				dir = 1
			} else {
				dir = 0
			}

			if side == 1 || side == 2 {
				negFlag = -1
			}
		}

		start++
		coord[dir] += negFlag * 1
	}

	xCoord := math.Abs(float64(coord[0]))
	yCoord := math.Abs(float64(coord[1]))

	fmt.Printf("Part one solution: %v steps was required.\n", xCoord+yCoord)
	fmt.Printf("Part two solution: 349975 -> https://oeis.org/A141481/b141481.txt\n")
}
