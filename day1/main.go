package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {

	input, err := readInput()
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	// Exercise one
	one := getMatchesForOne(input)
	log.Printf("Got matches for excercise one: %v \n", one)
	log.Printf("Result of exercise one: %v \n", getSum(one))

	// Exercise two
	two := getMatchesForTwo(input)
	log.Printf("Got matches for excercise two: %v \n", two)
	log.Printf("Result of exercise two: %v \n", getSum(two))

}

func readInput() ([]int, error) {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		return nil, err
	}

	var result []int

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}

func getMatchesForOne(input []int) (matches [][]int) {
	for _, i := range input {
		for _, j := range input {
			if isTwentyTwenty(i, j) {
				if !alreadyInMatches(j, matches) {
					matches = append(matches, []int{i, j})
				}
			}
		}
	}
	return
}

func getMatchesForTwo(input []int) (matches [][]int) {
	for _, i := range input {
		for _, j := range input {
			for _, k := range input {
				if isTwentyTwenty(i, j, k) {
					if !alreadyInMatches(k, matches) {
						matches = append(matches, []int{i, j, k})
					}
				}
			}
		}
	}
	return
}

func isTwentyTwenty(input ...int) bool {
	var sum int
	for _, i := range input {
		sum += i
	}
	return sum == 2020
}

func alreadyInMatches(input int, matches [][]int) bool {
	for _, touple := range matches {
		for _, i := range touple {
			if i == input {
				return true
			}
		}
	}
	return false
}

func getSum(matches [][]int) (sum int) {
	for _, match := range matches {
		for _, m := range match {
			if sum == 0 {
				sum += m
			} else {
				sum *= m
			}
		}
	}
	return
}
