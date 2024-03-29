package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func IsValidPartOne(p map[string]interface{}) bool {
	items := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	var count int
	for _, i := range items {
		if _, ok := p[i]; ok {
			count++
		}
	}
	if count == 7 {
		return true
	}
	return false
}

func IsValidPartTwo(p map[string]interface{}) bool {
	items := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	var count int
	for _, i := range items {
		//fmt.Println(reflect.TypeOf(p[i]))
		if _, ok := p[i]; ok {
			count++
		}
	}
	if count == 7 {
		// byr (Birth Year) - four digits; at least 1920 and at most 2002
		byr, _ := strconv.ParseInt(p["byr"].(string), 0, 16)
		if !(byr >= 1920 && byr <= 2002) {
			//fmt.Println(p["byr"], "byr nepasuje")
			return false
		}
		//// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		iyr, _ := strconv.ParseInt(p["iyr"].(string), 0, 16)
		if !(iyr >= 2010 && iyr <= 2020) {
			//fmt.Println(p["iyr"], "iyr nepasuje")
			return false
		}
		//// eyr (Expiration Year) - four digits; at least 2020 and at most 2030
		eyr, _ := strconv.ParseInt(p["eyr"].(string), 0, 16)
		if !(eyr >= 2020 && eyr <= 2030) {
			//fmt.Println(p["eyr"], "eyr nepasuje")
			return false
		}
		// hgt (Height) - a number followed by either cm or in:
		//     - If cm, the number must be at least 150 and at most 193.
		//     - If in, the number must be at least 59 and at most 76.
		hgt := regexp.MustCompile(`^(?P<value>\d+)(?P<units>[a-z]{2})$`)
		hgt_res := hgt.FindStringSubmatch(p["hgt"].(string))
		fmt.Println(hgt_res)
		if len(hgt_res) != 3 {
			//fmt.Println("missing units [cm|in]")
			return false
		}
		hight, _ := strconv.ParseInt(hgt_res[1], 0, 16)

		if hgt_res[2] == "cm" {
			if !(hight >= 150 && hight <= 193) {
				//fmt.Println(p["hgt"], "hgt nepasuje")
				return false
			}
		}
		if hgt_res[2] == "in" {
			if !(hight >= 59 && hight <= 76) {
				//fmt.Println(p["hgt"], "hgt nepasuje")
				return false
			}

		}
		// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f
		hcl := regexp.MustCompile(`^(?P<hash>#)(?P<value>[a-f0-9]{6})$`)
		hcl_res := hcl.FindStringSubmatch(p["hcl"].(string))
		//haircolor, _ := strconv.ParseInt(hcl_res[1], 0, 16)
		if len(hcl_res) == 0 {
			//fmt.Println("nie je tam zhoda")
			return false
		}
		// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth
		var ecl_count int = 0
		for _, v := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
			if p["ecl"] == v {
				ecl_count++
			}
		}
		if ecl_count == 0 {
			return false
		}
		// pid (Passport ID) - a nine-digit number, including leading zeroes
		pid := regexp.MustCompile(`^(?P<number>[0-9]{9})$`)
		//pid := regexp.MustCompile(`^(?P<hash>0+)(?P<rest>[0-9]{8})$`)
		pid_res := pid.FindStringSubmatch(p["pid"].(string))
		if len(pid_res) == 0 {
			//fmt.Println("prazdne")
			return false

		}

		// iconst not of above checks will maintch then return true :)
		return true
	}
	return false
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

	var entries []string
	var tmp string

	for _, i := range rawData {

		if i != "" {
			tmp = tmp + " " + i
		} else {
			entries = append(entries, tmp)
			tmp = ""
			continue
		}
	}

	var passportData []string
	for _, en := range entries {
		var t string
		t = strings.TrimLeft(strings.ReplaceAll(en, " ", ","), ",")
		passportData = append(passportData, t)
	}

	var couple string
	var pairs [][]string
	for _, i := range passportData {
		var pair []string
		for nr, x := range i {
			//fmt.Println(string(x))
			if string(x) != string(",") {
				couple = couple + string(x)
				// if this is the last charaster in a string you
				// are spliting at "," - the loop needs to be stopped and
				// last <key>:<value> should be appended to a slice :)
				if len(i)-1 == nr {
					pair = append(pair, couple)
					//fmt.Println(couple)
					couple = ""
					continue
				}
			} else {
				pair = append(pair, couple)
				//fmt.Println(couple)
				couple = ""
				continue
			}
		}

		pairs = append(pairs, pair)
	}

	var partOne int
	var partTwo int

	items := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	for _, pair := range pairs {
		// temporary empty Passport structt
		m := make(map[string]interface{})
		// getting through passports "pair"
		//fmt.Println(pair)
		for _, p := range pair {
			for _, item := range items {
				if strings.Contains(p, item) {
					//fmt.Println(p, item)
					m[item] = strings.Split(p, ":")[1]
				}

			}
		}

		//fmt.Println(m, IsValidPartOne(m))
		if IsValidPartOne(m) {
			partOne++
		}
		if IsValidPartTwo(m) {
			fmt.Println(m["byr"], "  ", m["iyr"], "  ", m["eyr"], "   ", m["hgt"], "  ", m["hcl"], "\t", m["ecl"], "\t", m["pid"], " ", m["cid"])
			partTwo++
		}
	}
	fmt.Println("Part1: Valid Passports: ", partOne)
	fmt.Println("Part2: Valid Passports: ", partTwo)
}
