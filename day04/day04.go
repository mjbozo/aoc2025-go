package day04

import (
	"aoc2025/utils"
	"fmt"
	"log"
	"time"
)

func Run() {
	input, err := utils.ReadInputLines("day04/input.txt", 4)
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
	accessible := 0
	width := len(input[0])
	height := len(input)

	for y := range height {
		for x := range width {
			if input[y][x] != '@' {
				continue
			}

			adjacent := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					currentY := y + dy
					currentX := x + dx

					if dy == 0 && dx == 0 {
						continue
					}

					if currentX >= 0 && currentX < width && currentY >= 0 && currentY < height && input[currentY][currentX] == '@' {
						adjacent++
					}
				}
			}

			if adjacent < 4 {
				accessible++
			}
		}
	}

	return accessible
}

func part2(input []string) int {
	width := len(input[0])
	height := len(input)
	totalRemoved := 0

	for {
		accessible := make([]utils.Pair[int, int], 0)
		for y := range height {
			for x := range width {
				if input[y][x] != '@' {
					continue
				}

				adjacent := 0
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						currentY := y + dy
						currentX := x + dx

						if dy == 0 && dx == 0 {
							continue
						}

						if currentX >= 0 && currentX < width && currentY >= 0 && currentY < height && input[currentY][currentX] == '@' {
							adjacent++
						}
					}
				}

				if adjacent < 4 {
					accessible = append(accessible, utils.Pair[int, int]{First: x, Second: y})
				}
			}
		}

		if len(accessible) == 0 {
			break
		}

		totalRemoved += len(accessible)
		for _, pos := range accessible {
			prefix := input[pos.Second][:pos.First]
			suffix := input[pos.Second][pos.First+1:]
			input[pos.Second] = prefix + "." + suffix
		}
	}

	return totalRemoved
}

