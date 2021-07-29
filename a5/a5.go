package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
)

func Unique(intSlice []float64) []float64 {
	keys := make(map[float64]bool)
	list := []float64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func findMinAndMax(a []float64) (min float64, max float64) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

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
	for _, entry := range rawData {
		if entry != "" {
			data = append(data, entry)
			//fmt.Println(nr, entry)
		}

	}
	//fmt.Println(data)
	//fmt.Println("------ ------")

	var d float64 = 0
	var u float64 = 127

	var l float64 = 0
	var r float64 = 7

	var tmpColumn float64 = 0
	var seatIds []float64

	for _, ent := range data {
		//fmt.Println(nr, ent)

		for id, i := range ent {
			//if id >= 7 {
			//	fmt.Println("Getting through indes: ", id)
			//}
			if id > 6 {
				// reset values up and down for rows
				d = 0
				u = 127
				//continue
			}
			if string(i) == "F" {
				// 0 - 63
				//u = (u-d)/2 + d - 1
				u = math.Round((u-d)/2+d) - 1
				//fmt.Printf("front: \t[%v] <%v, %v>\n", string(i), d, u)
			}
			if string(i) == "B" {
				//d = (u-d)/2 + d + 1
				d = math.Round((u-d)/2 + d)
				//fmt.Printf("back: \t[%v] <%v, %v>\n", string(i), d, u)
			}
			if id == 6 {
				//fmt.Printf("row:\t\t%v, %v\n", (d+u)/2, ((d+u)/2)*8)
				tmpColumn = ((d + u) / 2) * 8
			}

			if id >= 7 {
				if string(i) == "R" {
					// 0 - 63
					//u = (u-d)/2 + d - 1
					l = math.Round((r-l)/2 + l)
					//fmt.Printf("right: \t[%v] <%v, %v>\n", string(i), l, r)
				}
				if string(i) == "L" {
					//d = (u-d)/2 + d + 1
					r = math.Round((r-l)/2+l) - 1
					//fmt.Printf("left: \t[%v] <%v, %v>\n", string(i), l, r)
				}

			}
			// reset columns for next iteration
			if id == 9 {

				//fmt.Printf("column:\t\t%v: %v\n", ent, math.Round((r+l)/2))

				seatIds = append(seatIds, tmpColumn+math.Round((r+l)/2))
				tmpColumn = 0
				l = 0
				r = 7
			}
		}
	}
	//fmt.Println("seatIds: ", seatIds)
	min, max := findMinAndMax(seatIds)

	//fmt.Println("sorting:")
	sort.Float64s(seatIds)
	//fmt.Println(min, max)
	//fmt.Println(seatIds)

	list := Unique(seatIds)
	sort.Float64s(list)
	//fmt.Println(list)

	//for i := min; i < max; i++ {
	//	fmt.Println(i)
	//}

	fmt.Println("Part1: ", max)
	var counter float64 = min

	for _, val := range list {
		if val != counter {
			fmt.Println("Part 2: this is missing number", counter)
			counter = counter + 1
		}
		counter = counter + 1
	}

}

// F B F B B F F R L R
