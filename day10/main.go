package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
)

type adapter struct {
	output int
}

func (a *adapter) ValidInput(input int) bool {

	if input > a.output {
		return false
	}

	if a.output-input > 3 {
		return false
	}

	return true
}

func (a *adapter) Diff(input int) int {
	return a.output - input
}

func main() {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	adapters, err := readInput(file)
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	builtIn := getBuiltInAdapter(adapters)

	log.Printf("Answer to part one: %d", partOne(builtIn, adapters))
	log.Printf("Answer to part two: %d", partTwo(builtIn, adapters))
}

func partTwo(builtIn adapter, adapters []adapter) int {

	adapters = append([]adapter{adapter{}}, adapters...)
	adapters = append(adapters, builtIn)

	// Memoization map used to calculate number of combinations per step
	m := map[int]int{0: 1}
	for _, a := range adapters {
		for n := -3; n <= -1; n++ {
			m[a.output] += m[a.output+n]
		}
	}

	return m[builtIn.output]
}

func partOne(builtIn adapter, adapters []adapter) int {

	var joltage, ones, threes int
	var candidates, chain []adapter

	bag := make([]adapter, len(adapters))
	copy(bag, adapters)

	for len(chain) < len(adapters) {
		candidates, bag = popCandidatesFromBag(joltage, bag)
		for _, c := range candidates {
			chain = append(chain, c)
			switch c.Diff(joltage) {
			case 1:
				ones++
			case 3:
				threes++
			}
			joltage = c.output
		}
	}

	switch builtIn.Diff(joltage) {
	case 1:
		ones++
	case 3:
		threes++
	}

	return ones * threes
}

func popLowest(candidates []adapter) (adapter, []adapter) {
	return candidates[0], candidates[1:]
}

func candidatesFindDiffs(joltage, diff int, candidates []adapter) int {
	var result int
	for _, a := range candidates {
		if a.Diff(joltage) == diff {
			result++
		}
	}
	return result
}

func candidatesMaxJoltage(candidates []adapter) int {
	var joltage int
	for _, a := range candidates {
		if joltage < a.output {
			joltage = a.output
		}
	}
	return joltage
}

func popCandidatesFromBag(joltage int, adapters []adapter) (candidates, bag []adapter) {

	bag = make([]adapter, len(adapters))
	copy(bag, adapters)

	for i := len(adapters) - 1; i >= 0; i-- {
		if adapters[i].ValidInput(joltage) {
			candidates = append(candidates, adapters[i])
			bag = append(bag[:i], bag[i+1:]...)
		}
	}

	sort.Slice(candidates, func(i, j int) bool { return candidates[i].output < candidates[j].output })

	return candidates, bag
}

func getBuiltInAdapter(adapters []adapter) adapter {
	var highest int
	for _, a := range adapters {
		if a.output > highest {
			highest = a.output
		}
	}
	return adapter{output: highest + 3}
}

func readInput(file []byte) ([]adapter, error) {

	result := []adapter{}

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		result = append(result, adapter{output: line})
	}

	sort.Slice(result, func(i, j int) bool { return result[i].output < result[j].output })

	return result, nil
}
