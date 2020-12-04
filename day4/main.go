package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

type passport struct {
	BirthYear      string // byr
	IssueYear      string // iyr
	ExpirationYear string // eyr
	Height         string // hgt
	HairColor      string // hcl
	EyeColor       string // ecl
	PassportID     string // pid
	CountryID      string // cid
}

func (pass *passport) ValidOne() error {

	if pass.BirthYear == "" {
		return errors.New("no byr provided")
	}

	if pass.IssueYear == "" {
		return errors.New("no iyr provided")
	}

	if pass.ExpirationYear == "" {
		return errors.New("no eyr provided")
	}

	if pass.Height == "" {
		return errors.New("no hgt provided")
	}

	if pass.HairColor == "" {
		return errors.New("no hcl provided")
	}

	if pass.EyeColor == "" {
		return errors.New("no ecl provided")
	}

	if pass.PassportID == "" {
		return errors.New("no pid provided")
	}

	return nil
}

func (pass *passport) ValidTwo() error {

	// Issue Year byr

	if pass.BirthYear == "" {
		return errors.New("no byr provided")
	}

	if byr, err := strconv.Atoi(pass.BirthYear); err != nil {
		return fmt.Errorf("%w: failed to validate byr as integer", err)
	} else if byr < 1920 || byr > 2002 {
		return fmt.Errorf("byr needs to be between 1920 and 2002, got: %d", byr)
	}

	// Issue Year iyr

	if pass.IssueYear == "" {
		return errors.New("no iyr provided")
	}

	if iyr, err := strconv.Atoi(pass.IssueYear); err != nil {
		return fmt.Errorf("%w: failed to validate iyr as integer", err)
	} else if iyr < 2010 || iyr > 2020 {
		return fmt.Errorf("iyr needs to be between 2010 and 2020, got: %d", iyr)
	}

	// Expiration Year eyr

	if pass.ExpirationYear == "" {
		return errors.New("no eyr provided")
	}

	if eyr, err := strconv.Atoi(pass.ExpirationYear); err != nil {
		return fmt.Errorf("%w: failed to validate eyr as integer", err)
	} else if eyr < 2020 || eyr > 2030 {
		return fmt.Errorf("eyr needs to be between 2020 and 2030, got: %d", eyr)
	}

	// Height hgt

	if pass.Height == "" {
		return errors.New("no hgt provided")
	}

	hascm := strings.HasSuffix(pass.Height, "cm")
	if hascm {
		if cm, err := strconv.Atoi(strings.TrimSuffix(pass.Height, "cm")); err != nil {
			return fmt.Errorf("%w: failed to validate hgt as integer", err)
		} else if cm < 150 || cm > 193 {
			return fmt.Errorf("eyr needs to be between 150cm and 193cm, got: %d", cm)
		}
	}

	hasin := strings.HasSuffix(pass.Height, "in")
	if hasin {
		if in, err := strconv.Atoi(strings.TrimSuffix(pass.Height, "in")); err != nil {
			return fmt.Errorf("%w: failed to validate hgt as integer", err)
		} else if in < 59 || in > 76 {
			return fmt.Errorf("eyr needs to be between 59in and 76in, got: %d", in)
		}
	}

	if !hascm && !hasin {
		return fmt.Errorf("hgt needs to have suffix cm or in, got: %s", pass.Height)
	}

	// Hair Color hcl

	if pass.HairColor == "" {
		return errors.New("no hcl provided")
	}

	if !strings.HasPrefix(pass.HairColor, "#") {
		return fmt.Errorf("hgt needs to have prefix #, got: %s", pass.HairColor)
	}

	if utf8.RuneCountInString(pass.HairColor) != 7 {
		return fmt.Errorf("hcl needs to be 7 characters long")
	}

	// Eye color ecl

	if pass.EyeColor == "" {
		return errors.New("no ecl provided")
	}

	switch pass.EyeColor {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		break
	default:
		return fmt.Errorf("ecl needs to be either: amb, blu, brn, gry, grn, hzl, oth, got: %s", pass.EyeColor)
	}

	// PassportID pid

	if pass.PassportID == "" {
		return errors.New("no pid provided")
	}

	if pid := utf8.RuneCountInString(pass.PassportID); pid != 9 {
		return fmt.Errorf("pid needs to be 9 digits long, got: %d", pid)
	}

	if pid, err := strconv.Atoi(pass.PassportID); err != nil {
		return fmt.Errorf("%w: failed to validate pid as integer, got: %d", err, pid)
	}

	return nil
}

func main() {

	passports, err := readInput("input")
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	var one, two int
	for _, pass := range passports {
		if err := pass.ValidOne(); err == nil {
			one++
		}
		if err := pass.ValidTwo(); err != nil {
			log.Printf("%+v", err)
			continue
		}
		two++
	}

	log.Printf("Valid passports for part one: %d", one)
	log.Printf("Valid passports for part two: %d", two)
}

func readInput(inputfile string) ([]passport, error) {

	file, err := ioutil.ReadFile(inputfile)
	if err != nil {
		return nil, err
	}

	var passports []passport
	var pass passport

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {

		line := scanner.Text()
		if line == "" {
			passports = append(passports, pass)
			pass = passport{}
			continue
		}

		lineSplit := strings.Split(line, " ")

		for _, kv := range lineSplit {

			kvSplit := strings.Split(kv, ":")
			if len(kvSplit) < 2 {
				continue
			}

			switch kvSplit[0] {
			case "byr":
				pass.BirthYear = kvSplit[1]
			case "iyr":
				pass.IssueYear = kvSplit[1]
			case "eyr":
				pass.ExpirationYear = kvSplit[1]
			case "hgt":
				pass.Height = kvSplit[1]
			case "hcl":
				pass.HairColor = kvSplit[1]
			case "ecl":
				pass.EyeColor = kvSplit[1]
			case "pid":
				pass.PassportID = kvSplit[1]
			case "cid":
				pass.CountryID = kvSplit[1]
			}
		}
	}

	passports = append(passports, pass)

	return passports, nil
}
