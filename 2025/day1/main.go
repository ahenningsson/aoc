package main

import (
	"fmt"
	"main/utils"
	"math"
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

	start2 := time.Now()
	sum2 := part2(lines)
	elapsed2 := time.Since(start2)

	fmt.Println("Part 1 result: ", sum)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")

	fmt.Println("Part 2 result: ", sum2)
	fmt.Println("Part 2 took: ", elapsed2.Microseconds(), " μs")
}

func Mod(a, b int64) int64 {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}

	return m
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

func part2(lines []string) int {
	count := 0
	pointing := int64(50)

	for _, line := range lines {
		rot := line[:1]
		val, err := strconv.ParseInt(line[1:], 10, 64)
		utils.CheckErr(err)

		wholeLaps, tempval := evaluateSteps(val)
		count += int(wholeLaps)
		val = tempval

		switch rot {
		case "L":
			// This is stupid
			if pointing == 0 {
				count--
			}
			pointing -= val
			if pointing == 0 {
				count++
			} else if pointing < 0 {
				count++
				pointing = 100 - utils.Abs(pointing)
			}

		case "R":
			pointing += val
			if pointing == 100 {
				count++
				pointing = 0
			} else if pointing > 100 {
				count++
				pointing = utils.Abs(pointing) - 100
			}
		}

	}
	return count
}

func evaluateSteps(val int64) (int64, int64) {
	temp := Mod(val, 100)
	wholeLaps, _ := math.Modf(float64(val / 100))

	return int64(wholeLaps), temp
}
