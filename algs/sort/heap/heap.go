package heap

// Sort 堆排序
func Sort(s []int) {
	if len(s) <= 1 {
		return
	}

	buildMaxHeap(s)
	for l := len(s) - 1; l > 0; l-- {
		s[0], s[l] = s[l], s[0]
		maxHeapify(s[0:l])
	}
}

// buildMaxHeap 构建最大堆
func buildMaxHeap(s []int) {
	if len(s) <= 1 {
		return
	}
	for i := (len(s) - 1) / 2; i >= 0; i-- {
		maxHeapify(s[i:])
	}
}

// maxHeapify 维护最大堆性质
func maxHeapify(s []int) {
	if len(s) <= 1 {
		return
	}

	parent := 0
	for {
		child := 2*parent + 1
		if child >= len(s) {
			return
		}
		if child+1 < len(s) && s[child] < s[child+1] {
			child++
		}
		if s[parent] >= s[child] {
			return
		}
		s[child], s[parent] = s[parent], s[child]
		parent = child
	}
}
