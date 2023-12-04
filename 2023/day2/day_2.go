package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
	part1(lines)

	start := time.Now()
	power := 0
	for _, line := range lines {
		gameString := strings.Split(line, ": ")[1]
		games := strings.Split(gameString, ";")
		power += part2(games)
	}
	elapsed := time.Since(start)
	fmt.Println("Part 2 took: ", elapsed.Microseconds(), " μs")
	fmt.Println("Part 2 result: ", power)
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

func part1(lines []string) {
	start := time.Now()
	sum := 0
	for idx, line := range lines {

		gameString := strings.Split(line, ": ")[1]
		games := strings.Split(gameString, ";")
		validGame := true
	OuterLoop:
		for _, set := range games {
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
	}
	elapsed := time.Since(start)
	fmt.Println("Part 1 result: ", sum)
	println("Part 1 took: ", elapsed.Microseconds(), " μs")
}

/*
PART 2
*/
func part2(games []string) int {
	colorsMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, row := range games {
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
