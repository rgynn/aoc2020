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

type bag struct {
	Color string
	Rules []rule
}

type rule struct {
	Amount int
	Color  string
}

func (b *bag) Contains(color string) bool {
	for _, r := range b.Rules {
		if r.Color == color {
			return true
		}
	}
	return false
}

func (b *bag) ContainsAmount(color string) int {
	for _, r := range b.Rules {
		if r.Color == color {
			return r.Amount
		}
	}
	return 0
}

func main() {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	bags, err := readInput(file)
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	one := partOne(bags)
	two := partTwo(bags)

	log.Printf("Answer to part one: %d", one)
	log.Printf("Answer to part two: %d", two)
}

func partOne(bags map[string]bag) int {
	sum := 0
	for _, bg := range bags {
		if canContainGold(bg.Color, bags) && bg.Color != "shiny gold" {
			sum++
		}
	}
	return sum
}

func partTwo(bags map[string]bag) int {
	return countBags("shiny gold", bags)
}

func canContainGold(color string, bags map[string]bag) bool {
	if color == "shiny gold" {
		return true
	}
	for _, ib := range bags[color].Rules {
		if canContainGold(ib.Color, bags) {
			return true
		}
	}
	return false
}

func countBags(color string, bags map[string]bag) int {
	sum := 0
	for _, r := range bags[color].Rules {
		subCount := countBags(r.Color, bags)
		sum += r.Amount + r.Amount*subCount
	}
	return sum
}

func readInput(file []byte) (map[string]bag, error) {

	bags := map[string]bag{}

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), ".")
		ruleSplit := strings.Split(line, " bags contain ")

		color := ruleSplit[0]
		rules := strings.Split(ruleSplit[1], ", ")

		bg := bag{Color: color}

		if !(ruleSplit[1] == "no other bags") {
			for _, r := range rules {

				r = strings.TrimSuffix(r, " bags")
				r = strings.TrimSuffix(r, " bag")
				ruleSplit = strings.Split(r, " ")

				num, err := strconv.Atoi(ruleSplit[0])
				if err != nil {
					return nil, err
				}

				color := fmt.Sprintf("%s %s", ruleSplit[1], ruleSplit[2])
				bg.Rules = append(bg.Rules, rule{Amount: num, Color: color})
			}
		}

		bags[bg.Color] = bg
	}

	return bags, nil
}
