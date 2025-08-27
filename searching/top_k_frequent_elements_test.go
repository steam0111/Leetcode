package searching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Idea :
// - using map for count numbres
// - using backet sort for get frequent elements
func TopKFrequent(nums []int, k int) []int {
	if k <= 0 || len(nums) == 0 {
		return []int{}
	}

	numsCount := make(map[int]int)

	for _, val := range nums {
		numsCount[val]++
	}

	// in case [7, 7, 7] where len = 3 and bucket index = 3
	buckets := make([][]int, len(nums)+1)

	for key, val := range numsCount {
		buckets[val] = append(buckets[val], key)
	}

	res := make([]int, 0, k)

	// skip 0 bucket
	for i := len(buckets) - 1; i > 0; i-- {
		for _, val := range buckets[i] {
			res = append(res, val)

			if len(res) == k {
				return res
			}
		}
	}

	return res
}

func TestTopKFrequent(t *testing.T) {
	t.Run("base test", func(t *testing.T) {
		got := TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
		assert.Equal(t, []int{1, 2}, got)
	})

	t.Run("equals count", func(t *testing.T) {
		got := TopKFrequent([]int{500, 500, 500, 600, 600, 600}, 2)
		assert.Equal(t, []int{500, 600}, got)
	})

	t.Run("single negative", func(t *testing.T) {
		got := TopKFrequent([]int{-1, -1}, 1)
		assert.Equal(t, []int{-1}, got)
	})
}
