# Advent of Code 2025

My advent of code project for 2025. See https://adventofcode.com/

## Structure

This repository contains solutions for all 25 days of Advent of Code 2025. Each day has its own directory with the following structure:

```
dayXX/
├── solution.py    # Python solution for both parts
├── input.txt      # Your personal puzzle input
└── example.txt    # Example input from puzzle description
```

## Usage

### Running a Specific Day

To run the solution for a specific day:

```bash
python run.py <day_number>
```

Example:
```bash
python run.py 1    # Runs day 1
python run.py 15   # Runs day 15
```

### Running All Days

To run all implemented solutions:

```bash
python run.py all
```

### Running an Individual Solution

You can also run each day's solution directly:

```bash
cd day01
python solution.py
```

## Workflow

For each day:

1. Visit https://adventofcode.com/2025/day/X (where X is the day number)
2. Read the puzzle description
3. Copy the example input into `dayXX/example.txt`
4. Copy your personal puzzle input into `dayXX/input.txt`
5. Implement your solution in `dayXX/solution.py`
6. Run and test your solution

## Requirements

- Python 3.6 or higher
- No external dependencies required (solutions use only standard library)

## Notes

- Each solution is self-contained in its day directory
- Input files are git-ignored by default (except placeholder comments)
- Solutions should work with both example and actual puzzle inputs
