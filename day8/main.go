package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var instructions []instruction
var changedOnce bool
var pos, steps, one, two int

type instruction struct {
	Operation string
	Argument  int
	Executed  bool
}

func (i *instruction) One() {

	if i.Executed {
		return
	}

	switch i.Operation {
	case "acc":
		one += i.Argument
		pos++
	case "jmp":
		pos += i.Argument
	case "nop":
		pos++
	}

	i.Executed = true

	instructions[pos].One()
}

func (i *instruction) Two(program []instruction) bool {

	if steps >= len(program) {
		return true
	}

	program[pos].Executed = true

	switch i.Operation {
	case "acc":
		two += i.Argument
		pos++
	case "jmp":
		pos += i.Argument
	case "nop":
		pos++
	}

	steps++

	if pos == len(program) {
		return false
	}

	return program[pos].Two(program)
}

func main() {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	if err := readInput(file); err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	instructions[0].One()
	log.Printf("Answer to part one: %d", one)

	for i := range instructions {

		program := reset()

		switch program[i].Operation {
		case "jmp":
			program[i].Operation = "nop"
		case "nop":
			program[i].Operation = "jmp"
		}

		if infiniteLoop := program[0].Two(program); infiniteLoop {
			continue
		} else {
			break
		}
	}

	// 1521
	// 1695 too high
	log.Printf("Answer to part two: %d", two)
}

func reset() []instruction {

	pos = 0
	two = 0
	steps = 0

	for i := range instructions {
		instructions[i].Executed = false
	}

	result := make([]instruction, len(instructions))
	copy(result, instructions)

	return result
}

func readInput(file []byte) error {

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {

		line := scanner.Text()
		lineSplit := strings.Split(line, " ")

		arg, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			return err
		}

		instructions = append(instructions, instruction{Operation: lineSplit[0], Argument: arg})
	}

	return nil
}
