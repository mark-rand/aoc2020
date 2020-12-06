package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func countRecords(currentRecord []string, records int) (answerCount int) {
	m := make(map[string]int)
	for _, entry := range currentRecord {
		m[entry]++
	}
	fmt.Println(len(m))
	fmt.Println(m)
	answerCount = 0

	for _, count := range(m) {
		if count == records {
			answerCount++
		}
	}
	return
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var answerCount = 0
	var records = 0
	var currentRecord []string

	for scanner.Scan() {
		thisLine := scanner.Text()
		if len(thisLine) >= 1 {
			records++
			parts := strings.Split(thisLine, "")
			currentRecord = append(currentRecord, parts...)
		} else {
			answerCount += countRecords(currentRecord, records)
			currentRecord = make([]string, 0)
			records = 0
		}
	}

	file.Close()

	fmt.Println(answerCount)

}
