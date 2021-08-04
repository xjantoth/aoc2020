package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: " + os.Args[0] + " <fileName>")
		os.Exit(1)
	}
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawData := strings.Split(string(content), "\n")

	var data []string
	for _, val := range rawData {
		if val != "" {
			data = append(data, val)
		}
	}

	m := make(map[string]bool)
	var w []string
	w = append(w, "shiny gold")

	wLenPrevious := len(w)
	wLenNext := 0

	var iterations int = 0 // artificial valuee

	fmt.Println("Length of map: ", len(m))

	for wLenPrevious != wLenNext {
		// previous = 1

		iterations++
		wLenPrevious = wLenNext

		for nl, row := range data {
			r := regexp.MustCompile(`\sbags?`)
			row := r.ReplaceAllString(row, "") // remove words "bag/bags"
			// ...
			// [key]					    [value]
			// light red     contain        1 bright white, 2 muted yellow.
			// dark orange   contain        3 bright white, 4 muted yellow.
			// ...
			key, value := strings.TrimSpace(strings.Split(row, "contain")[0]), strings.Split(row, "contain")[1] // whiteat can be stored in bag :)
			fmt.Printf("%v - %v:   %v\n", nl, key, value)

			for _, k := range w {
				if strings.Contains(value, k) {
					if _, ok := m[key]; !ok {
						m[key] = true
						w = append(w, key)
					}
				}
			}

		}
		// next  = 2
		wLenNext = len(w)
	}
	fmt.Println("Length of map: ", len(m))
	fmt.Println(m)

}
