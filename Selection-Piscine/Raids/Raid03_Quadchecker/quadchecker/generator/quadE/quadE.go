package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 && j == 1) || (i == y && j == x) {
				fmt.Print("A")
			} else if (i == 1 && j == x) || (i == y && j == 1) {
				fmt.Print("C")
			} else if i == 1 || i == y || j == 1 || j == x {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}