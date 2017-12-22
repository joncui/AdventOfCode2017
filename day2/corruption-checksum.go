package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

	rows := strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n")
	checkSum1 := 0
	checkSum2 := 0

	for _, row := range rows {
		columns := strArrToIntArr(strings.Split(row, " "))

		checkSum1 += findMinMaxDiff(columns)
		checkSum2 += findDivisibleMinMax(columns)
	}

	fmt.Printf("The solution to the checksum is %d.\n", checkSum1)
	fmt.Printf("The solution to the checksum is %d.\n", checkSum2)
}

func findMinMaxDiff(values []int) int {
	min := math.MaxUint8
	max := 0

	for _, v := range values {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return max - min
}

func findDivisibleMinMax(values []int) int {
	for i, v1 := range values {
		for _, v2 := range values[i+1:] {
			if min, max, success := getMinMax(v1, v2); success {
				return max / min
			}
		}
	}

	return 0
}

func convertStringToInt(str string) int {
	value, err := strconv.Atoi(str)
	check(err)

	return value
}

func strArrToIntArr(strArr []string) (intArr []int) {
	for _, v := range strArr {
		intArr = append(intArr, convertStringToInt(v))
	}

	return
}

func getMinMax(x, y int) (int, int, bool) {
	min := y
	max := x

	if x < y {
		max = y
		min = x
	}

	if max%min == 0 {
		return min, max, true
	}

	return 0, 0, false
}
