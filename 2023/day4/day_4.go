package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
	err := os.Chdir("./2023/day4")
	check(err)

	lines, err := readLines("./input.txt")
	check(err)
	// // End boiler plate
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	start := time.Now()
	part1Res := 0
	for _, line := range lines {
		row := strings.Split(line, ":")
		gameRow := strings.Split(row[1], "|")
		winningNumbersStr := strings.Split(strings.TrimSpace(gameRow[0]), " ")
		winningNumbers := make(map[int]bool)
		for _, numb := range winningNumbersStr {
			if len(numb) == 0 {
				continue
			}
			val, err := strconv.ParseInt(numb, 10, 32)
			check(err)
			winningNumbers[int(val)] = true // Add the winning numbers to a map
		}

		ownNumbersStr := strings.Split(strings.TrimSpace(gameRow[1]), " ")
		winCount := 0
		for _, numb := range ownNumbersStr {
			if len(numb) == 0 {
				continue
			}
			val, err := strconv.ParseInt(numb, 10, 32)
			check(err)
			if winningNumbers[int(val)] {
				winCount += 1
			}
		}

		if winCount > 0 {
			part1Res += int(math.Pow(2, float64(winCount-1)))
		}

	}

	elapsed := time.Since(start)
	fmt.Println("Part 1 result: ", part1Res)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")
}

func part2(lines []string) {
	start := time.Now()
	scratchCards := map[int]int{}

	for idx, line := range lines {
		currentGame := idx + 1
		nextGame := currentGame + 1

		row := strings.Split(line, ":")
		gameRow := strings.Split(row[1], "|")
		winningNumbersStr := strings.Split(strings.TrimSpace(gameRow[0]), " ")
		winningNumbers := make(map[int]bool)
		for _, numb := range winningNumbersStr {
			if len(numb) == 0 {
				continue
			}
			val, err := strconv.ParseInt(numb, 10, 32)
			check(err)
			winningNumbers[int(val)] = true // Add the winning numbers to a map
		}

		ownNumbersStr := strings.Split(strings.TrimSpace(gameRow[1]), " ")
		matchingNumbers := 0

		for _, numb := range ownNumbersStr {
			if len(numb) == 0 {
				continue
			}

			val, err := strconv.ParseInt(numb, 10, 32)
			check(err)
			// Check if the number is a winning number
			if winningNumbers[int(val)] {
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

	elapsed := time.Since(start)
	fmt.Println("Part 2 result: ", totalVal)
	fmt.Println("Part 2 took: ", elapsed.Microseconds(), " μs")
}
