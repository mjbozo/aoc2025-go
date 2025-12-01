package main

import (
	"aoc2025/day01"
	// "aoc2025/day02"
	// "aoc2025/day03"
	// "aoc2025/day04"
	// "aoc2025/day05"
	// "aoc2025/day06"
	// "aoc2025/day07"
	// "aoc2025/day08"
	// "aoc2025/day09"
	// "aoc2025/day10"
	// "aoc2025/day11"
	// "aoc2025/day12"
	"aoc2025/daybreaker"
	"aoc2025/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatalln(utils.Red("[ ERR ] Not enough arguments. Provide the day you want to run"))
	}

	if args[1] == "-c" {
		if len(args) < 3 {
			log.Fatalln(utils.Red("[ ERR ] Not enough arguments. Provide the day you want to create"))
		}

		err := daybreaker.Create(args[2:])
		if err != nil {
			log.Fatalln(utils.Red(fmt.Sprintf("[ ERR ] Daybreaker failed: %s", err.Error())))
		}

		return
	}

	day := args[1]

	if day == "all" {
		runAll()
		return
	}

	dayNum := strings.TrimPrefix(day, "day")
	if len(dayNum) == 1 {
		// then i forgot to prepend day 1-9 with a zero
		day = fmt.Sprintf("day0%s", dayNum)
	}

	switch day {
	case "day01":
		day01.Run()
	// case "day02":
	// 	day02.Run()
	// case "day03":
	// 	day03.Run()
	// case "day04":
	// 	day04.Run()
	// case "day05":
	// 	day05.Run()
	// case "day06":
	// 	day06.Run()
	// case "day07":
	// 	day07.Run()
	// case "day08":
	// 	day08.Run()
	// case "day09":
	// 	day09.Run()
	// case "day10":
	// 	day10.Run()
	// case "day11":
	// 	day11.Run()
	// case "day12":
	// 	day12.Run()
	default:
		fmt.Printf("%s not completed yet\n", day)
	}
}

func runAll() {
	fmt.Println("DAY 01")
	day01.Run()
	// fmt.Println("\nDAY 02")
	// day02.Run()
	// fmt.Println("\nDAY 03")
	// day03.Run()
	// fmt.Println("\nDAY 04")
	// day04.Run()
	// fmt.Println("\nDAY 05")
	// day05.Run()
	// fmt.Println("\nDAY 06")
	// day06.Run()
	// fmt.Println("\nDAY 07")
	// day07.Run()
	// fmt.Println("\nDAY 08")
	// day08.Run()
	// fmt.Println("\nDAY 09")
	// day09.Run()
	// fmt.Println("\nDAY 10")
	// day10.Run()
	// fmt.Println("\nDAY 11")
	// day11.Run()
	// fmt.Println("\nDAY 12")
	// day12.Run()
}
