package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []int

	for scanner.Scan() {
		thisLine, _ := strconv.Atoi(scanner.Text())
		txtlines = append(txtlines, thisLine)
	}

	file.Close()

Mainloop:
	for _, line := range txtlines {
		for _, line2 := range txtlines {
			for _, line3 := range txtlines {
			if line+line2+line3 == 2020 {
				fmt.Printf("Found %d %d %d... Product %d\n", line, line2, line3, line * line2 * line3)
				break Mainloop
			}
		}
		}
	}
}
