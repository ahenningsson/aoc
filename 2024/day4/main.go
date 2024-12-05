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

	start := time.Now()
	var inputSlice [][]string
	for _, line := range lines {
		val := strings.Split(line, "")
		inputSlice = append(inputSlice, val)
	}
	val := part1(inputSlice)
	elapsed := time.Since(start)

	fmt.Println("Part 1 result: ", val)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")
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
			row++
			currentCol++
		}

		if len(diagonal) > 3 {
			occurrences += findMatch(diagonal, reg)
		}
	}

	for r := 1; r < rows; r++ {
		diagonal := ""
		col := 0
		currentRow := r

		for currentRow < rows && col < cols {
			diagonal += slice[currentRow][col]
			currentRow++
			col++
		}
		if len(diagonal) > 3 {
			occurrences += findMatch(diagonal, reg)
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
			row++
			currentCol--
		}
		if len(diagonal) > 3 {
			occurrences += findMatch(diagonal, reg)
		}
	}

	// Check diagonals starting from each row of the last column
	for r := 1; r < rows; r++ {
		diagonal := ""
		col := cols - 1
		currentRow := r

		for currentRow < rows && col >= 0 {
			diagonal += slice[currentRow][col]
			currentRow++
			col--
		}

		if len(diagonal) > 3 {
			occurrences += findMatch(diagonal, reg)
		}
	}
	return occurrences
}

func findMatch(str string, reg *regexp.Regexp) int {
	// fmt.Println(str)
	matches := reg.FindAllStringSubmatch(str, -1)
	return len(matches)
}

func part1(inputSlice [][]string) int {
	reg, _ := regexp.Compile(`XMAS|SAMX`)

	occurrences := 0
	for _, line := range inputSlice {
		matches := reg.FindAllStringSubmatch(strings.Join(line, ""), -1)
		occurrences += len(matches)
	}

	transposed := transpose(inputSlice)
	for _, line := range transposed {
		matches := reg.FindAllStringSubmatch(strings.Join(line, ""), -1)
		occurrences += len(matches)
	}

	occurrences += checkTopLeftToBottomRightDiagonal(inputSlice, reg)
	occurrences += checkTopRightToBottomLeftDiagonal(inputSlice, reg)

	return occurrences
}
