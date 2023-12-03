package main

import (
	"advent2023/internal/aoc"
	"strconv"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

type cord struct {
	x int
	y int
}

type board struct {
	grid     map[cord]uuid.UUID
	parts    map[uuid.UUID]part
	included map[uuid.UUID]part
	symbols  []symbol
	gears    []gear
}

type part struct {
	id     uuid.UUID
	number int
}

type symbol struct {
	repr string
	cord cord
}

type gear struct {
	symbol
	ratio int
}


func parseBoard(input []string) (b board) {
	b = board{
		grid:     map[cord]uuid.UUID{},
		parts:    map[uuid.UUID]part{},
		included: map[uuid.UUID]part{},
		symbols:  []symbol{},
		gears:    []gear{},
	}

	for y, row := range input {
		var partSB strings.Builder
		partID := uuid.New()
		for x, col := range row {

			if unicode.IsDigit(col) {
				partSB.WriteRune(col)
				b.grid[cord{x, y}] = partID
				// only continue if not at end of row
				if x < len(row)-1 {
					continue
				}
			}

			// non-digit or end or row after digits parsed
			// part number completed, store it
			if partSB.Len() > 0 {
				partNumber, _ := strconv.Atoi(partSB.String())
				b.parts[partID] = part{id: partID, number: partNumber}
				partID = uuid.New()
				partSB.Reset()
				if x == len(row)-1 {
					continue
				}
			}

			if col == '.' {
				continue
			}

			// symbol remaning, store it
			b.symbols = append(b.symbols, symbol{repr: string(col), cord: cord{x, y}})
		}
	}
	for _, s := range b.symbols {

		neighbors := map[uuid.UUID]part{}
		for _, c := range s.cord.neighbors() {
			if id, ok := b.grid[c]; ok {
				b.included[id] = b.parts[id]
				neighbors[id] = b.parts[id]
			}
		}
		if s.repr == "*" && len(neighbors) == 2 {
			ratio := 1
			for _, p := range neighbors {
				ratio = ratio * p.number
			}
			b.gears = append(b.gears, gear{
				symbol: s,
				ratio:  ratio,
			})
		}
	}
	return b
}

func (c cord) neighbors() (cords []cord) {
	cords = []cord{
		{c.x - 1, c.y - 1},
		{c.x - 1, c.y},
		{c.x - 1, c.y + 1},
		{c.x, c.y - 1},
		{c.x, c.y + 1},
		{c.x + 1, c.y - 1},
		{c.x + 1, c.y},
		{c.x + 1, c.y + 1},
	}
	return cords
}

func solvePart1(input []string) (solution int) {
	b := parseBoard(input)
	solution = 0
	for _, p := range b.included {
		solution += p.number
	}
	return solution
}

func solvePart2(input []string) (solution int) {
	b := parseBoard(input)
	solution = 0
	for _, s := range b.gears {
		solution += s.ratio
	}
	return solution
}

func main() {
	day := aoc.Day3
	aoc.SolveExample(day, aoc.Part1, solvePart1, 4361)
	aoc.SolvePuzzle(day, aoc.Part1, solvePart1, 533784)
	aoc.SolveExample(day, aoc.Part2, solvePart2, 467835)
	aoc.SolvePuzzle(day, aoc.Part2, solvePart2, 78826761)
}
