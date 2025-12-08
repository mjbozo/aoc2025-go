package day08

import (
	"aoc2025/utils"
	"log"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := utils.ReadInputLines("example.txt", 8)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := 40
	result := part1(input, 10)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

func TestPart2(t *testing.T) {
	input, err := utils.ReadInputLines("example.txt", 8)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := 25272
	result := part2(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

