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

	type group struct {
		members int
		answers map[rune]int
	}

	groups := []group{}
	grp := group{members: 0, answers: map[rune]int{}}

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || line == "\n" {
			groups = append(groups, grp)
			grp = group{members: 0, answers: map[rune]int{}}
			continue
		}

		for _, r := range line {
			if _, ok := grp.answers[r]; ok {
				grp.answers[r]++
			} else {
				grp.answers[r] = 1
			}
		}

		grp.members++
	}

	groups = append(groups, grp)

	var one, two int
	for _, grp := range groups {
		for _, answers := range grp.answers {
			one++
			if answers == grp.members {
				two++
			}
		}
	}

	log.Printf("Answer to part one: %d", one)
	log.Printf("Answer to part two: %d", two)
}
