package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

	column := 0
	trees := 0
	rightIncrement := 1
	downIncrement := 2
	for counter, line := range txtlines {
		// fmt.Printf("%s %d\n", string(line[column % len(line)]), column)
		if counter%downIncrement == 0 {
			if string(line[column%len(line)]) == "#" {
				// fmt.Printf("Tree at %d %d\n", counter, column)
				trees++
			}
			column = column + rightIncrement
		}
	}

	fmt.Printf("Trees: %d\n", trees)
}
