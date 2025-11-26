package utils

import (
	"fmt"
	"strings"
)

// Grid data type, implemented as slice of slices
type Grid[T any] [][]T

// Constructor for grid struct
func BuildGrid[T, U any](inputLines []U, mapper func(v U) []T) *Grid[T] {
	grid := Grid[T]{}

	for _, row := range inputLines {
		mappedElements := mapper(row)
		grid = append(grid, mappedElements)
	}

	fmt.Println(grid)
	return &grid
}

// Default formatting for grid
//
// Pads columns to maintain element alignment
func (g Grid[T]) String() string {
	if len(g) == 0 {
		return "[[]]"
	}

	columnMaxLength := make([]int, len(g[0]))
	for _, row := range g {
		for i, val := range row {
			length := len(fmt.Sprintf("%v", val))
			if length > columnMaxLength[i] {
				columnMaxLength[i] = length
			}
		}
	}

	var output string
	for _, row := range g {
		var line string = "[ "
		for i, val := range row {
			padLength := columnMaxLength[i] - len(fmt.Sprintf("%v", val))
			line += fmt.Sprintf("%v%s ", val, strings.Repeat(" ", padLength))
		}
		output += line + "]" + "\n"
	}
	return output
}

// Gets neighbours in 4 cardinal directions around point, if they exist
//
// Returned in order: UP, RIGHT, DOWN, LEFT
func (g Grid[T]) Neighbours(x, y int) []T {
	neighbours := make([]T, 0)

	if y > 0 {
		neighbours = append(neighbours, g[y-1][x])
	}

	if x < len(g[0])-1 {
		neighbours = append(neighbours, g[y][x+1])
	}

	if y < len(g)-1 {
		neighbours = append(neighbours, g[y+1][x])
	}

	if x > 0 {
		neighbours = append(neighbours, g[y][x-1])
	}

	return neighbours
}
