package main

import (
	"fmt"
	"main/utils"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	err := os.Chdir("../2024/day4")
	utils.CheckErr(err)

	lines, err := utils.ReadLines("./input.txt")
	utils.CheckErr(err)

	var inputSlice [][]string
	for _, line := range lines {
		val := strings.Split(line, "")
		inputSlice = append(inputSlice, val)
	}

	start := time.Now()
	val := part1(inputSlice)
	elapsed := time.Since(start)

	start2 := time.Now()
	val2 := part2(inputSlice)
	elapsed2 := time.Since(start2)

	fmt.Println("Part 1 result: ", val)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")

	fmt.Println("Part 2 result: ", val2)
	fmt.Println("Part 2 took: ", elapsed2.Microseconds(), " μs")
}

func transpose(slice [][]string) [][]string {
	rows := len(slice)
	cols := len(slice[0])

	var transposed = make([][]string, cols)
	for c := range transposed {
		transposed[c] = make([]string, rows)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			transposed[c][r] = slice[r][c]
		}
	}

	return transposed
}

func checkTopLeftToBottomRightDiagonal(slice [][]string, reg *regexp.Regexp) int {
	rows := len(slice)
	cols := len(slice[0])

	occurrences := 0

	// Check diagonals starting from each column on the first row
	for c := 0; c < cols; c++ {
		diagonal := ""
		row := 0
		currentCol := c

		for row < rows && currentCol < cols {
			diagonal += slice[row][currentCol]
			if len(diagonal) == 4 {
				if findMatch(diagonal, reg) {
					occurrences++
				}
				diagonal = diagonal[1:] // Never let the string become more than len 4. This is because in the situation where the string is XMASXAMS and it regex the entire string, it will return 1, not 2.
			}
			row++
			currentCol++
		}
	}

	for r := 1; r < rows; r++ {
		diagonal := ""
		col := 0
		currentRow := r

		for currentRow < rows && col < cols {
			diagonal += slice[currentRow][col]
			if len(diagonal) == 4 {
				if findMatch(diagonal, reg) {
					occurrences++
				}
				diagonal = diagonal[1:]
			}
			currentRow++
			col++
		}
	}
	return occurrences
}

func checkTopRightToBottomLeftDiagonal(slice [][]string, reg *regexp.Regexp) int {
	rows := len(slice)
	cols := len(slice[0])

	occurrences := 0

	// Check diagonals starting from each column on the first row
	for c := cols - 1; c >= 0; c-- {
		diagonal := ""
		row := 0
		currentCol := c

		for row < rows && currentCol >= 0 {
			diagonal += slice[row][currentCol]
			if len(diagonal) == 4 {
				if findMatch(diagonal, reg) {
					occurrences++
				}
				diagonal = diagonal[1:]
			}
			row++
			currentCol--
		}
	}

	// Check diagonals starting from each row of the last column
	for r := 1; r < rows; r++ {
		diagonal := ""
		col := cols - 1
		currentRow := r

		for currentRow < rows && col >= 0 {
			diagonal += slice[currentRow][col]
			if len(diagonal) == 4 {
				if findMatch(diagonal, reg) {
					occurrences++
				}
				diagonal = diagonal[1:]
			}
			currentRow++
			col--
		}
	}
	return occurrences
}

func findMatch(str string, reg *regexp.Regexp) bool {
	matches := reg.MatchString(str)

	return matches
}

func part1(inputSlice [][]string) int {
	reg, _ := regexp.Compile(`XMAS|SAMX`)

	rows := len(inputSlice)
	cols := len(inputSlice[0])

	occurrences := 0

	transposed := transpose(inputSlice)
	for r := 0; r < rows; r++ {
		substring := ""
		substring2 := ""
		for c := 0; c < cols; c++ {
			substring += inputSlice[r][c]
			substring2 += transposed[r][c]

			if len(substring) == 4 {
				if reg.MatchString(substring) {
					occurrences++
				}
				if reg.MatchString(substring2) {
					occurrences++
				}
				substring = substring[1:]
				substring2 = substring2[1:]
			}
		}
	}

	// Check diagonals
	occurrences += checkTopLeftToBottomRightDiagonal(inputSlice, reg)
	occurrences += checkTopRightToBottomLeftDiagonal(inputSlice, reg)

	return occurrences
}

func part2(inputSlice [][]string) int {
	reg, _ := regexp.Compile(`MAS|SAM`)

	rows := len(inputSlice)
	cols := len(inputSlice[0])

	occurrences := 0

	for r := 1; r < (rows - 1); r++ {
		for c := 1; c < (cols - 1); c++ {
			centerCharacter := inputSlice[r][c]

			if centerCharacter != "A" {
				continue
			}

			// Get diagonals
			tl_br_diag := inputSlice[r-1][c-1] + centerCharacter + inputSlice[r+1][c+1]
			tr_bl_diag := inputSlice[r-1][c+1] + centerCharacter + inputSlice[r+1][c-1]

			if reg.MatchString(tl_br_diag) && reg.MatchString(tr_bl_diag) {
				occurrences++
			}
		}
	}

	return occurrences
}
