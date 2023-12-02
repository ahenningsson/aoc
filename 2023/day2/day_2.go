package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// boiler plate
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// boiler plate
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
	err := os.Chdir("./2023/day2")
	check(err)

	lines, err := readLines("./input.txt")
	check(err)
	// End boiler plate

	sum := 0
	power := 0

	for idx, line := range lines {

		gameString := strings.Split(line, ": ")[1]
		games := strings.Split(gameString, ";")

		sum += part1(games, idx)
		power += part2(games)

	}

	println("Part 1: ", sum)
	println("Part 2: ", power)
}

/*
PART 1
*/
const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

func isCubeCountValid(color string, val int) bool {
	if color == "red" {
		if val > MAX_RED {
			return false
		}
	}
	if color == "green" {
		if val > MAX_GREEN {
			return false
		}
	}
	if color == "blue" {
		if val > MAX_BLUE {
			return false
		}
	}
	return true
}

func part1(gameRows []string, idx int) int {
	sum := 0
	validGame := true
OuterLoop:
	for _, set := range gameRows {
		colorPair := strings.Split(strings.TrimSpace(set), ",")
		for _, cvpair := range colorPair {
			cvpair = strings.TrimSpace(cvpair)
			val := strings.Split(cvpair, " ")

			valInt, err := strconv.Atoi(val[0])
			check(err)

			if !isCubeCountValid(val[1], valInt) {
				validGame = false
				break OuterLoop
			}
		}
	}
	if validGame {
		sum += idx + 1
	}
	return sum
}

/*
PART 2
*/
func part2(gameRows []string) int {
	colorsMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, row := range gameRows {
		row = strings.TrimSpace(row)
		colorPairs := strings.Split(row, ",")

		for _, pair := range colorPairs {
			pair = strings.TrimSpace(pair)
			val := strings.Split(pair, " ")

			valInt, err := strconv.Atoi(val[0])
			check(err)

			if valInt > colorsMap[val[1]] {
				colorsMap[val[1]] = valInt
			}
		}
	}
	return colorsMap["red"] * colorsMap["green"] * colorsMap["blue"]
}
