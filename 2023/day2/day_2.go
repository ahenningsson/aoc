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

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

func checkColor(color string, val int) bool {
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

func part1(gameRows []string, idx int) int {
	validGame := true
	sum := 0

	for _, set := range gameRows {
		set = strings.TrimSpace(set)
		colorPair := strings.Split(set, ",")

		for _, cvpair := range colorPair {
			cvpair = strings.TrimSpace(cvpair)
			val := strings.Split(cvpair, " ")

			valInt, err := strconv.Atoi(val[0])
			check(err)

			if checkColor(val[1], valInt) && validGame {
				validGame = true
			} else {
				validGame = false
			}
		}
	}
	if validGame {
		sum += idx + 1
	}

	return sum
}

// part 2
func part2(gameRows []string) int {
	red := 0
	green := 0
	blue := 0
	for _, row := range gameRows {
		row = strings.TrimSpace(row)
		colorPairs := strings.Split(row, ",")

		for _, pair := range colorPairs {
			pair = strings.TrimSpace(pair)
			val := strings.Split(pair, " ")

			valInt, err := strconv.Atoi(val[0])
			check(err)

			if val[1] == "red" && valInt > red {
				red = valInt
			}

			if val[1] == "green" && valInt > green {
				green = valInt
			}

			if val[1] == "blue" && valInt > blue {
				blue = valInt
			}
		}
	}
	return red * green * blue
}
