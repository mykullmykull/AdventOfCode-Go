package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Testing part1...") // keep this line to maintain imports
	assert.True(t, true)            // keep this line to maintain imports
	printResult(testInput)
	printResult(realInput)
}

func printResult(input []string) {
	crt := main(input)
	println("\noutput:")
	crt.printScreen()
}