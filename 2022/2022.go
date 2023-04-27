package main

import (
	"fmt"

	"github.com/mykullmykull/AdventOfCode-Go/tree/master/2022/day1"
	"github.com/mykullmykull/AdventOfCode-Go/tree/master/2022/day2"
	"github.com/mykullmykull/AdventOfCode-Go/tree/master/2022/day3"
	"github.com/mykullmykull/AdventOfCode-Go/tree/master/2022/day4"
	"github.com/mykullmykull/AdventOfCode-Go/tree/master/2022/day5"
	"github.com/mykullmykull/AdventOfCode-Go/tree/master/2022/day6"
	"github.com/mykullmykull/AdventOfCode-Go/tree/master/2022/day7"
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

	fmt.Println()
	fmt.Print("day3a: ")
	fmt.Print(day3.SumPrioritiesOfMatchingTypes(day3.Real))

	fmt.Println()
	fmt.Print("day3b: ")
	fmt.Print(day3.SumPrioritiesOfBadges(day3.Real))

	fmt.Println()
	fmt.Print("day4a: ")
	fmt.Print(day4.CountOverlaps(day4.Real, 1))

	fmt.Println()
	fmt.Print("day4b: ")
	fmt.Print(day4.CountOverlaps(day4.Real, 2))

	fmt.Println()
	fmt.Print("day5a: ")
	fmt.Print(day5.TopCratesAfterMoves(day5.RealCrates, day5.RealMoves, 1))

	fmt.Println()
	fmt.Print("day5b: ")
	fmt.Print(day5.TopCratesAfterMoves(day5.RealCrates, day5.RealMoves, 2))

	fmt.Println()
	fmt.Print("day6a: ")
	fmt.Print(day6.FirstIndexAfterFourUnique(day6.Real))

	fmt.Println()
	fmt.Print("day6b: ")
	fmt.Print(day6.FirstIndexAfterFourteenUnique(day6.Real))

	fmt.Println()
	fmt.Print("day7a: ")
	fmt.Print(day7.SumSizeOfBigDirectories(day7.Test))
}
