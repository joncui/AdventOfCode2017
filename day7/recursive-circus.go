package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	name      string
	weight    int
	children  []string
	sumWeight int
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	rawData, err := ioutil.ReadFile("data")
	check(err)

	sequences := strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n")

	nameWeightMap := parseSequences(sequences)
	// nameNodeMap := make(map[string]Node)

	fmt.Printf("%v\n", nameWeightMap)
}

func parseSequences(sequences []string) (nameWeightMap map[string]Node) {
	nameWeightMap = make(map[string]Node)

	re := regexp.MustCompile(`([a-z]+) \(([0-9]+)\)[\-\> ]*(.*)`)

	for _, sequence := range sequences {
		program := re.FindAllStringSubmatch(sequence, -1)[0][1:]

		name := program[0]
		weight, err := strconv.Atoi(program[1])
		check(err)

		if len(program) == 3 && len(program[2]) > 1 {
			children := strings.Split(program[2], ", ")
			nameWeightMap[name] = Node{name, weight, children, weight}
		} else {
			nameWeightMap[name] = Node{name, weight, nil, weight}
		}
	}

	return
}
