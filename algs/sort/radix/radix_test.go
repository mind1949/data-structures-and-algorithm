package radix

import (
	"algorithm/algs/sort"
	"testing"
)

func TestSort(t *testing.T) {
	sort.TestSortFn(t, Sort)
}
