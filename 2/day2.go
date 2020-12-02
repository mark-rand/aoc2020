package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// IsPasswordValid is password valid?
func IsPasswordValid(passwordSpec string) bool {
	re := regexp.MustCompile("(\\d+)-(\\d+) (\\w): (\\w+)")
	res := re.FindAllStringSubmatch(passwordSpec, 1)[0]
	minOccurence, _ := strconv.Atoi(res[1])
	maxOccurence, _ := strconv.Atoi(res[2])
	char := res[3]
	password := res[4]
	occurences := strings.Count(password, char)
	return occurences >= minOccurence && occurences <= maxOccurence
}

// IsPasswordValid2 is password valid?
func IsPasswordValid2(passwordSpec string) bool {
	re := regexp.MustCompile("(\\d+)-(\\d+) (\\w): (\\w+)")
	res := re.FindAllStringSubmatch(passwordSpec, 1)[0]
	firstPosition, _ := strconv.Atoi(res[1])
	secondPosition, _ := strconv.Atoi(res[2])
	char := res[3]
	password := res[4]
	return (char == string(password[firstPosition-1]) != (char == string(password[secondPosition-1])))
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		thisLine := scanner.Text()
		txtlines = append(txtlines, thisLine)
	}

	file.Close()

	counter := 0
	for _, line := range txtlines {
		if IsPasswordValid2(line) {
			counter++
		}
	}
	fmt.Println(counter)
}
