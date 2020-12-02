package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type policy struct {
	MinOccurance int
	MaxOccurance int
	Letter       rune
	Password     string
}

func (p *policy) String() string {
	return fmt.Sprintf("MinOccurance: %d, MaxOccurance: %d, Letter: %s, Password: %s", p.MinOccurance, p.MaxOccurance, string(p.Letter), p.Password)
}

func (p *policy) ValidForOne() bool {
	var sum int
	for _, l := range p.Password {
		if l == p.Letter {
			sum++
		}
	}
	if sum < p.MinOccurance || sum > p.MaxOccurance {
		return false
	}
	return true
}

func (p *policy) ValidForTwo() bool {

	var first, second bool

	if rune(p.Password[p.MinOccurance-1]) == p.Letter {
		log.Printf("First rule match with toboggan index: %d, slice index: %d, letter from password: %s is the same as policy letter: %s in password: %s\n", p.MinOccurance, p.MinOccurance-1, string(p.Password[p.MinOccurance-1]), string(p.Letter), p.Password)
		first = true
	}

	if rune(p.Password[p.MaxOccurance-1]) == p.Letter {
		log.Printf("Second rule match with toboggan index: %d, slice index: %d, letter from Password: %s is the same as policy letter: %s in password: %s\n", p.MaxOccurance, p.MaxOccurance-1, string(p.Password[p.MaxOccurance-1]), string(p.Letter), p.Password)
		second = true
	}

	if first && second {
		return false
	}

	return first || second
}

func main() {

	input, err := readInput()
	if err != nil {
		log.Fatalf("FATAL: %+v", err)
	}

	var one, two []*policy
	for _, p := range input {
		if p.ValidForOne() {
			one = append(one, p)
		}
		if p.ValidForTwo() {
			two = append(two, p)
		}
	}

	log.Printf("Got result for exercise one: %d out of %d passwords \n", len(one), len(input))
	log.Printf("Got result for exercise two: %d out of %d passwords \n", len(two), len(input))
}

func readInput() ([]*policy, error) {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		return nil, err
	}

	var result []*policy

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {

		line := scanner.Text()
		policy, err := extractPolicy(line)
		if err != nil {
			return nil, err
		}

		result = append(result, policy)
	}

	return result, nil
}

func extractPolicy(line string) (*policy, error) {

	split := strings.Split(line, " ")
	minmax := strings.Split(split[0], "-")

	min, err := strconv.Atoi(minmax[0])
	if err != nil {
		return nil, err
	}

	max, err := strconv.Atoi(minmax[1])
	if err != nil {
		return nil, err
	}

	letter := []rune(strings.TrimSuffix(split[1], ":"))[0]
	password := split[2]

	return &policy{
		MinOccurance: min,
		MaxOccurance: max,
		Letter:       letter,
		Password:     password,
	}, nil
}
