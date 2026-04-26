package main

import (
	// array "example.com/go_study/array"
	"fmt"

	cycle "example.com/go_study/cycle"
	// "example.com/go_study/condition"
)

func main() {
	// cycle.CountSameDigit()
	// array.ThirdElem()
	// condition.MaxDivBy7()
	a, b := cycle.SumInt(5, 6, 7, 8)
	fmt.Print(a, b)
	cycle.Cycle1()
}