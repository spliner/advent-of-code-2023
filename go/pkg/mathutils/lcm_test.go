package mathutils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeastCommonMultiple(t *testing.T) {
	type testCase struct {
		slice    []int
		expected int
	}
	testCases := []testCase{
		{[]int{6, 10}, 30},
		{[]int{10, 18, 25}, 450},
	}
	for _, c := range testCases {
		name := fmt.Sprintf("%v should return %d", c.slice, c.expected)
		t.Run(name, func(t *testing.T) {
			lcm := LeastCommonMultiple(c.slice)
			assert.Equal(t, c.expected, lcm)
		})
	}
}
