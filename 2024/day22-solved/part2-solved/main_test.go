package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Testing...") // keep this line to maintain imports
	assert.True(t, true)      // keep this line to maintain imports
	assert.Equal(t, expected, main(testInput, 2000))
	fmt.Printf("output: %d\n", main(realInput, 2000))
}

func Test_getPrice(t *testing.T) {
	assert.Equal(t, 0, getPrice(9670))
	assert.Equal(t, 1, getPrice(2340981))
	assert.Equal(t, 2, getPrice(324082))
	assert.Equal(t, 3, getPrice(6897283))
	assert.Equal(t, 4, getPrice(234))
	assert.Equal(t, 5, getPrice(50985))
	assert.Equal(t, 6, getPrice(6))
	assert.Equal(t, 7, getPrice(587))
	assert.Equal(t, 8, getPrice(948278))
	assert.Equal(t, 9, getPrice(999999999))
}

func Test_getPricesAndChanges(t *testing.T) {
	prices, changes := getPricesAndChanges(123, 9)
	assert.Equal(t, []int{0, 6, 5, 4, 4, 6, 4, 4, 2}, prices)
	assert.Equal(t, []int{-3, 6, -1, -1, 0, 2, -2, 0, -2}, changes)
}

func Test_buyBananasFromMonkey(t *testing.T) {
	prices, changes := getPricesAndChanges(123, 10)
	m := monkey{123, prices, changes}
	assert.Equal(t, 6, buyBananasFromMonkey(m, seq{-1, -1, 0, 2}))

	s := seq{-2, 1, -1, 3}
	prices, changes = getPricesAndChanges(1, 2000)
	m = monkey{1, prices, changes}
	assert.Equal(t, 7, buyBananasFromMonkey(m, s))

	prices, changes = getPricesAndChanges(2, 2000)
	m = monkey{2, prices, changes}
	assert.Equal(t, 7, buyBananasFromMonkey(m, s))

	prices, changes = getPricesAndChanges(3, 2000)
	m = monkey{3, prices, changes}
	assert.Equal(t, 0, buyBananasFromMonkey(m, s))

	prices, changes = getPricesAndChanges(2024, 2000)
	m = monkey{2024, prices, changes}
	assert.Equal(t, 9, buyBananasFromMonkey(m, s))
}

func Test_buyBananas(t *testing.T) {
	monkeys := parseMonkeys(testInput, 2000)
	assert.Equal(t, 23, buyBananas(monkeys, seq{-2, 1, -1, 3}))
}
