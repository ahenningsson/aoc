package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	err := os.Chdir("./2023/day1")
	check(err)
	// End boiler plate

	// Part 1
	lines, err := readLines("./input.txt")
	check(err)

	sum := 0
	rePart1 := regexp.MustCompile("[0-9]+")
	start := time.Now()
	for _, line := range lines {
		strSlice := rePart1.FindAllString(line, -1)

		if len(strSlice) == 0 {
			continue
		}

		sum += part1(strSlice)
	}

	elapsed := time.Since(start)
	fmt.Println("Part 1 result: ", sum)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")

	// Part 2
	lines, err = readLines("./input.txt")
	check(err)

	start = time.Now()
	numberStringSlice := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	reNumbers := regexp.MustCompile("[0-9]")
	sum = 0
	for _, line := range lines {
		result := ""

		for i := 0; i < len(line); i++ {
			if len(line) == 0 {
				continue
			}

			// Check if the current character (line[i]) is a number

			if reNumbers.MatchString(string(line[i])) {
				result += string(line[i])
			}

			for idx, val := range numberStringSlice {
				if strings.HasPrefix(line[i:], val) {
					result += strconv.Itoa(idx + 1) // Index of the numberStringSlice is -1 from the number. 1 is at index 0, 2 is at index 1, etc.
				}
			}
		}
		// since part1 solution takes a []string, i need to split the result string
		sum += part1(strings.Split(result, ""))
	}
	elapsed = time.Since(start)
	fmt.Println("Part 2 result: ", sum)
	fmt.Println("Part 2 took: ", elapsed.Microseconds(), " μs")
}

func part1(strSlice []string) int {
	// get first number
	first := strSlice[0]
	if len(first) > 1 {
		first = first[:1]
	}

	// get last number
	last := strSlice[len(strSlice)-1]
	if len(last) > 1 {
		last = last[len(last)-1:]
	}

	strMerge := first + last
	intMerge, err := strconv.Atoi(strMerge)
	check(err)

	return intMerge
}
