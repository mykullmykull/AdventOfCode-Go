package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Testing...") // keep this line to maintain imports
	assert.True(t, true)      // keep this line to maintain imports
	// assert.Equal(t, expected, main(testInput))
	// fmt.Printf("output: %d\n", main(realInput))
}

func Test_convert(t *testing.T) {
	assert.Equal(t, "1", convertToSnafu(1))
	assert.Equal(t, "2", convertToSnafu(2))
	assert.Equal(t, "1=", convertToSnafu(3))
	assert.Equal(t, "1-", convertToSnafu(4))
	assert.Equal(t, "10", convertToSnafu(5))
	assert.Equal(t, "11", convertToSnafu(6))
	assert.Equal(t, "12", convertToSnafu(7))
	assert.Equal(t, "2=", convertToSnafu(8))
	assert.Equal(t, "2-", convertToSnafu(9))
	assert.Equal(t, "20", convertToSnafu(10))
	assert.Equal(t, "1=0", convertToSnafu(15))
	assert.Equal(t, "1-0", convertToSnafu(20))
	assert.Equal(t, "1=11-2", convertToSnafu(2022))
	assert.Equal(t, "1-0---0", convertToSnafu(12345))
	assert.Equal(t, "1121-1110-1=0", convertToSnafu(314159265))
	assert.Equal(t, "1=-0-2", convertToSnafu(1747))
	assert.Equal(t, "906", convertToSnafu(12111))
	assert.Equal(t, "2=0=", convertToSnafu(198))
	assert.Equal(t, "21", convertToSnafu(11))
	assert.Equal(t, "2=01", convertToSnafu(201))
	assert.Equal(t, "111", convertToSnafu(31))
	assert.Equal(t, "20012", convertToSnafu(1257))
	assert.Equal(t, "112", convertToSnafu(32))
	assert.Equal(t, "1=-1=", convertToSnafu(353))
	assert.Equal(t, "1-12", convertToSnafu(107))
	assert.Equal(t, "12", convertToSnafu(7))
	assert.Equal(t, "1=", convertToSnafu(3))
	assert.Equal(t, "122", convertToSnafu(37))
}
