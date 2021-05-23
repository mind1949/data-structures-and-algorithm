package radix

import (
	"algorithm/algs/sort"
	"testing"
)

func TestSort(t *testing.T) {
	for i, s := range [][]int{
		{1},
		{2, 0},
		{3, 3, 4, 2, 1, 10, 5},
		{4, 3, 2, 1, 0, 10, 8, 21, 6},
		{1, 9, 4, 9, 10, 1},
		{100, 99, 30, 21, 5, 3, 10000, 9898, 7, 3, 3, 4, 2, 1, 10, 5},
	} {
		before := make([]int, len(s))
		copy(before, s)

		Sort(s)
		if !sort.IsSorted(s) {
			t.Fatalf("expect cases[%d] is sorted, but got: %v ---> %v", i, before, s)
		} else {
			t.Logf("%v ---> %v", before, s)
		}
	}
}
