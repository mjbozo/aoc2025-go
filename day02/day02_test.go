package day02

import (
	"aoc2025/utils"
	"log"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := utils.ReadInput("example.txt", 2)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := 1227775554
	result := part1(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

func TestPart2(t *testing.T) {
	input, err := utils.ReadInput("example.txt", 2)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	expected := 4174379265
	result := part2(input)
	if result != expected {
		t.Fatalf(utils.Red("Expected %d, got %d\n"), expected, result)
	}
}

