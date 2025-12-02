package day02

import (
	"aoc2025/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func Run() {
	input, err := utils.ReadInput("day02/input.txt", 2)
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

func part1(input string) int {
	sum := 0

	ranges := strings.Split(input, ",")
	c := make(chan int, len(ranges))
	q := make(chan int)
	for _, r := range ranges {
		go findInvalidIds(r, c, q)
	}

	finished := 0
	for {
		select {
		case x := <-c:
			sum += x
		case <-q:
			finished++
			if finished == len(ranges) {
				return sum
			}
		}
	}
}

func findInvalidIds(r string, c chan int, q chan int) {
	parts := strings.Split(r, "-")
	lower, _ := strconv.Atoi(parts[0])
	upper, _ := strconv.Atoi(parts[1])

	for i := lower; i <= upper; i++ {
		str := fmt.Sprintf("%d", i)
		l := len(str)
		if l%2 == 0 && str[0:l/2] == str[l/2:] {
			c <- i
		}
	}

	q <- 0
}

func part2(input string) int {
	sum := 0

	ranges := strings.Split(input, ",")
	c := make(chan int, len(ranges))
	q := make(chan int)
	for _, r := range ranges {
		go findInvalidIdsAgain(r, c, q)
	}

	finished := 0
	for {
		select {
		case x := <-c:
			sum += x
		case <-q:
			finished++
			if finished == len(ranges) {
				return sum
			}
		}
	}
}

func findInvalidIdsAgain(r string, c chan int, q chan int) {
	parts := strings.Split(r, "-")
	lower, _ := strconv.Atoi(parts[0])
	upper, _ := strconv.Atoi(parts[1])

	for i := lower; i <= upper; i++ {
		str := fmt.Sprintf("%d", i)
		l := len(str)

		for x := 1; x <= l/2; x++ {
			pattern := str[:x]
			if l%len(pattern) != 0 {
				continue
			}

			matches := true
			for z := 0; z < l; z += x {
				if str[z:z+x] != pattern {
					matches = false
					break
				}
			}

			if matches {
				c <- i
				break
			}
		}
	}

	q <- 0
}
