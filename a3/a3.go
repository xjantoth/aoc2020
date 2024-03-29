package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func HitTrees(data []string, right int, down int) int {
	var treesCount int = 0
	for nr, v := range data {
		//fmt.Println(v)
		var isTree bool = false
		if down > 1 && (nr)%down == 1 {
			continue
		}

		idx := right * (nr / down)
		if len(v) <= idx {
			v = strings.Repeat(v, nr)
		}

		if string(v[idx]) == "#" {
			isTree = true
			//fmt.Println(nr)
			treesCount++
		}
		fmt.Printf("row: %v --> [idx]: %v (%v, %v) | \n", nr, idx, string(v[idx]), isTree)
	}
	fmt.Println("Trees encountered: ", treesCount)
	return treesCount
}

func main() {
	content, err := ioutil.ReadFile("a3.txt")
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

	var part2 []int
	for _, i := range []int{1, 3, 5, 7} {
		part2 = append(part2, HitTrees(r, i, 1))

	}

	part2 = append(part2, HitTrees(r, 1, 2))
	fmt.Println(part2)

	var res int = 1
	for _, i := range part2 {
		res = res * i
	}
	fmt.Println(res)

}
