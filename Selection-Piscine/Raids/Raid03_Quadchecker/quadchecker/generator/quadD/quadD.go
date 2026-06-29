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
			if j == 1 {
				fmt.Print("A")
			} else if j == x {
				fmt.Print("C")
			} else if i == 1 || i == y {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}