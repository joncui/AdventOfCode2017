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

	sequences := strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n")

	for _, sequence := range sequences {
		sum := findSum(sequence)
		fmt.Printf("The part 1 solution to the captcha is %d.\n", sum)

		sum = findSumPt2(sequence)
		fmt.Printf("The part 2 solution to the captcha is %d.\n", sum)
	}
}

func findSum(sequence string) (sum int) {
	initDigit, err := strconv.Atoi(string(sequence[0]))
	check(err)

	prevDigit := initDigit
	for i := 1; i < len(sequence); i++ {
		currDigit, err := strconv.Atoi(string(sequence[i]))
		check(err)

		if prevDigit == currDigit {
			sum += currDigit
		}

		prevDigit = currDigit
	}

	if prevDigit == initDigit {
		sum += initDigit
	}

	return
}

func findSumPt2(sequence string) (sum int) {
	length := len(sequence)
	half := length / 2

	for i := 0; i < length; i++ {
		first, err := strconv.Atoi(string(sequence[i]))
		check(err)

		secondI := (i + half) % length
		// secondI := int(math.Mod(float64(i+half), float64(length)))
		second, err := strconv.Atoi(string(sequence[secondI]))
		check(err)

		if first == second {
			sum += first
		}

	}

	return
}
