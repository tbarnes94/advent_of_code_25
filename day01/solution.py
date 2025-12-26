#!/usr/bin/env python3
"""
Advent of Code 2025 - Day 1
https://adventofcode.com/2025/day/1
"""

def parse_input(filename):
    """Parse the input file and return the data."""
    with open(filename, 'r') as f:
        lines = f.read().strip().split('\n')
    return [int(line) for line in lines]


def part1(data):
    """Solve part 1 of the puzzle - sum of all numbers."""
    return sum(data)


def part2(data):
    """Solve part 2 of the puzzle - product of all numbers."""
    result = 1
    for num in data:
        result *= num
    return result


def main():
    """Main function to run both parts of the solution."""
    # Test with example input
    example_data = parse_input('example.txt')
    print("Example Part 1:", part1(example_data))
    print("Example Part 2:", part2(example_data))
    
    # Solve with real input
    data = parse_input('input.txt')
    print("Part 1:", part1(data))
    print("Part 2:", part2(data))


if __name__ == "__main__":
    main()
