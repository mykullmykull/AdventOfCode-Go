package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Testing...") // keep this line to maintain imports
	assert.True(t, true)      // keep this line to maintain imports
	test_mix(t)
	test_prune(t)
	test_evolve(t)
	assert.Equal(t, expected, main(testInput, 2000))
	fmt.Printf("output: %d\n", main(realInput, 2000))
}

func test_mix(t *testing.T) {
	assert.Equal(t, 37, mix(42, 15))
}

func test_prune(t *testing.T) {
	assert.Equal(t, 16113920, prune(100000000))
}

func test_evolve(t *testing.T) {
	expectations := []int{
		15887950,
		16495136,
		527345,
		704524,
		1553684,
		12683156,
		11100544,
		12249484,
		7753432,
		5908254,
	}
	for i, e := range expectations {
		assert.Equal(t, e, evolve(123, i+1))
	}
}
