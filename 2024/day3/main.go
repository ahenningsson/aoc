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
	err := os.Chdir("../2024/day3")
	utils.CheckErr(err)

	start := time.Now()
	value := part1()
	elapsed := time.Since(start)

	start2 := time.Now()
	value2 := part2()
	elapsed2 := time.Since(start2)

	fmt.Println("Part 1 result: ", value)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")

	fmt.Println("Part 2 result: ", value2)
	fmt.Println("Part 2 took: ", elapsed2.Microseconds(), " μs")
}

func part1() int {
	lines, err := utils.ReadLines("./input.txt")
	utils.CheckErr(err)
	reg, _ := regexp.Compile(`mul[(](\d+),(\d+)[)]`)

	var matches [][]string
	sum := 0
	for _, line := range lines {
		matches = reg.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			numb1, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println("Error parsing numbers:", err)
				return 0
			}
			numb2, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Println("Error parsing numbers:", err)
				return 0
			}

			sum += numb1 * numb2
		}
	}

	return sum
}

func part2() int {
	lines, err := utils.ReadLines("./input.txt")
	utils.CheckErr(err)
	reg, _ := regexp.Compile(`don't\(\)|do\(\)|mul\((\d+),(\d+)\)`)

	keepGoing := true
	sum := 0
	for _, line := range lines {
		matches := reg.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "don't()" {
				keepGoing = false
			} else if match[0] == "do()" {
				keepGoing = true
			}

			if keepGoing && strings.Contains(match[0], "mul") {
				numb1, err := strconv.Atoi(match[1])
				if err != nil {
					fmt.Println("Error parsing numbers:", err)
					return 0
				}
				numb2, err := strconv.Atoi(match[2])
				if err != nil {
					fmt.Println("Error parsing numbers:", err)
					return 0
				}

				sum += numb1 * numb2
			}
		}
	}

	return sum
}
