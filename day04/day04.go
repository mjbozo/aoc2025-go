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
			deltas := []utils.Pair[int, int]{
				{First: -1, Second: -1}, {First: -1, Second: 0}, {First: -1, Second: 1}, {First: 0, Second: -1},
				{First: 0, Second: 1}, {First: 1, Second: -1}, {First: 1, Second: 0}, {First: 1, Second: 1},
			}

			for _, delta := range deltas {
				dx, dy := delta.First, delta.Second
				if dy == 0 && dx == 0 {
					continue
				}

				currentX, currentY := x+dx, y+dy
				if currentX >= 0 && currentX < width && currentY >= 0 && currentY < height && input[currentY][currentX] == '@' {
					adjacent++
				}
			}

			if adjacent < 4 {
				accessible++
			}
		}
	}

	return accessible
}

type Cell struct {
	accessibiltyScore  int
	neighbourPositions []utils.Pair[int, int]
}

func (c *Cell) String() string {
	return fmt.Sprintf("[%d]", c.accessibiltyScore)
}

func part2(input []string) int {
	width, height := len(input[0]), len(input)
	purgeList := make(chan *Cell, width*height)

	cells := utils.BuildGrid(input, func(i int, line string) []*Cell {
		row := make([]*Cell, 0)
		for j, c := range line {
			if c == '.' {
				row = append(row, nil)
			} else {
				cell := &Cell{neighbourPositions: make([]utils.Pair[int, int], 0)}
				deltas := []utils.Pair[int, int]{
					{First: -1, Second: -1}, {First: -1, Second: 0}, {First: -1, Second: 1}, {First: 0, Second: -1},
					{First: 0, Second: 1}, {First: 1, Second: -1}, {First: 1, Second: 0}, {First: 1, Second: 1},
				}

				for _, delta := range deltas {
					dx, dy := delta.First, delta.Second
					if dy == 0 && dx == 0 {
						continue
					}

					currentX, currentY := j+dx, i+dy
					if currentX >= 0 && currentX < width && currentY >= 0 && currentY < height && input[currentY][currentX] == '@' {
						cell.accessibiltyScore++
						cell.neighbourPositions = append(cell.neighbourPositions, utils.Pair[int, int]{First: currentX, Second: currentY})
					}
				}

				row = append(row, cell)
				if cell.accessibiltyScore < 4 {
					purgeList <- cell
				}
			}
		}

		return row
	})

	purged := 0
	for len(purgeList) > 0 {
		purged++
		x := <-purgeList

		for _, pos := range x.neighbourPositions {
			(*cells)[pos.Second][pos.First].accessibiltyScore--
			if (*cells)[pos.Second][pos.First].accessibiltyScore == 3 {
				// just dropped below 4 - add to purge list
				purgeList <- (*cells)[pos.Second][pos.First]
			}
		}
	}

	return purged
}
