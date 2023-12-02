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
		first := ""
		last := ""
		for _, c := range line {
			if _, err := strconv.Atoi(string(c)); err == nil {
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

func unSpellOut(s string) int {
	type digitIndex struct {
		digit string
		index int
	}
	firstIndex := map[string]int{}
	lastIndex := map[string]int{}
	digits := map[string]string{
		"one":   "1",
		"1":     "1",
		"two":   "2",
		"2":     "2",
		"three": "3",
		"3":     "3",
		"four":  "4",
		"4":     "4",
		"five":  "5",
		"5":     "5",
		"six":   "6",
		"6":     "6",
		"seven": "7",
		"7":     "7",
		"eight": "8",
		"8":     "8",
		"nine":  "9",
		"9":  "9",
	}

	for k := range digits {
		f := strings.Index(s, k)
		if f != -1 {
			firstIndex[k] = f
		}
		l := strings.LastIndex(s, k)
		if l != -1 {
			lastIndex[k] = l
		}
	}
	first := digitIndex{digit: "", index: 5000}
	last := digitIndex{digit: "", index: -1}
	for k, v := range firstIndex {
		if v < first.index {
			first = digitIndex{digit: k, index: v}
		}
	}
	for k, v := range lastIndex {
		if v > last.index {
			last = digitIndex{digit: k, index: v}
		}
	}
	if first.index == last.index {
		value, _ := strconv.Atoi(digits[first.digit])
		fmt.Printf("%40s %v %v %d\n", s, first, last, value)
		return value
	}
	value, _ := strconv.Atoi(strings.Join([]string{digits[first.digit], digits[last.digit]}, ""))
	fmt.Printf("%40s %v %v %d\n", s, first, last, value)
	return value
}

// 1tbbsmdhtwonedtt
func part2(input []string) (solution int) {
	sum := 0
	for _, line := range input {
		value := unSpellOut(line)
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
