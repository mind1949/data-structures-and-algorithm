package sort

import "testing"

func TestSortFn(t *testing.T, sortFn func(s []int)) {
	for i, s := range getCases() {
		before := make([]int, len(s))
		copy(before, s)

		sortFn(s)
		if !isSorted(s) {
			t.Fatalf("expect cases[%d] is sorted, but got: %v ---> %v", i, before, s)
		} else {
			t.Logf("%v ---> %v", before, s)
		}
	}
}

func getCases() [][]int {
	return [][]int{
		{1},
		{2, 0},
		{3, 3, 4, 2, 1, 10, 5},
		{4, 3, 2, 1, 0, 10, 8, 21, 6},
		{1, 9, 4, 9, 10, 1},
		{100, 99, 30, 21, 5, 3, 10000, 9898, 7, 3, 3, 4, 2, 1, 10, 5},
	}
}

func isSorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
