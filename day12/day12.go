package day12

import (
	"aoc2025/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func Run() {
	input, err := utils.ReadInput("day12/input.txt", 12)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	start := time.Now()
	part1Result := part1(input)
	elapsed := time.Since(start)
	fmt.Printf("Part 1: %d (%v)\n", part1Result, elapsed)
}

type shape [][]int

type region struct {
	width          int
	height         int
	requiredShapes []int
}

func part1(input string) int {
	shapes, regions := parseInput(input)
	validRegions := 0

	for _, region := range regions {
		spotsRequired := 0
		for i, req := range region.requiredShapes {
			for range req {
				for _, row := range shapes[i] {
					for _, cell := range row {
						if cell == 1 {
							spotsRequired++
						}
					}
				}
			}
		}

		areaSize := region.height * region.width

		if areaSize >= spotsRequired {
			validRegions++
		}
	}

	return validRegions
}

func parseInput(input string) ([]shape, []region) {
	sections := strings.Split(input, "\n\n")
	shapes := make([]shape, 0)
	regions := make([]region, 0)

	for _, section := range sections {
		lines := strings.Split(section, "\n")
		if len(lines) > 1 && (strings.HasPrefix(lines[1], "#") || strings.HasPrefix(lines[1], ".")) {
			newShape := make(shape, 0)
			for _, line := range lines[1:] {
				row := make([]int, 0)
				for _, c := range line {
					if c == '#' {
						row = append(row, 1)
					} else {
						row = append(row, 0)
					}
				}
				newShape = append(newShape, row)
			}
			shapes = append(shapes, newShape)
		} else {
			regionLines := strings.Split(section, "\n")
			for _, line := range regionLines {
				parts := strings.Split(line, ": ")
				dimensions := strings.Split(parts[0], "x")
				width, _ := strconv.Atoi(dimensions[0])
				height, _ := strconv.Atoi(dimensions[1])

				required := make([]int, 0)
				for _, n := range strings.Split(parts[1], " ") {
					x, _ := strconv.Atoi(n)
					required = append(required, x)
				}

				regions = append(regions, region{width: width, height: height, requiredShapes: required})
			}
		}
	}

	return shapes, regions
}
