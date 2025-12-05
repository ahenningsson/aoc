package main

import (
	"fmt"
	"main/utils"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	err := os.Chdir("./2025/day5")
	utils.CheckErr(err)

	lines, err := utils.ReadLines("./new_input.txt")
	utils.CheckErr(err)

	fresh := [][]int64{}
	spoiled := []int64{}

	freshLine := true
	for _, line := range lines {

		if line == "" {
			freshLine = false
			continue
		}

		if freshLine {
			// Read fresh
			freshRange := strings.Split(line, "-")

			f1, err := strconv.ParseInt(freshRange[0], 10, 64)
			utils.CheckErr(err)
			f2, err := strconv.ParseInt(freshRange[1], 10, 64)
			utils.CheckErr(err)

			vals := []int64{f1, f2}

			fresh = append(fresh, vals)

		} else {
			// Read spolied
			s, err := strconv.ParseInt(line, 10, 64)
			utils.CheckErr(err)
			spoiled = append(spoiled, s)
		}

	}

	start := time.Now()
	sum := part1(fresh, spoiled)
	elapsed := time.Since(start)

	start2 := time.Now()
	sum2 := part2(fresh)
	elapsed2 := time.Since(start2)

	fmt.Println("Part 1 result: ", sum)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")

	fmt.Println("Part 2 result: ", sum2)
	fmt.Println("Part 2 took: ", elapsed2.Microseconds(), " μs")
}

func part1(fresh [][]int64, spoiled []int64) int {

	res := 0

	for _, s := range spoiled {
		for _, franges := range fresh {
			f1 := franges[0]
			f2 := franges[1]

			if s >= f1 && s <= f2 {
				res++
				break
			}

		}
	}
	return res
}

func part2(fresh [][]int64) int64 {

	// sort fresh according on start value
	sort.Slice(fresh, func(i, j int) bool {
		return fresh[i][0] < fresh[j][0]
	})

	// Merge the overlaps
	isMerge := true
	for isMerge {

		for i, curr := range fresh {
			if i == len(fresh)-1 {
				continue
			}
			e1 := curr[1]

			s2 := fresh[i+1][0]
			e2 := fresh[i+1][1]

			if e1 >= s2 {
				// The ending of the current range should be changed to the max value of the ending from current or the next
				curr[1] = max(e1, e2)
				fresh = remove(fresh, i+1)
				isMerge = true
				break
			} else {
				isMerge = false
			}
		}
	}

	res := int64(0)
	// Calculate the length: (b-a)+1
	for _, r := range fresh {
		res += r[1] - r[0] + 1
	}

	return res
}

func remove(slice [][]int64, index int) [][]int64 {
	return append(slice[:index], slice[index+1:]...)
}
