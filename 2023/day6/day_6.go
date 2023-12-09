package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// check if error is nil
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// read lines of the input file
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
	err := os.Chdir("./2023/day6")
	check(err)

	lines, err := readLines("./input.txt")
	check(err)
	// // End boiler plate

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	// Read the input file
	start := time.Now()
	raceTime := make([]int, 0, len(lines[0]))
	distance := make([]int, 0, len(lines[0]))
	for idx, line := range lines {
		val := strings.Split(line, ":")[1]
		val = strings.TrimSpace(val)
		for _, v := range strings.Fields(val) {
			numb, err := strconv.ParseInt(v, 10, 64)
			check(err)
			if idx == 0 {
				raceTime = append(raceTime, int(numb))
			}
			if idx == 1 {
				distance = append(distance, int(numb))
			}
		}
	}

	waysToWin := []int{}
	for idx, val := range raceTime {
		winCount := 0
		for s := 0; s <= val; s++ {
			// s is the seconds I hold the button
			if (s * (raceTime[idx] - s)) > distance[idx] {
				winCount++
			}
		}
		waysToWin = append(waysToWin, winCount)
	}

	result := 1
	for _, win := range waysToWin {
		result *= win
	}

	elapsed := time.Since(start)
	fmt.Println("Part 1 result:", result)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")
}

func part2(lines []string) {
	start := time.Now()
	var raceTimeStr, distanceStr strings.Builder
	for idx, line := range lines {
		val := strings.Split(line, ":")[1]
		val = strings.TrimSpace(val)
		for _, v := range strings.Fields(val) {
			if idx == 0 {
				raceTimeStr.WriteString(v)
			}
			if idx == 1 {
				distanceStr.WriteString(v)
			}
		}
	}

	raceTime, err := strconv.ParseInt(raceTimeStr.String(), 10, 64)
	check(err)
	distance, err := strconv.ParseInt(distanceStr.String(), 10, 64)
	check(err)

	// calculate the number of possible ways to win
	var s int64 // the number of seconds I hold the button
	for s = 0; s < raceTime; s++ {
		// When I hold the button for s time and I win, it means that all the options between s and endtime - s are also winning options
		// This means that I dont have to iterate through all the options
		if (s * (raceTime - s)) > distance {
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Part 2 result:", int(raceTime-(2*s+1)))
	fmt.Println("Part 2 took: ", elapsed.Microseconds(), " μs")
}
