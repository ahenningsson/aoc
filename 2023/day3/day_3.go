package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

// check if error is nil
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// read lines of the input file
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // append each line to the lines slice
	}
	return lines, scanner.Err()
}

func main() {

	// Boiler plate
	err := os.Chdir("./2023/day3")
	check(err)

	lines, err := readLines("./input.txt")
	check(err)
	// // End boiler plate

	maxRows := len(lines)
	maxCols := len(lines[0])

	part1(lines, maxRows, maxCols)
	part2(lines, maxRows, maxCols)
}

func part1(lines []string, maxRows int, maxCols int) {
	start := time.Now()
	partNumberVal := 0

	for row, line := range lines {
		currentColumn := 0 // column index
		start := 0         // keep track of where the number starts

		for currentColumn < maxCols {
			start = currentColumn
			num := ""

			for currentColumn < maxCols && isNumber(string(line[currentColumn])) {
				num += string(line[currentColumn])
				currentColumn += 1
			}

			if len(num) == 0 {
				currentColumn += 1
				continue
			}

			// Check if number is valid
			val, err := strconv.Atoi(num)
			check(err)

			if isSymbol(lines, maxRows, maxCols, row, start-1) || isSymbol(lines, maxRows, maxCols, row, currentColumn) {
				partNumberVal += val
			}

			// check if there's a symbol above or below the number (diagonally as well)
			for x := start - 1; x < currentColumn+1; x++ {
				if isSymbol(lines, maxRows, maxCols, row-1, x) || isSymbol(lines, maxRows, maxCols, row+1, x) {
					partNumberVal += val
				}
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Part 1 result: ", partNumberVal)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")
}

func part2(lines []string, maxRows int, maxCols int) {
	start := time.Now()
	sum := 0

	// Create a 3d matrix that holds the part numbers.
	// Idea is if the innermost slice has 2 elements, then it is a gear
	gearRatio := make([][][]int, maxRows)
	for i := range gearRatio {
		gearRatio[i] = make([][]int, maxCols)
	}

	for row, line := range lines {
		currentColumn := 0 // column index
		start := 0         // keep track of where the number starts

		for currentColumn < maxCols {
			start = currentColumn
			num := ""

			// Create the number, since it is continuous there's a while-loop here
			for currentColumn < maxCols && isNumber(string(line[currentColumn])) {
				num += string(line[currentColumn])
				currentColumn += 1
			}

			if len(num) == 0 {
				currentColumn += 1
				continue
			}

			// Check if number is valid
			val, err := strconv.Atoi(num)
			check(err)

			// check if there's a star to the left or right of the number
			if flag, status := isStar(lines, maxRows, maxCols, row, start-1); status {
				if flag {
					gearRatio[row][start-1] = append(gearRatio[row][start-1], val)
				}
			}
			if flag, status := isStar(lines, maxRows, maxCols, row, currentColumn); status {
				if flag {
					gearRatio[row][currentColumn] = append(gearRatio[row][currentColumn], val)
				}
			}

			// check if there's a star above or below the number (diagonally as well)
			for x := start - 1; x < currentColumn+1; x++ {

				if flag, status := isStar(lines, maxRows, maxCols, row-1, x); status {
					if flag {
						gearRatio[row-1][x] = append(gearRatio[row-1][x], val)
					}
				}
				if flag, status := isStar(lines, maxRows, maxCols, row+1, x); status {
					if flag {
						gearRatio[row+1][x] = append(gearRatio[row+1][x], val)
					}
				}
			}
		}
	}

	for i := 0; i < maxRows; i++ {
		for j := 0; j < maxCols; j++ {
			if string(lines[i][j]) == "*" && len(gearRatio[i][j]) == 2 {
				sum += gearRatio[i][j][0] * gearRatio[i][j][1]
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Part 2 result: ", sum)
	fmt.Println("Part 2 took: ", elapsed.Microseconds(), " μs")
}

// helper function to check if a character is a symbol
func isStar(lines []string, maxRows int, maxCols int, r int, c int) (bool, bool) {
	if r < 0 || r >= maxRows || c < 0 || c >= maxCols {
		return false, false
	}
	if string(lines[r][c]) == "*" {
		return true, true
	}
	return false, true
}

// helper function to check if a character is a number
func isNumber(char string) bool {
	if _, err := strconv.Atoi(char); err != nil {
		return false
	}
	return true
}

// helper function to check if a character is a symbol
func isSymbol(lines []string, maxRows int, maxCols int, r int, c int) bool {
	if r < 0 || r >= maxRows || c < 0 || c >= maxCols {
		return false
	}
	if _, err := strconv.Atoi(string(lines[r][c])); err == nil {
		return false
	}
	if string(lines[r][c]) == "." {
		return false
	}
	return true
}
