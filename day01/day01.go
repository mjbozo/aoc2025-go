package day01

import (
	"aoc2025/utils"
	"fmt"
	"log"
	"strconv"
	"time"
)

func Run() {
	input, err := utils.ReadInputLines("day01/input.txt", 1)
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
	current := 50
	zeros := 0

	for _, rotation := range input {
		amount, _ := strconv.Atoi(rotation[1:])
		remaining := amount % 100

		if rotation[0] == 'L' {
			current = (current + (100 - remaining)) % 100
		} else {
			current = (current + remaining) % 100
		}

		if current == 0 {
			zeros++
		}
	}

	return zeros
}

func part2(input []string) int {
	current := 50
	zeros := 0

	for _, rotation := range input {
		amount, _ := strconv.Atoi(rotation[1:])
		remaining := amount % 100
		zeros += amount / 100

		if rotation[0] == 'L' {
			if (current != 0 && current-remaining < 0) || current-remaining == 0 {
				zeros++
			}
			current = (current + (100 - remaining)) % 100
		} else {
			if current+remaining > 99 {
				zeros++
			}
			current = (current + remaining) % 100
		}
	}

	return zeros
}

