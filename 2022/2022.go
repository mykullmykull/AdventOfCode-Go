package main

import (
	"fmt"

	"github.com/mykullmykull/AdventOfCode-Go/tree/master/2022/day1"
)

func main() {
	fmt.Print("day1a: ")
	fmt.Print(day1.MaxCalories(day1.Real))
	fmt.Println()

	fmt.Print("day1b: ")
	fmt.Println(day1.TopThreeCalories(day1.Real))
	fmt.Println()
}
