package main

import (
	"fmt"
	"main/utils"
	"os"
	"strings"
	"time"
)

func main() {
	err := os.Chdir("../2024/day6")
	utils.CheckErr(err)

	lines, err := utils.ReadLines("./input.txt")
	utils.CheckErr(err)

	var puzzleMap [][]string
	for _, line := range lines {
		val := strings.Split(line, "")
		puzzleMap = append(puzzleMap, val)
	}

	start := time.Now()
	val := part1(puzzleMap)
	elapsed := time.Since(start)

	start2 := time.Now()
	val2 := part2(puzzleMap)
	elapsed2 := time.Since(start2)

	fmt.Println("Part 1 result: ", val)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")

	fmt.Println("Part 2 result: ", val2)
	fmt.Println("Part 2 took: ", elapsed2.Microseconds(), " μs")
}

func turn(direction string) string {
	switch direction {
	case "^":
		return ">"
	case ">":
		return "v"
	case "v":
		return "<"
	case "<":
		return "^"
	}

	return ""
}

func part1(puzzleMap [][]string) int {
	rows := len(puzzleMap)
	cols := len(puzzleMap[0])

	visited := make([][]int, rows)
	for i := range visited {
		visited[i] = make([]int, cols)
	}

	r, c := 0, 0
	direction := ""
	// Find start position and direction
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if puzzleMap[i][j] == "v" || puzzleMap[i][j] == "^" || puzzleMap[i][j] == "<" || puzzleMap[i][j] == ">" {
				direction = puzzleMap[i][j]
				r, c = i, j
				break
			}
		}
	}

	isOnMap := true
	for isOnMap {
		// Mark square as visited
		visited[r][c] = 1

		// Find next step
		var nextR, nextC int

		switch direction {
		case "^":
			nextR, nextC = r-1, c
		case ">":
			nextR, nextC = r, c+1
		case "v":
			nextR, nextC = r+1, c
		case "<":
			nextR, nextC = r, c-1
		}

		// check so that it is not out of bounds
		if nextR < 0 || nextR >= rows || nextC < 0 || nextC >= cols {
			isOnMap = false
			break
		}

		// Check if obstacle
		if puzzleMap[nextR][nextC] == "#" {
			direction = turn(direction)
		} else {
			r, c = nextR, nextC
		}
	}
	sum := 0
	for _, row := range visited {
		for _, col := range row {
			sum += col
		}
	}
	return sum
}

type Visited struct {
	Row int
	Col int
}

func getGuardPositions(puzzleMap [][]string) []Visited {
	rows := len(puzzleMap)
	cols := len(puzzleMap[0])

	var visited []Visited
	visitedSet := make(map[Visited]bool) // Use a map to track unique positions

	r, c := 0, 0
	direction := ""
	// Find start position and direction
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if puzzleMap[i][j] == "v" || puzzleMap[i][j] == "^" || puzzleMap[i][j] == "<" || puzzleMap[i][j] == ">" {
				direction = puzzleMap[i][j]
				r, c = i, j
				break
			}
		}
	}

	isOnMap := true
	for isOnMap {
		currState := Visited{Row: r, Col: c}

		// Add position as visited
		if !visitedSet[currState] {
			visitedSet[currState] = true
			visited = append(visited, currState)
		}

		// Find next step
		var nextR, nextC int

		switch direction {
		case "^":
			nextR, nextC = r-1, c
		case ">":
			nextR, nextC = r, c+1
		case "v":
			nextR, nextC = r+1, c
		case "<":
			nextR, nextC = r, c-1
		}

		// check so that it is not out of bounds
		if nextR < 0 || nextR >= rows || nextC < 0 || nextC >= cols {
			isOnMap = false
			break
		}

		// Check if obstacle
		if puzzleMap[nextR][nextC] == "#" {
			direction = turn(direction)
		} else {
			r, c = nextR, nextC
		}
	}

	return visited
}

type VisitedState struct {
	Row       int    // row position
	Col       int    // column position
	Direction string // direction ('^', '>', 'v', '<')
}

// This was the second attempt, from 6 seconds to 1.1 seconds in execution time.
func part2(puzzleMap [][]string) int {
	rows := len(puzzleMap)
	cols := len(puzzleMap[0])

	// Result
	obstaclePositions := 0

	// Guard start
	originalR, originalC := 0, 0
	originalDirection := ""
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if puzzleMap[i][j] == "v" || puzzleMap[i][j] == "^" || puzzleMap[i][j] == "<" || puzzleMap[i][j] == ">" {
				originalDirection = puzzleMap[i][j]
				originalR, originalC = i, j
				break
			}
		}
	}

	// Guard positions
	guardPositions := getGuardPositions(puzzleMap)

	for _, gPos := range guardPositions {
		if gPos.Row == originalR && gPos.Col == originalC {
			continue
		}
		r, c := gPos.Row, gPos.Col

		originalValue := puzzleMap[r][c]

		// Temporarily new obstacle and run the simulation
		if originalValue == "." {
			puzzleMap[r][c] = "0"
		}

		if simulation(puzzleMap, originalR, originalC, originalDirection, rows, cols) {
			obstaclePositions++
		}

		puzzleMap[r][c] = originalValue
	}

	return obstaclePositions
}

func simulation(puzzleMap [][]string, startR, startC int, startDir string, rowsLen, colsLen int) bool {
	visited := make(map[VisitedState]bool)

	r, c := startR, startC
	direction := startDir

	isRunning := true
	for isRunning {
		// Mark square as visited
		// Idea is that if the guard revists the same square as previously + with the same current direction, a loop exists
		currState := VisitedState{Row: r, Col: c, Direction: direction}

		if visited[currState] {
			return true
		}
		visited[currState] = true

		// Find next step
		var nextR, nextC int

		switch direction {
		case "^":
			nextR, nextC = r-1, c
		case ">":
			nextR, nextC = r, c+1
		case "v":
			nextR, nextC = r+1, c
		case "<":
			nextR, nextC = r, c-1
		}

		// check so that it is not out of bounds
		if nextR < 0 || nextR >= rowsLen || nextC < 0 || nextC >= colsLen {
			return false
		}

		// Check if obstacle
		if puzzleMap[nextR][nextC] == "#" || puzzleMap[nextR][nextC] == "0" {
			direction = turn(direction)
		} else {
			r, c = nextR, nextC
		}
	}
	return false
}

// This was the first solution. Brute force and slow.
func part2_works(puzzleMap [][]string) int {
	rows := len(puzzleMap)
	cols := len(puzzleMap[0])

	// Result
	obstaclePositions := 0

	// Guard start
	originalR, originalC := 0, 0
	originalDirection := ""
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if puzzleMap[i][j] == "v" || puzzleMap[i][j] == "^" || puzzleMap[i][j] == "<" || puzzleMap[i][j] == ">" {
				originalDirection = puzzleMap[i][j]
				originalR, originalC = i, j
				break
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			originalValue := puzzleMap[i][j]
			// Add new obstacle on each position and run the simulation
			if originalValue == "." {
				puzzleMap[i][j] = "0"
			}

			if simulation(puzzleMap, originalR, originalC, originalDirection, rows, cols) {
				obstaclePositions++
			}

			puzzleMap[i][j] = originalValue
		}
	}

	return obstaclePositions
}
