package day09

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
	input, err := utils.ReadInputLines("day09/input.txt", 9)
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

type pos utils.Pair[int, int]
type edge utils.Triple[int, int, int] // First constant, Second and Third is range of line

func part1(input []string) int {
	points := make([]pos, 0)
	for _, line := range input {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, pos{x, y})
	}
	slices.SortFunc(points, func(a, b pos) int {
		if a.First == b.First {
			return a.Second - b.Second
		}
		return a.First - b.First
	})

	largest := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			xDiff := points[j].First - points[i].First + 1
			yDiff := points[j].Second - points[i].Second + 1
			area := xDiff * yDiff
			largest = max(largest, area)
		}
	}

	return largest
}

func part2(input []string) int {
	horizontalEdges := make([]edge, 0)
	verticalEdges := make([]edge, 0)

	points := make([]pos, 0)
	for _, line := range input {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, pos{x, y})
	}

	for i := 1; i < len(points); i++ {
		current := points[i]
		prev := points[i-1]
		if current.First == prev.First {
			verticalEdges = append(verticalEdges, edge{current.First, min(current.Second, prev.Second), max(current.Second, prev.Second)})
		} else {
			horizontalEdges = append(horizontalEdges, edge{current.Second, min(current.First, prev.First), max(current.First, prev.First)})
		}
	}

	// wrap around
	current := points[0]
	prev := points[len(points)-1]
	if current.First == prev.First {
		verticalEdges = append(verticalEdges, edge{current.First, min(current.Second, prev.Second), max(current.Second, prev.Second)})
	} else {
		horizontalEdges = append(horizontalEdges, edge{current.Second, min(current.First, prev.First), max(current.First, prev.First)})
	}

	largest := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			// check if box crosses any horizontal or vertical line
			minX := min(points[i].First, points[j].First)
			minY := min(points[i].Second, points[j].Second)
			maxX := max(points[i].First, points[j].First)
			maxY := max(points[i].Second, points[j].Second)

			if intersectsEdge(horizontalEdges, minY, minX, maxY, maxX) ||
				intersectsEdge(verticalEdges, minX, minY, maxX, maxY) {
				continue
			}

			xDiff := utils.Abs(points[j].First-points[i].First) + 1
			yDiff := utils.Abs(points[j].Second-points[i].Second) + 1
			largest = max(largest, xDiff*yDiff)
		}
	}

	return largest
}

func intersectsEdge(edges []edge, a, b, c, d int) bool {
	for _, line := range edges {
		contant, variableMin, variableMax := line.First, line.Second, line.Third
		if contant > a && contant < c && variableMax > b && variableMin < d {
			return true
		}
	}
	return false
}
