package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

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
	err := os.Chdir("../2024/day1")
	check(err)

	lines, err := readLines("./input.txt")
	check(err)

	var slice1 []int
	var slice2 []int

	for _, line := range lines {
		parts := strings.Fields(line)

		leftNum, err1 := strconv.Atoi(parts[0])
		rightNum, err2 := strconv.Atoi(parts[1])
		if err1 == nil && err2 == nil {
			slice1 = append(slice1, leftNum)
			slice2 = append(slice2, rightNum)
		} else {
			fmt.Println("Error parsing numbers:", err1, err2)
		}
	}

	start := time.Now()
	sum := part1(slice1, slice2)
	elapsed := time.Since(start)

	fmt.Println("Part 1 result: ", sum)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")

	start = time.Now()
	sum = part2(slice1, slice2)
	elapsed = time.Since(start)

	fmt.Println("Part 2 result: ", sum)
	fmt.Println("Part 2 took: ", elapsed.Microseconds(), " μs")
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(slice1, slice2 []int) int {
	if len(slice1) == 0 || len(slice2) == 0 {
		return 0
	}

	sort.Slice(slice1, func(i, j int) bool {
		return slice1[i] < slice1[j]
	})
	sort.Slice(slice2, func(i, j int) bool {
		return slice2[i] < slice2[j]
	})

	sum := 0

	for i := range slice1 {
		sum += Abs(slice2[i] - slice1[i])
	}

	return sum
}

func part2(slice1, slice2 []int) int {
	if len(slice1) == 0 || len(slice2) == 0 {
		return 0
	}

	simularityScore := 0

	for i := range slice1 {

		occurrences := 0
		for _, val := range slice2 {
			if slice1[i] == val {
				occurrences += 1
			}
		}
		simularityScore += slice1[i] * occurrences
		occurrences = 0
	}

	return simularityScore
}
