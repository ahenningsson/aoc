package main

import (
	"fmt"
	"main/utils"
	"os"
	"strconv"
	"time"
)

func main() {
	err := os.Chdir("./2025/day3")
	utils.CheckErr(err)

	lines, err := utils.ReadLines("./input.txt")
	utils.CheckErr(err)

	values := make([][]int64, len(lines))
	for i, line := range lines {
		vals := make([]int64, len(line))
		for j := 0; j < len(line); j++ {
			// Works if values are between 0 and 9.
			// For example 3 has value 51, '0' is 48. 51-48 = 3
			// This is the fastest int conversion
			vals[j] = int64(line[j] - '0')
		}
		values[i] = vals
	}

	start := time.Now()
	sum := part1(values)
	elapsed := time.Since(start)

	fmt.Println("Part 1 result: ", sum)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")

	start2 := time.Now()
	sum2 := part2(values)
	elapsed2 := time.Since(start2)

	fmt.Println("Part 2 result: ", sum2)
	fmt.Println("Part 2 took: ", elapsed2.Microseconds(), " μs")
}

func part1(values [][]int64) int64 {

	results := []int64{}

	for _, bank := range values {
		highestNumb := int64(0)
		for i, battery1 := range bank {
			for _, battery2 := range bank[i+1:] {
				maxbatstr := fmt.Sprintf("%d%d", battery1, battery2)
				maxbat, err := strconv.ParseInt(maxbatstr, 10, 64)
				utils.CheckErr(err)

				if maxbat > highestNumb {
					highestNumb = maxbat
				}
			}
		}
		results = append(results, highestNumb)
	}

	return Sum(results)
}

func part2(values [][]int64) int64 {

	results := []int64{}

	for _, bank := range values {
		results = append(results, evaluateWindow(bank))
	}

	return Sum(results)
}

func evaluateWindow(bank []int64) int64 {
	wRes := []byte{}

	wStart := 0

	for rem := range 12 {
		wEnd := len(bank) - (12 - rem)
		maxDig := int64(0)
		maxIdx := 0

		// Get the highest value of current window
		for i := wStart; i <= wEnd; i++ {
			if bank[i] > maxDig {
				maxDig = bank[i]
				maxIdx = i
			}
		}

		wRes = append(wRes, byte('0')+byte(maxDig))
		wStart = maxIdx + 1
	}

	res, err := strconv.ParseInt(string(wRes), 10, 64)
	utils.CheckErr(err)

	return res
}

func Sum(vals []int64) int64 {
	res := int64(0)
	for _, val := range vals {
		res += val
	}

	return res
}
