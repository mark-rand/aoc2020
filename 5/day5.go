package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"sort"
)

//DecodeSeatNumber returns row and seat
func DecodeSeatNumber(seatStr string) (seatNumber int64) {
	seatStr = strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1").Replace(seatStr)
	row64, _ := strconv.ParseInt(seatStr[:7], 2, 8)
	seat64, _ := strconv.ParseInt(seatStr[7:], 2, 8)
	seatNumber = row64*8 + seat64
	return
}

func getInput() []string {
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
	return txtlines
}

func main() {
	input := getInput()
	ids := []int64{}
	for _, line := range input {
		seat := DecodeSeatNumber(line)
		ids = append(ids, seat)
		fmt.Println(seat)
	}

	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	fmt.Println("Max ID =", ids[len(ids)-1])

	for idx, id := range ids {
		if ids[idx+1] != id+1 {
			fmt.Println("My ID =", id+1)
			return
		}
	}
}
