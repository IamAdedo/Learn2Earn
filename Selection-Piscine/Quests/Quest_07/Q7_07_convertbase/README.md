# Instructions

Write a function ConvertBase that converts a number represented as a string from one base to another.

The function receives three arguments:

nbr: a string representing a numeric value in a base.

baseFrom: a string representing the base nbr it's using.

baseTo: a string representing the base nbr should be represented in the returned value.

`baseFrom` and `baseTo` will always be valid bases. Negative numbers will not be tested.


## Expected function

func ConvertBase(nbr, baseFrom, baseTo string) string {

}

## Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	result := piscine.ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
}


## And its output:

$ go run .
43
$
