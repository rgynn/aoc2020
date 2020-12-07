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
	Color  string
	Amount int
	Rules  []bag
}

func (b *bag) Contains(color string) bool {
	for _, bg := range b.Rules {
		if bg.Color == color {
			return true
		}
	}
	return false
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

	containsDirectly := []bag{}

	for _, bg := range bags {
		if bg.Contains("shiny gold") {
			containsDirectly = append(containsDirectly, bg)
			log.Printf("%s bags contain shiny gold bags directly", bg.Color)
		}
	}

	containsInDirectly := []bag{}
	for _, bg := range bags {
		for _, cdbg := range containsDirectly {
			if bg.Contains(cdbg.Color) {
				containsInDirectly = append(containsInDirectly, bg)
				log.Printf("%s bags contain shiny gold bags indirectly", bg.Color)
			}
		}
	}

	one := len(containsDirectly) + len(containsInDirectly)

	log.Printf("Answer to part one: %d", one)
}

func readInput(file []byte) ([]bag, error) {

	bags := []bag{}

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), ".")
		ruleSplit := strings.Split(line, " bags contain ")

		color := ruleSplit[0]
		rules := strings.Split(ruleSplit[1], ", ")

		bg := bag{Color: color, Rules: []bag{}}

		if !(ruleSplit[1] == "no other bags") {
			for _, rule := range rules {

				rule = strings.TrimSuffix(rule, " bags")
				rule = strings.TrimSuffix(rule, " bag")
				ruleSplit = strings.Split(rule, " ")

				num, err := strconv.Atoi(ruleSplit[0])
				if err != nil {
					return nil, err
				}

				color := fmt.Sprintf("%s %s", ruleSplit[1], ruleSplit[2])
				bg.Rules = append(bg.Rules, bag{Amount: num, Color: color})
			}
		}

		bags = append(bags, bg)
	}

	return bags, nil
}
