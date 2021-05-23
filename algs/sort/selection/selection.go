package selection

func Sort(s []int) {
	if len(s) < 2 {
		return
	}
	maxIdx := 0
	for i := 1; i < len(s); i++ {
		if s[maxIdx] < s[i] {
			maxIdx = i
		}
	}
	s[maxIdx], s[len(s)-1] = s[len(s)-1], s[maxIdx]
	Sort(s[:len(s)-1])
}
