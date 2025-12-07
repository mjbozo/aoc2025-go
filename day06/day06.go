package day06

import (
	"aoc2025/utils"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"
)

func Run() {
	input, err := utils.ReadInputLines("day06/input.txt", 6)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	start := time.Now()
	part1Result := part1(input)
	elapsed := time.Since(start)
	fmt.Printf("Part 1: %d (%v)\n", part1Result, elapsed)

	start = time.Now()
	part2Result := part2(input)
	elapsed = time.Since(start)
	fmt.Printf("Part 2: %d (%v)\n", part2Result, elapsed)
}

func part1(input []string) int {
	lines := make([][]string, 0)
	for _, line := range input {
		normalised := strings.Fields(line)
		lines = append(lines, normalised)
	}

	total := 0
	rows := len(lines)
	for x := 0; x < len(lines[0]); x++ {
		operation := lines[rows-1][x]
		current := 0
		if operation == "*" {
			current = 1
		}

		for y := len(lines) - 2; y >= 0; y-- {
			val, _ := strconv.Atoi(lines[y][x])
			switch operation {
			case "+":
				current += val
			case "*":
				current *= val
			}
		}

		total += current
	}

	return total
}

func part2(input []string) int {
	operations := strings.Fields(input[len(input)-1])
	slices.Reverse(operations)

	// input from http response cuts leading spaces
	// need to reformat input first
	longest := 0
	for _, line := range input {
		longest = max(longest, len(line))
	}

	for i, line := range input {
		diff := longest - len(line)
		input[i] = strings.Repeat(" ", diff) + line
	}

	grid := make([][]int, 0)
	values := make([]int, 0)
	for x := len(input[0]) - 1; x >= 0; x-- {
		num := 0
		for y := 0; y < len(input)-1; y++ {
			if input[y][x] != ' ' {
				num = num*10 + int(input[y][x]-'0')
			}
		}

		if num == 0 {
			grid = append(grid, values)
			values = make([]int, 0)
		} else {
			values = append(values, num)
		}
	}

	grid = append(grid, values)
	total := 0
	for i, row := range grid {
		current := row[0]
		for _, num := range row[1:] {
			switch operations[i] {
			case "+":
				current += num
			case "*":
				current *= num
			}
		}
		total += current
	}

	return total
}
