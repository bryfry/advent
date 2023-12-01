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
	found := false
	digits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for k, v := range digits {
		s = strings.ReplaceAll(s, k, v)
		strings.Index()
	}
	return s
}

func part2(input []string) (solution int) {
	sum := 0
	for _, line := range input {
		line = unSpellOut(line)
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
		fmt.Println(line, value)
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
	//p2, err := InputPuzzle(1, 2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println("Example 1:", part1(e1))
	fmt.Println("Puzzle  1:", part1(p1))
	fmt.Println("Example 2:", part2(e2))
	//fmt.Println("Puzzle  2:", part2(p2))
	
}













