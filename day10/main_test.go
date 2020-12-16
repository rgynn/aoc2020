package main

import (
	"io/ioutil"
	"testing"
)

func TestPartOneTwo(t *testing.T) {

	file, err := ioutil.ReadFile("input_test_2")
	if err != nil {
		t.Fatal(err)
	}

	adapters, err := readInput(file)
	if err != nil {
		t.Fatal(err)
	}

	builtIn := getBuiltInAdapter(adapters)

	one := partOne(builtIn, adapters)
	if one != 220 {
		t.Fatalf("expected answer to part one to be: %d, got: %d", 220, one)
	}
}

func TestPartOneOne(t *testing.T) {

	file, err := ioutil.ReadFile("input_test_1")
	if err != nil {
		t.Fatal(err)
	}

	adapters, err := readInput(file)
	if err != nil {
		t.Fatal(err)
	}

	builtIn := getBuiltInAdapter(adapters)

	one := partOne(builtIn, adapters)
	if one != 35 {
		t.Fatalf("expected answer to part one to be: %d, got: %d", 35, one)
	}
}

func TestGetBuiltInAdapter(t *testing.T) {

	adapters := []adapter{
		adapter{output: 1},
		adapter{output: 2},
		adapter{output: 3},
		adapter{output: 4},
	}

	builtin := getBuiltInAdapter(adapters)

	if builtin.output != 7 {
		t.Fatalf("expected built-in adapter output to be: %d, got: %d", 7, builtin.output)
	}
}

func TestPopCandidatesFromBag(t *testing.T) {

	adapters := []adapter{
		adapter{output: 1},
		adapter{output: 2},
		adapter{output: 3},
		adapter{output: 4},
	}

	candidates, bag := popCandidatesFromBag(0, adapters)

	if len(candidates) != 3 {
		t.Fatalf("expected num candidates to be: %d, got: %d", 3, len(candidates))
	}

	if len(bag) != 1 {
		t.Fatalf("expected num adapters in bag to be: %d, got: %d", 1, len(bag))
	}
}

func TestCandidatesMaxJoltage(t *testing.T) {

	candidates := []adapter{
		adapter{output: 1},
		adapter{output: 2},
		adapter{output: 3},
	}

	if max := candidatesMaxJoltage(candidates); max != 3 {
		t.Fatalf("expected max joltage from candidates to be: %d, got: %d", 3, max)
	}
}

func TestCandidatesFindDiffs(t *testing.T) {

	candidates := []adapter{
		adapter{output: 1},
		adapter{output: 2},
		adapter{output: 3},
	}

	if ones := candidatesFindDiffs(0, 1, candidates); ones != 1 {
		t.Fatalf("expected diffs from candidates to be: %d, got: %d", 1, ones)
	}

	if threes := candidatesFindDiffs(0, 3, candidates); threes != 1 {
		t.Fatalf("expected diffs from candidates to be: %d, got: %d", 1, threes)
	}
}
