package main

import (
	"fmt"

	"github.com/mykullmykull/AdventOfCode-Go/tree/master/2022/day1"
	"github.com/mykullmykull/AdventOfCode-Go/tree/master/2022/day2"
)

func main() {
	fmt.Println()
	fmt.Print("day1a: ")
	fmt.Print(day1.MaxCalories(day1.Real))

	fmt.Println()
	fmt.Print("day1b: ")
	fmt.Print(day1.TopThreeCalories(day1.Real))

	fmt.Println()
	fmt.Print("day2a: ")
	fmt.Print(day2.FinalScore(day2.Real))

	fmt.Println()
	fmt.Print("day2b: ")
	fmt.Print(day2.FinalScoreB(day2.Real))
}
