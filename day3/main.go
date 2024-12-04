package day3

import (
	"fmt"
	"regexp"

	"github.com/underscoren/aoc2024/util"
)

func Main() {
	text := util.ReadFile("day3/input")

	// regex to match all mul() instructions. Captures just the digits.
	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	textMatches := mulRegex.FindAllStringSubmatch(text, -1)

	var totalSum int

	for _, regexMatch := range textMatches {
		numbers := util.AtoiFatalSlice(regexMatch[1:])

		totalSum += numbers[0] * numbers[1]
	}

	fmt.Printf("Day 3 sum of mul(): %d\n", totalSum)

	// regex to match all do, don't and mul instructions. Captures the do/don't/mul and digits for mul
	stateRegex := regexp.MustCompile(`(do(?:n't)?)\(\)|(mul)\((\d+),(\d+)\)`)

	var totalStateAware int

	enabled := true
	for _, matches := range stateRegex.FindAllStringSubmatch(text, -1) {
		instruction := matches[1]
		if instruction == "" {
			instruction = matches[2]
		}

		switch instruction { // 1st capture group will always be the instruction
		case "do":
			enabled = true
		case "don't":
			enabled = false
		case "mul":
			if !enabled {
				continue
			}

			numbers := util.AtoiFatalSlice(matches[3:])
			totalStateAware += numbers[0] * numbers[1]
		}
	}

	fmt.Printf("Day 3 do/don't sum of mul(): %d\n", totalStateAware)
}
