package count

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

	cnt := make([]int, max-min+1)
	for _, v := range s {
		cnt[v-min]++
	}

	s = s[:0]
	for i := 0; i < len(cnt); i++ {
		c := cnt[i]
		for j := 0; j < c; j++ {
			s = append(s, i+min)
		}
	}
}
