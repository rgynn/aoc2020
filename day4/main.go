package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"strings"
)

type passport struct{}

func main() {

	passports, err := readInput()
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	log.Println(passports)
}

func readInput() ([]passport, error) {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		return nil, err
	}

	var passports []passport

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		if len(lineSplit) < 1 {
			continue
		}
		for _, kv := range lineSplit {
			log.Println(kv)
		}
	}

	return passports, nil
}
