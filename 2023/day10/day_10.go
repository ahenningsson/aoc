package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// check if error is nil
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Boiler plate
	err := os.Chdir("./2023/day10")
	check(err)

	input, err := os.ReadFile("./input.txt")
	check(err)
	// fmt.Printf("MyInput: %q", string(input))
	// // End boiler plate

	lines := strings.Split(string(input), "\n")

	part1(lines)
}

func part1(lines []string) {
	start := time.Now()

	elapsed := time.Since(start)
	fmt.Println("Part 1 took: ", elapsed.Microseconds(), " μs")
}
