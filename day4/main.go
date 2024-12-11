package day4

import (
	"fmt"

	"github.com/underscoren/aoc2024/util"
)

// tests whether a given string exists in a given direction in a slice of strings
func testDirection(lines []string, pos []int, direction []int, text string) bool {
	x, y := pos[0], pos[1]
	dx, dy := direction[0], direction[1]

	for _, char := range text {
		if y < 0 || x < 0 || y >= len(lines) || x >= len(lines[y]) || lines[y][x] != byte(char) {
			return false
		}

		x += dx
		y += dy
	}

	return true
}

func Main() {
	lines := util.ReadLines("day4/input")

	directions := [][]int{
		{0, 1},   // right
		{-1, 1},  // down-right
		{-1, 0},  // down
		{-1, -1}, // down-left
		{0, -1},  // left
		{1, -1},  // up-left
		{1, 0},   // up
		{1, 1},   // up-right
	}

	var xmasCount int

	// check each position
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			// check all directions
			for _, dir := range directions {
				if testDirection(lines, []int{x, y}, dir, "XMAS") {
					xmasCount++
				}
			}
		}
	}

	fmt.Printf("Day 4 XMAS count: %d\n", xmasCount)

	var x_masCount int

	diagonals := [][]int{
		{-1, 1},  // down-right
		{-1, -1}, // down-left
		{1, -1},  // up-left
		{1, 1},   // up-right
	}

	// X-MAS check
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			var found int // number of MAS found at this position

			// if we hit an "A", we're possibly in the middle of a MAS
			if line[x] == 'A' {
				// check all diagonals
				for _, dir := range diagonals {
					dx, dy := dir[0], dir[1]

					// offset by the inverse of the checking direction, then check for MAS
					if testDirection(lines, []int{x + (dx * (-1)), y + (dy * (-1))}, dir, "MAS") {
						found++
					}
				}
			}

			// if we found 2 MAS, then we have an X-MAS
			if found == 2 {
				x_masCount++
			}
		}
	}

	fmt.Printf("Day 4 X-MAS count: %d\n", x_masCount)
}
