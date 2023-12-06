package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	err := os.Chdir("./2023/day5")
	check(err)

	input, err := os.ReadFile("./demo_input.txt")
	check(err)
	fmt.Printf("MyInput: %q", string(input))
	// // End boiler plate
	lines := strings.Split(string(input), "\n\n")
	seeds := lines[0]
	part1(seeds, lines[1:])
}

func part1(seeds string, lines []string) {
	for idx, line := range lines {
		fmt.Println(idx, line)
	}
}
