package day09

import (
	"aoc2025/utils"
	"fmt"
	"log"
	"math"
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
			xDiff := utils.Abs(points[j].First-points[i].First) + 1
			yDiff := utils.Abs(points[j].Second-points[i].Second) + 1
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
	var bestPointA pos
	var bestPointB pos

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
			if xDiff*yDiff > largest {
				largest = xDiff * yDiff
				bestPointA = points[i]
				bestPointB = points[j]
			}
		}
	}

	visualise(points, bestPointA, bestPointB)
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

func visualise(points []pos, a, b pos) {
	minX := math.MaxInt
	minY := math.MaxInt
	maxX := 0
	maxY := 0

	pointsLookup := utils.HashSet[pos]{}
	edgeLookup := utils.HashSet[pos]{}

	for _, point := range points {
		minX = min(minX, point.First)
		minY = min(minY, point.Second)
		maxX = max(maxX, point.First)
		maxY = max(maxY, point.Second)
	}

	xRange := float64(maxX - minX)
	yRange := float64(maxY - minY)
	xScale := 200.0 / xRange
	yScale := 120.0 / yRange

	for i, point := range points {
		newX := int(float64(point.First-minX) * xScale)
		newY := int(float64(point.Second-minY) * yScale)
		points[i] = pos{newX, newY}
		pointsLookup.Insert(pos{newX, newY})

		if i > 0 {
			if points[i].First == points[i-1].First {
				for y := min(points[i].Second, points[i-1].Second); y <= max(points[i].Second, points[i-1].Second); y++ {
					edgeLookup.Insert(pos{points[i].First, y})
				}
			}

			if points[i].Second == points[i-1].Second {
				for x := min(points[i].First, points[i-1].First); x <= max(points[i].First, points[i-1].First); x++ {
					edgeLookup.Insert(pos{x, points[i].Second})
				}
			}
		}
	}

	maxX = 0
	maxY = 0

	for point := range pointsLookup {
		minX = min(minX, point.First)
		minY = min(minY, point.Second)
		maxX = max(maxX, point.First)
		maxY = max(maxY, point.Second)
	}

	a.First = int(float64(a.First-minX) * xScale)
	a.Second = int(float64(a.Second-minY) * yScale)
	b.First = int(float64(b.First-minX) * xScale)
	b.Second = int(float64(b.Second-minY) * yScale)

	boxLookup := utils.HashSet[pos]{}
	for y := min(a.Second, b.Second); y <= max(a.Second, b.Second); y++ {
		for x := min(a.First, b.First); x <= max(a.First, b.First); x++ {
			boxLookup.Insert(pos{x - 4, y})
		}
	}

	var output string
	for y := minY - 1; y <= maxY+1; y++ {
		row := ""
		for x := minX - 1; x <= maxX+1; x++ {
			if edgeLookup.Contains(pos{x, y}) && !pointsLookup.Contains(pos{x, y}) {
				row += utils.Green("X")
			} else if pointsLookup.Contains(pos{x, y}) {
				row += utils.Red("#")
			} else if boxLookup.Contains(pos{x, y}) {
				row += utils.Blue("O")
			} else {
				row += "."
			}
		}
		output += row + "\n"
	}

	fmt.Println(output)
}
