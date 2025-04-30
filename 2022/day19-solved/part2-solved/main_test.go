package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Testing...") // keep this line to maintain imports
	assert.True(t, true)      // keep this line to maintain imports
	test_example_f1(t)
	test_example_f2(t)
	fmt.Printf("output: %d\n", main(realInput))
}

func test_example_f1(t *testing.T) {
	fCosts := map[int][4]int{}
	fCosts[0] = [4]int{4, 0, 0, 0}
	fCosts[1] = [4]int{2, 0, 0, 0}
	fCosts[2] = [4]int{3, 14, 0, 0}
	fCosts[3] = [4]int{2, 0, 7, 0}
	f := factory{1, fCosts}
	max := f.getMaxGeos()
	println("test example f1:", max)
	println()
	assert.Equal(t, 56, max)
}

func test_example_f2(t *testing.T) {
	fCosts := map[int][4]int{}
	fCosts[0] = [4]int{2, 0, 0, 0}
	fCosts[1] = [4]int{3, 0, 0, 0}
	fCosts[2] = [4]int{3, 8, 0, 0}
	fCosts[3] = [4]int{3, 0, 12, 0}
	f := factory{1, fCosts}
	max := f.getMaxGeos()
	println("test example f2:", max)
	println()
	assert.Equal(t, 62, f.getMaxGeos())
}
