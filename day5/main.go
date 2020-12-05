package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
)

func main() {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	var one, two int
	seats := [128][8]bool{}

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {
		instructions := scanner.Text()
		y, x := getSeatPosition(instructions)
		seats[y][x] = true
		seatID := getSeatID(y, x)
		if one < seatID {
			one = seatID
		}
	}

	for y, row := range seats {
		for x, seat := range row {
			if !seat {
				if x > 0 && seats[y][x-1] &&
					x < 7 && seats[y][x+1] {
					two = getSeatID(y, x)
				}
			}
		}
	}

	log.Printf("Answer to part one: %d", one)
	log.Printf("Answer to part two: %d", two)
}

func getSeatPosition(instructions string) (int, int) {

	minY := 0
	maxY := 127

	minX := 0
	maxX := 7

	for _, instruction := range instructions {
		switch instruction {
		case 'F', 'B':
			minY, maxY = getRange(instruction, minY, maxY)
		case 'L', 'R':
			minX, maxX = getRange(instruction, minX, maxX)
		}
	}

	return minY, minX
}

func getSeatID(row, col int) int {
	return row*8 + col
}

func getRange(instruction rune, min, max int) (int, int) {
	half := (max - min + 1) / 2
	switch instruction {
	case 'F', 'L':
		return min, max - half
	default:
		return min + half, max
	}
}
