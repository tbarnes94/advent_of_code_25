package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ParseInput reads and parses the input file
func parseInput(filename string) ([]string, []int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	numbers := make([]int, 0, len(lines))
	directions := make([]string, 0, len(lines))

	for _, line := range lines {
		dir := string(line[0])
		directions = append(directions, dir)
		line = line[1:]

		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			return nil, nil, err
		}

		numbers = append(numbers, num)
	}

	return directions, numbers, nil
}

// SpinLock simulates the spin lock and calculates the password
func spinLock(directions []string, numbers []int) (int, error) {
	password := 0
	curr := 50
	denom := 100
	for i, num := range numbers {
		dir := directions[i]
		var times int
		if dir == "L" {
			if num > curr {
				times = (num - curr + denom - 1) / denom
			}
			curr = (denom + curr - num%denom) % denom
		} else if dir == "R" {
			if curr+num >= denom {
				times = (curr + num) / denom
			}
			curr = (curr + num) % denom
		}
		password += times
	}
	return password, nil
}

func main() {
	// Test with example input
	directions, numbers, err := parseInput("example.txt")
	if err != nil {
		fmt.Printf("Error reading example.txt: %v\n", err)
	}

	password, _ := spinLock(directions, numbers)
	fmt.Printf("Example password: %d\n", password)

	// Solve with real input
	directions, numbers, err = parseInput("input.txt")
	if err != nil {
		fmt.Printf("Error reading input.txt: %v\n", err)
		os.Exit(1)
	}

	password, _ = spinLock(directions, numbers)
	fmt.Printf("Input password: %d\n", password)
}
