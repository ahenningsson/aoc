package main

import (
	"fmt"
	"main/utils"
	"os"
	"strconv"
	"time"
)

func main() {
	err := os.Chdir("./2025/day1")
	utils.CheckErr(err)

	lines, err := utils.ReadLines("./input.txt")
	utils.CheckErr(err)

	start := time.Now()
	sum := part1(lines)
	elapsed := time.Since(start)

	fmt.Println("Part 1 result: ", sum)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")
}

func part1(lines []string) int {
	count := 0
	pointing := int64(50)

	for _, line := range lines {
		rot := line[:1]
		val, err := strconv.ParseInt(line[1:], 10, 64)
		utils.CheckErr(err)

		switch rot {
		case "L":
			pointing -= val

		case "R":
			pointing += val
		}

		if pointing%100 == 0 {
			count += 1
		}
	}
	return count
}
