package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
)

const (
	slopeMaxX = 31
)

type toboggan struct {
	x, y       int
	velX, velY int
	hits       int
}

func (t *toboggan) Step() {
	t.StepX()
	t.StepY()
}

func (t *toboggan) StepX() {
	if (t.x + t.velX) >= slopeMaxX {
		t.x = (t.x + t.velX) - slopeMaxX
	} else {
		t.x += t.velX
	}
}

func (t *toboggan) StepY() {
	t.y += t.velY
}

func (t *toboggan) CheckTree(slope []string) bool {

	if t.y >= len(slope) {
		return false
	}

	return slope[t.y][t.x] == '#'
}

func (t *toboggan) Go(slope []string) {

	for t.y < len(slope) {
		t.Step()
		if hit := t.CheckTree(slope); hit {
			t.hits++
		}
	}

	return
}

func main() {

	slope, err := makeSlope()
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	// exercise one
	one := toboggan{velX: 3, velY: 1}
	one.Go(slope)

	// exercise two
	var two int
	for i, t := range []toboggan{
		toboggan{velX: 1, velY: 1},
		toboggan{velX: 3, velY: 1},
		toboggan{velX: 5, velY: 1},
		toboggan{velX: 7, velY: 1},
		toboggan{velX: 1, velY: 2},
	} {
		t.Go(slope)
		if i == 0 {
			two += t.hits
		} else {
			two *= t.hits
		}
	}

	log.Printf("Answer to exercise one: %d", one.hits)
	log.Printf("Answer to exercise two: %d", two)
}

func makeSlope() ([]string, error) {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		return nil, err
	}

	slope := []string{}

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {
		slope = append(slope, scanner.Text())
	}

	return slope, nil
}
