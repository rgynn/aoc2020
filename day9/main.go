package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
)

func main() {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	stream, err := readInput(file)
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	var one int
	for i := 25; i < len(stream); i++ {
		preamble := slidingWindow(i, 25, stream)
		if !valid(stream[i], preamble) {
			one = stream[i]
			break
		}
	}

	log.Printf("Answer to part one: %d", one)

	var two int
	var matches []int

streamloop:
	for i, x := range stream {

		if x == one {
			matches = []int{}
			continue
		}

		// step backwards from i in stream
		for y := i; y >= 0; y-- {

			matches = append(matches, stream[y])

			var sum int
			for _, z := range matches {

				sum += z

				if sum < one {
					continue
				}

				if sum == one {
					sort.Slice(matches, func(i, j int) bool { return matches[i] > matches[j] })
					two = matches[0] + matches[len(matches)-1]
					break streamloop
				}

				if sum > one {
					matches = []int{}
				}
			}
		}
	}

	log.Printf("Answer to part two: %d", two)
}

func valid(input int, preamble []int) bool {

	for _, px := range preamble {
		for _, py := range preamble {
			if px == py {
				continue
			}
			if px+py == input {
				return true
			}
		}
	}

	return false
}

func slidingWindow(offset, size int, input []int) []int {

	if len(input) <= size {
		return input
	}

	result := []int{}

	for i := 0; i < size; i++ {
		if offset+i == len(input) {
			return result
		}
		result = append(result, input[(offset-size)+i])
	}

	return result
}

func readInput(file []byte) (stream []int, err error) {

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {
		line := scanner.Text()

		l, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		stream = append(stream, l)
	}

	return
}
