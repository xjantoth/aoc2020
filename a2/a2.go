package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Password struct {
	min    int    //
	max    int    //
	letter string //
	pass   string //
}

func (p Password) Check() bool {
	occurance := strings.Count(p.pass, p.letter)
	if p.min <= occurance && p.max >= occurance {
		return true
	} else {
		return false
	}

}

func (p Password) Part2() bool {
	left := strings.Count(string(p.pass[p.min-1]), p.letter)
	right := strings.Count(string(p.pass[p.max-1]), p.letter)

	if left != right {
		return true
	} else {
		return false

	}

}

func main() {
	content, err := ioutil.ReadFile("a2.txt")
	if err != nil {
		fmt.Println(err)
	}

	data := strings.Split(string(content), "\n")

	var r []string
	for _, str := range data {
		if str != "" {
			//t, _ := strconv.Atoi(str)
			r = append(r, str)
		}
	}
	//fmt.Println(r)

	var part1 int
	var part2 int

	for _, v := range r {
		myRange := strings.Split(v, " ")[0]
		min, _ := strconv.Atoi(strings.Split(myRange, "-")[0])
		max, _ := strconv.Atoi(strings.Split(myRange, "-")[1])
		letter := strings.Replace(strings.Split(v, " ")[1], ":", "", -1)
		pass := strings.Split(v, " ")[2]
		//fmt.Println("min: ", min, "max: ", max, "letter: ", letter, "password: ", pass)
		tmp := Password{
			min:    min,
			max:    max,
			letter: letter,
			pass:   pass,
		}
		if tmp.Check() {
			part1++
		}
		if tmp.Part2() {
			part2++
		}
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)

	//1 1 false
	//1 0 true
	//0 1 true
	//1 1 false

}
