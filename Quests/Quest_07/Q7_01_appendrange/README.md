# Instructions

Write a function that takes two integers, `min` and `max`, as parameters and returns a slice of integers containing all values from `min` (inclusive) to `max` (exclusive).

If `min` is greater than or equal to `max`, the function must return a `nil` slice.

The use of `make` is not allowed for this exercise.

## Expected function

func AppendRange(min, max int) []int {

}

## Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AppendRange(5, 10))
	fmt.Println(piscine.AppendRange(10, 5))
}


## And its output:

$ go run .
[5 6 7 8 9]
[]
$
