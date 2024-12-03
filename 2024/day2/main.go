package main

import (
	"fmt"
	"main/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

func getIntSlice(line string) []int {
	var slice []int

	parts := strings.Split(line, " ")

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
	err := os.Chdir("../2024/day2")
	utils.CheckErr(err)

	lines, err := utils.ReadLines("./input.txt")
	utils.CheckErr(err)

	start := time.Now()
	safeReports := 0
	for _, line := range lines {

		intSlice := getIntSlice(line)
		safeReports += part1(intSlice)
	}
	elapsed := time.Since(start)

	fmt.Println("Part 1 result: ", safeReports)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")

	start = time.Now()
	safeReports = 0
	for _, line := range lines {

		intSlice := getIntSlice(line)
		safeReports += part2(intSlice)
	}
	elapsed = time.Since(start)

	fmt.Println("Part 2 result: ", safeReports)
	fmt.Println("Part 2 took: ", elapsed.Microseconds(), " μs")
}

func part1(reportSlice []int) int {

	safeReports := 0
	val := isSafe(reportSlice)
	if val {
		safeReports++
	}

	return safeReports
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isSafe(slice []int) bool {
	trend := ""
	for i := 1; i < len(slice); i++ {
		res, cat := compareFunc(slice[i-1], slice[i])
		if !res {
			return false
		}

		if trend == "" {
			trend = cat
		} else if trend != cat {
			return false
		}
	}

	return true
}

func compareFunc(a, b int) (bool, string) {
	diff := Abs(a - b)

	if diff > 3 || diff < 1 {
		return false, "neither"
	}

	if a > b {
		return true, "decrease"
	} else if b > a {
		return true, "increase"
	} else {
		return false, "neither"
	}
}

func part2(reportSlice []int) int {
	safeReports := 0

	for i := range reportSlice {
		newSlice := append([]int{}, reportSlice[:i]...)
		newSlice = append(newSlice, reportSlice[i+1:]...)

		if isSafe(newSlice) {
			safeReports++
			break
		}
	}
	return safeReports
}
