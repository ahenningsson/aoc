package main

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
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
	err := os.Chdir("./2023/day4")
	check(err)

	lines, err := readLines("./input.txt")
	check(err)
	// // End boiler plate
	part1Res := 0
	for _, line := range lines {
		part1Res += part1(line)
	}
	println("Part 1: ", part1Res)

	println("Part 2: ", part2(lines))
}

func part1(line string) int {
	row := strings.Split(line, ":")
	gameRow := strings.Split(row[1], "|")
	winningNumbersStr := strings.Split(strings.TrimSpace(gameRow[0]), " ")
	winningNumbers := []int{}
	for _, numb := range winningNumbersStr {
		if len(numb) == 0 {
			continue
		}
		val, err := strconv.Atoi(numb)
		check(err)
		winningNumbers = append(winningNumbers, val)
	}

	ownNumbersStr := strings.Split(strings.TrimSpace(gameRow[1]), " ")
	ownNumbers := []int{}
	for _, numb := range ownNumbersStr {
		if len(numb) == 0 {
			continue
		}
		val, err := strconv.Atoi(numb)
		check(err)
		ownNumbers = append(ownNumbers, val)
	}

	winCount := 0
	for _, numb := range ownNumbers {
		if slices.Contains(winningNumbers, numb) {
			winCount += 1
		}
	}

	if winCount > 0 {
		return int(math.Pow(2, float64(winCount-1)))
	}
	return 0
}

func part2(lines []string) int {
	scratchCards := map[int]int{}

	for idx, line := range lines {
		currentGame := idx + 1
		nextGame := currentGame + 1

		row := strings.Split(line, ":")
		gameRow := strings.Split(row[1], "|")
		winningNumbersStr := strings.Split(strings.TrimSpace(gameRow[0]), " ")
		winningNumbers := []int{}
		for _, numb := range winningNumbersStr {
			if len(numb) == 0 {
				continue
			}
			val, err := strconv.Atoi(numb)
			check(err)
			winningNumbers = append(winningNumbers, val)
		}

		ownNumbersStr := strings.Split(strings.TrimSpace(gameRow[1]), " ")
		ownNumbers := []int{}
		for _, numb := range ownNumbersStr {
			if len(numb) == 0 {
				continue
			}
			val, err := strconv.Atoi(numb)
			check(err)
			ownNumbers = append(ownNumbers, val)
		}

		matchingNumbers := 0
		for _, numb := range ownNumbers {
			if slices.Contains(winningNumbers, numb) {
				matchingNumbers++
			}
		}

		scratchCards[currentGame]++ // Add the current game to the map
		for i := nextGame; i < nextGame+matchingNumbers; i++ {
			scratchCards[i] += scratchCards[currentGame]
		}
	}

	totalVal := 0
	for k := range scratchCards {
		totalVal += scratchCards[k]
	}

	return totalVal
}
