package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"strconv"
)

func validate(currentRecord []string) bool {
	m := make(map[string]string)
	for _, entry := range currentRecord {
		z := strings.Split(entry, ":")
		m[z[0]] = z[1]
	}
	if m["byr"] >= "1920" && m["byr"] <= "2002" && m["iyr"] >= "2010" && m["iyr"] <= "2020" &&
		m["eyr"] >= "2020" && m["eyr"] <= "2030" && len(m["hgt"]) > 0 && len(m["hcl"]) > 0 &&
		len(m["ecl"]) > 0 && len(m["pid"]) > 0 {
		re := regexp.MustCompile("(\\d+)(\\w+)")
		res := re.FindAllStringSubmatch(m["hgt"], 1)[0]
		value, _ := strconv.Atoi(res[1])
		units := res[2]
		fmt.Printf("Value %d\n", value) 
		fmt.Println("UNITS: " + units)
		if (value >= 150 && value <= 193 && units == "cm") || (value >= 59 && value <= 76 && units == "in") {
			hairColourValid, _ := regexp.MatchString("#[0-9a-z]{6}", m["hcl"])
			eyeColourValid, _ := regexp.MatchString("(amb|blu|brn|gry|grn|hzl|oth)", m["ecl"])
			pidValid, _ := regexp.MatchString("^[0-9]{9}$", m["pid"])
			fmt.Printf("%t %t %t\n", eyeColourValid, hairColourValid, pidValid)
			if eyeColourValid && hairColourValid && pidValid {
				fmt.Printf("Valid: ")
				fmt.Println(m)
				return true
			}
		}
	}
	fmt.Printf("Invalid: ")
	fmt.Println(m)
	return false
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var validRecords = 0
	var currentRecord []string
	// var allRecords []string

	for scanner.Scan() {
		thisLine := scanner.Text()
		if len(thisLine) > 1 {
			parts := strings.Split(thisLine, " ")
			currentRecord = append(currentRecord, parts...)
		} else {
			if validate(currentRecord) {
				validRecords++
			}
			currentRecord = make([]string, 0)
		}
	}

	if validate(currentRecord) {
		validRecords++
	}

	file.Close()

	fmt.Println(validRecords)

}
