package main

import (
	"fmt"
	"main/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	err := os.Chdir("./2025/day2")
	utils.CheckErr(err)

	lines, err := utils.ReadLines("./input.txt")
	utils.CheckErr(err)

	values := []string{}
	for _, line := range lines {
		str := strings.Split(line, ",")

		for _, val := range str {
			if val != "" {
				values = append(values, val)
			}
		}
	}

	start := time.Now()
	sum := part1(values)
	elapsed := time.Since(start)

	start2 := time.Now()
	sum2 := part2(values)
	elapsed2 := time.Since(start2)

	fmt.Println("Part 1 result: ", sum)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")

	fmt.Println("Part 2 result: ", sum2)
	fmt.Println("Part 2 took: ", elapsed2.Microseconds(), " μs")
}

func part1(values []string) int64 {

	invalidIDs := []int64{}

	for _, val := range values {
		r := strings.Split(val, "-")
		start, err := strconv.ParseInt(r[0], 10, 64)
		utils.CheckErr(err)
		end, err := strconv.ParseInt(r[1], 10, 64)
		utils.CheckErr(err)

		for x := start; x <= end; x++ {

			if evaluateIfInvalid(x) {
				invalidIDs = append(invalidIDs, x)
			}
		}
	}

	return sum(invalidIDs)
}

// Evaluate if a value should be added to invalid list
func evaluateIfInvalid(val int64) bool {

	str := strconv.Itoa(int(val))
	if str[:1] == "0" {
		return true
	}

	// len of string
	strLength := len(str)
	firstString := str[:strLength/2]
	lastString := str[strLength/2:]

	if firstString == lastString {
		return true
	}

	return false
}

func part2(values []string) int64 {

	invalidIDs := []int64{}

	for _, val := range values {
		r := strings.Split(val, "-")
		start, err := strconv.ParseInt(r[0], 10, 64)
		utils.CheckErr(err)
		end, err := strconv.ParseInt(r[1], 10, 64)
		utils.CheckErr(err)

		for x := start; x <= end; x++ {

			if evaluateIfInvalidPart2(x) {
				invalidIDs = append(invalidIDs, x)
			}
		}
	}

	return sum(invalidIDs)
}

func sum(input []int64) int64 {
	sum := int64(0)

	for _, x := range input {
		sum += x
	}

	return sum
}

// Evaluate if a value should be added to invalid list
func evaluateIfInvalidPart2(val int64) bool {

	str := strconv.Itoa(int(val))
	if str[:1] == "0" {
		// fmt.Println("Leading 0:", val)
		return true
	}

	// len of string
	strLength := len(str)
	if strLength%2 == 0 {
		firstString := str[:strLength/2]
		lastString := str[strLength/2:]

		if firstString == lastString {
			// fmt.Println("String equal first and last:", val)
			return true
		}
	}

	// If a pattern is repeated its invalid
	// Check if all characters are same
	allSame := false
	for i := 0; i < strLength-1; i++ {
		if str[i] == str[i+1] {
			allSame = true
		} else {
			allSame = false
			break
		}
	}
	if allSame {
		// fmt.Println("All is same:", val)
		return true
	}

	// Check if the number has a repeating pattern, example: 123123123 => 123 is repeated 3 times

	for i := range str {

		compStr := fmt.Sprintf(`\W*(%s)*`, str[:i])
		if len(str[:i]) == 0 {
			continue
		}

		// If the index is higher than half the strings total length, there cant be any patterns
		if i > len(str)/2 {
			break
		}

		r, _ := regexp.Compile(compStr)
		result := r.FindAllString(str, -1)
		if len(result[0]) == len(str) {
			return true
		}

	}

	return false
}
