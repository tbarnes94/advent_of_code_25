#!/usr/bin/env python3
"""
Runner script for Advent of Code 2025
Run a specific day's solution or all solutions.
"""

import sys
import os
import importlib.util


def run_day(day_num):
    """Run the solution for a specific day."""
    day_dir = f"day{day_num:02d}"
    solution_path = os.path.join(day_dir, "solution.py")
    
    if not os.path.exists(solution_path):
        print(f"Error: {solution_path} does not exist")
        return False
    
    print(f"\n{'='*60}")
    print(f"Running Day {day_num}")
    print(f"{'='*60}")
    
    # Import and run the solution
    spec = importlib.util.spec_from_file_location(f"day{day_num}", solution_path)
    module = importlib.util.module_from_spec(spec)
    
    # Change to the day's directory so relative file paths work
    original_dir = os.getcwd()
    os.chdir(day_dir)
    
    try:
        spec.loader.exec_module(module)
        if hasattr(module, 'main'):
            module.main()
        else:
            print("Warning: No main() function found in solution.py")
    except Exception as e:
        print(f"Error running day {day_num}: {e}")
    finally:
        os.chdir(original_dir)
    
    return True


def main():
    """Main function to handle command line arguments."""
    if len(sys.argv) < 2:
        print("Usage: python run.py <day_number>")
        print("Example: python run.py 1")
        print("         python run.py all  (to run all days)")
        sys.exit(1)
    
    arg = sys.argv[1]
    
    if arg.lower() == "all":
        for day in range(1, 26):
            run_day(day)
    else:
        try:
            day = int(arg)
            if day < 1 or day > 25:
                print("Error: Day must be between 1 and 25")
                sys.exit(1)
            run_day(day)
        except ValueError:
            print("Error: Invalid day number")
            sys.exit(1)


if __name__ == "__main__":
    main()
