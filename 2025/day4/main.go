package main

import (
	"fmt"
	"main/utils"
	"os"
	"time"
)

func main() {
	err := os.Chdir("./2025/day4")
	utils.CheckErr(err)

	lines, err := utils.ReadLines("./demo_input.txt")
	utils.CheckErr(err)

	grid := [][]byte{}
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}

	start := time.Now()
	sum := part1(grid)
	elapsed := time.Since(start)

	fmt.Println("Part 1 result: ", sum)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")

	start2 := time.Now()
	sum2 := part2(grid)
	elapsed2 := time.Since(start2)

	fmt.Println("Part 2 result: ", sum2)
	fmt.Println("Part 2 took: ", elapsed2.Microseconds(), " μs")
}

func part1(grid [][]byte) int64 {
	rows := len(grid)
	cols := len(grid[0])

	target := byte('@')

	acceptedPaperroll := 0

	for r := range rows {
		for c := range cols {

			// Only check if it is paperroll
			if grid[r][c] != target {
				continue
			}

			if countAdjacent(grid, r, c, rows, cols, target) < 4 {
				acceptedPaperroll++
			}

		}
	}

	return int64(acceptedPaperroll)
}

func part2(originalGrid [][]byte) int {
	rows := len(originalGrid)
	cols := len(originalGrid[0])

	target := byte('@')

	result := 0

	for {
		acceptedPaperroll := 0

		for r := range rows {
			for c := range cols {

				if originalGrid[r][c] != target {
					continue
				}

				if countAdjacent(originalGrid, r, c, rows, cols, target) < 4 {
					acceptedPaperroll++
					originalGrid[r][c] = byte('x')
				}
			}
		}
		result += acceptedPaperroll

		if acceptedPaperroll == 0 {
			break
		}
	}

	return result
}

func countAdjacent(grid [][]byte, r int, c int, rows int, cols int, target byte) int {
	// adjacent coordinates
	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0

	for _, dir := range dirs {
		newR := r + dir[0]
		newC := c + dir[1]

		// check out of bounds
		if newR < 0 || newR >= rows || newC < 0 || newC >= cols {
			continue
		}

		if grid[newR][newC] == target {
			count++
		}

		if count == 4 {
			return count
		}
	}

	return count
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
