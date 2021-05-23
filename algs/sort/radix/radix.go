package radix

// Sort 基数排序，只能排序正数
func Sort(s []int) {
	if len(s) <= 1 {
		return
	}

	max := s[0]
	min := s[0]
	for i := 1; i < len(s); i++ {
		if max < s[i] {
			max = s[i]
		}
		if min > s[i] {
			min = s[i]
		}
	}

	for bit := 1; max/bit > 0; bit *= 10 {
		bitSort(s, bit)
	}

	return
}

func bitSort(s []int, bit int) {
	bitCounts := make([]int, 10)

	for i := 0; i < len(s); i++ {
		num := (s[i] / bit) % 10
		bitCounts[num]++
	}

	for j := 1; j < len(bitCounts); j++ {
		bitCounts[j] += bitCounts[j-1]
	}

	tmp := make([]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		num := (s[i] / bit) % 10
		tmp[bitCounts[num]-1] = s[i]
		bitCounts[num]--
	}

	copy(s, tmp)
}
