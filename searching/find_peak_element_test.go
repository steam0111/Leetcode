package searching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Idea :
// - use binary search for log n
// - pick always be in one side bcs nums[-1] == nums[n] == - ∞
func FindPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func TestFindPeakElement(t *testing.T) {
	t.Run("base test", func(t *testing.T) {
		got := FindPeakElement([]int{10, 20, 50, 15, 10, 25})
		assert.Equal(t, 2, got)
	})
}
