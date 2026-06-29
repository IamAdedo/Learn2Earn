package main

import (
	"os"
)

func main() {

	grid, ok := ParseGrid(os.Args[1:])

	if !ok {

		PrintError()

		return

	}

	solved, unique := SolveSudoku(grid)

	if !unique {

		PrintError()

		return

	}

	PrintGrid(solved)

}
