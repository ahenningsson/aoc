package main

import (
	"fmt"
	"main/utils"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func getIntSlice(line string) []int {
	var slice []int

	parts := strings.Split(line, ",")

	for _, c := range parts {
		numb, err1 := strconv.Atoi(c)
		if err1 != nil {
			fmt.Println("Error parsing numbers:", err1)
			return []int{}
		} else {
			slice = append(slice, numb)

		}
	}

	return slice
}

func main() {
	err := os.Chdir("../2024/day5")
	utils.CheckErr(err)

	lines, err := utils.ReadLines("./input.txt")
	utils.CheckErr(err)

	var rulesMap = make(map[int][]int)
	var updateSlice [][]int

	addToRules := true
	for _, line := range lines {
		if len(line) == 0 {
			addToRules = false
			continue
		}

		if addToRules {
			parts := strings.Split(line, "|")
			key, err1 := strconv.Atoi(parts[0])
			if err1 != nil {
				fmt.Println("Error parsing numbers:", err1)
				return
			}
			value, err1 := strconv.Atoi(parts[1])
			if err1 != nil {
				fmt.Println("Error parsing numbers:", err1)
				return
			}

			rulesMap[key] = append(rulesMap[key], value)
		} else {
			intSlice := getIntSlice(line)
			updateSlice = append(updateSlice, intSlice)
		}
	}

	start := time.Now()
	val := part1(rulesMap, updateSlice)
	elapsed := time.Since(start)

	start2 := time.Now()
	val2 := part2(rulesMap, updateSlice)
	elapsed2 := time.Since(start2)

	fmt.Println("Part 1 result: ", val)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")

	fmt.Println("Part 2 result: ", val2)
	fmt.Println("Part 2 took: ", elapsed2.Microseconds(), " μs")
}

func part1(rules map[int][]int, updates [][]int) int {

	var correctUpdates [][]int
	for _, updateLine := range updates {
		isCorrect := checkIfCorrect(updateLine, rules)
		if isCorrect {
			correctUpdates = append(correctUpdates, updateLine)
		}
	}

	// Calculate sum
	sum := 0
	for _, correctSlice := range correctUpdates {
		sum += correctSlice[len(correctSlice)/2]
	}

	return sum
}

func checkIfCorrect(updateLine []int, rules map[int][]int) bool {
	isCorrect := true
	for i, update := range updateLine {
		// Check if other values are in rules
		for _, k := range updateLine[i+1:] {
			if !slices.Contains(rules[update], k) {
				isCorrect = false
			}
		}
	}
	return isCorrect
}

func part2(rules map[int][]int, updates [][]int) int {

	var inCorrectUpdates [][]int
	for _, updateLine := range updates {
		isCorrect := true
		isCorrect = checkIfCorrect(updateLine, rules)
		if !isCorrect {
			inCorrectUpdates = append(inCorrectUpdates, updateLine)
		}
	}
	// Fix incorrect updates recursively
	for _, updateLine := range inCorrectUpdates {
		changed := true

		for changed {
			changed = false

			for j, update := range updateLine {

				for _, k := range updateLine[j+1:] {
					if !slices.Contains(rules[update], k) {
						updateLine[j], updateLine[j+1] = updateLine[j+1], updateLine[j]
						changed = true
						break
					}
				}
			}

			if checkIfCorrect(updateLine, rules) {
				break
			}
		}

	}

	// Calculate sum
	sum := 0
	for _, InCorrectSlice := range inCorrectUpdates {
		sum += InCorrectSlice[len(InCorrectSlice)/2]
	}

	return sum
}
