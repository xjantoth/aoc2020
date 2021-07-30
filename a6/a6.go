package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: " + os.Args[0] + " <fileName>")
		os.Exit(1)
	}
	//fmt.Println("Usage: " + os.Args[0] + " file name")
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawData := strings.Split(string(content), "\n")
	var data []string
	var tmp string

	for _, i := range rawData {

		if i != "" {
			tmp = tmp + " " + i
		} else {
			data = append(data, tmp)
			tmp = ""
			continue
		}
	}

	fmt.Println(len(data))
	fmt.Println(data)

	var numOfpeople []int
	var answersTotal int = 0
	var correctAnsPart2 int = 0

	for _, group := range data {
		var spaces int = 0

		// Part 1.
		var groupTmp string = strings.TrimLeft(string(group), " ")

		set := make(map[rune]bool)

		for _, char := range groupTmp {
			//fmt.Println(char)
			if char == 32 {
				spaces++
			}

			// Part 1
			if _, ok := set[char]; !ok && char != 32 {
				set[char] = true
			}

		}

		// Part 2.
		numOfpeople = append(numOfpeople, spaces+1)

		people := spaces + 1

		var keys []rune
		for k := range set {
			keys = append(keys, k)
		}

		//  if one person -> then each letter once
		part2map := make(map[rune]int)

		for _, char := range groupTmp {
			if _, ok := part2map[char]; ok {
				var tmpVal int = part2map[char]
				part2map[char] = tmpVal + 1
			} else {
				part2map[char] = 1
			}

		}
		// deleting space 32
		delete(part2map, 32)

		for k, v := range part2map {
			if v == people {

				fmt.Println(k, v)
				correctAnsPart2++
			}
		}
		fmt.Printf("People: %v p2: %v\n", people, part2map)

		answersTotal = answersTotal + len(set)

	}
	fmt.Println(numOfpeople)
	fmt.Println("Part 1: ", answersTotal)
	fmt.Println("Part 2: ", correctAnsPart2)
}
