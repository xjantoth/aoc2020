package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("a1.txt")
	if err != nil {
		fmt.Println(err)
	}

	data := strings.Split(string(content), "\n")

	var r []int
	for _, str := range data {
		if str != "" {
			t, _ := strconv.Atoi(str)
			r = append(r, t)
		}
	}

	for _, v := range r {
		for _, vv := range r {
			if v+vv == 2020 {
				fmt.Println(v, " + ", vv, " = ", v+vv)
				fmt.Println(v, " * ", vv, " = ", v*vv)
				break
			}
		}
	}

	fmt.Println("---------------------")
	for _, v := range r {
		for _, vv := range r {
			for _, vvv := range r {
				if v+vv+vvv == 2020 {
					fmt.Println(v, " + ", vv, " + ", vvv, " = ", v+vv+vvv)
					fmt.Println(v, " * ", vv, " * ", vvv, " = ", v*vv*vvv)
					os.Exit(0)
				}

			}

		}
	}
}
