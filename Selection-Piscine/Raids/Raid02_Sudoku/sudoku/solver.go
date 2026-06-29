package main

import "github.com/01-edu/z01"

const size = 9

type Grid [size][size]int

func ParseGrid(args []string) (Grid, bool) {

	var grid Grid

	if len(args) != size {

		return grid, false

	}

	for row := 0; row < size; row++ {

		if len(args[row]) != size {

			return grid, false

		}

		for col := 0; col < size; col++ {

			char := args[row][col]

			if char == '.' {

				continue

			}

			if char < '1' || char > '9' {

				return grid, false

			}

			value := int(char - '0')

			if !isSafe(grid, row, col, value) {

				return grid, false

			}

			grid[row][col] = value

		}

	}

	return grid, true

}

func SolveSudoku(grid Grid) (Grid, bool) {

	solutions := 0

	var solved Grid

	search(&grid, &solutions, &solved)

	return solved, solutions == 1

}

func search(grid *Grid, solutions *int, solved *Grid) {

	if *solutions > 1 {

		return

	}

	row, col, found := findEmpty(*grid)

	if !found {

		*solutions++

		if *solutions == 1 {

			*solved = *grid

		}

		return

	}

	for value := 1; value <= 9; value++ {

		if !isSafe(*grid, row, col, value) {

			continue

		}

		grid[row][col] = value

		search(grid, solutions, solved)

		grid[row][col] = 0

	}

}

func findEmpty(grid Grid) (int, int, bool) {

	for row := 0; row < size; row++ {

		for col := 0; col < size; col++ {

			if grid[row][col] == 0 {

				return row, col, true

			}

		}

	}

	return 0, 0, false

}

func isSafe(grid Grid, row, col, value int) bool {

	for i := 0; i < size; i++ {

		if grid[row][i] == value {

			return false

		}

		if grid[i][col] == value {

			return false

		}

	}

	startRow := (row / 3) * 3

	startCol := (col / 3) * 3

	for r := startRow; r < startRow+3; r++ {

		for c := startCol; c < startCol+3; c++ {

			if grid[r][c] == value {

				return false

			}

		}

	}

	return true

}

func PrintGrid(grid Grid) {

	for row := 0; row < size; row++ {

		for col := 0; col < size; col++ {

			if col > 0 {

				z01.PrintRune(' ')

			}

			z01.PrintRune(rune(grid[row][col]) + '0')

		}

		z01.PrintRune('\n')

	}

}

func PrintError() {

	for _, char := range "Error\n" {

		z01.PrintRune(char)

	}

}
