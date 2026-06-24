# Instructions

Write a function ConcatParams that takes a slice of strings as input and returns a single string made by concatenating all elements of the slice separated by a newline (\n).

## Expected function

func ConcatParams(args []string) string {

}

## Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(piscine.ConcatParams(test))
}

## And its output:

$ go run .
Hello
how
are
you?
$
