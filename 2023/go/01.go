package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	example = "e"
	puzzle  = "p"
)

type InputType string

func Input(day int, part int, inputType InputType) (input []string, err error) {

	inputFilename := fmt.Sprintf("../input/%02d-%s%di.txt", day, inputType, part)
	file, err := os.Open(inputFilename)
	if err != nil {
		return input, fmt.Errorf("failed to open %s: %w", inputFilename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return input, fmt.Errorf("failed to scan by lines: %w", err)
	}
	return input, nil

}

func InputExample(day int, part int) (input []string, err error) {
	return Input(day, part, example)
}

func InputPuzzle(day int, part int) (input []string, err error) {
	return Input(day, part, puzzle)
}

func part1(input []string) (solution int) {
	sum := 0
	for _, line := range input {
		//num := []string{}
		first := ""
		last := ""
		for _, c := range line {
			if _, err := strconv.Atoi(string(c)); err == nil {
				//num = append(num, string(c))
				if first == "" {
					first = string(c)
				}
				last = string(c)
			}
		}
		value, err := strconv.Atoi(strings.Join([]string{first, last}, ""))
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + value
	}
	return sum
}

func unSpellOut(s string) string {
	digits := map[string]string{
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}

	for k := range digits {
		s = strings.ReplaceAll(s, k, digits[k])
	}
	return s
}

// 1tbbsmdhtwonedtt
func part2(input []string) (solution int) {
	sum := 0
	for _, line := range input {
		fixedLine := unSpellOut(line)
		first := ""
		last := ""
		for _, c := range fixedLine {
			if _, err := strconv.Atoi(string(c)); err == nil {
				//num = append(num, string(c))
				if first == "" {
					first = string(c)
				} else {
					last = string(c)
				}
			}
		}
		value, err := strconv.Atoi(strings.Join([]string{first, last}, ""))
		fmt.Printf("%40s %40s %10d\n", line, fixedLine, value)
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + value
	}
	return sum
}

func main() {
	e1, err := InputExample(1, 1)
	if err != nil {
		log.Fatal(err)
	}
	p1, err := InputPuzzle(1, 1)
	if err != nil {
		log.Fatal(err)
	}
	e2, err := InputExample(1, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Example 1:", part1(e1))
	fmt.Println("Puzzle  1:", part1(p1))
	fmt.Println("Example 2:", part2(e2))
	p2, err := InputPuzzle(1, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Puzzle  2:", part2(p2))

}
