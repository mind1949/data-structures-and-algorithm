package insertion

func Sort(s []int) {
	if len(s) < 2 {
		return
	}
	for i := 1; i < len(s); i++ {
		v := s[i]
		for j := 0; j < i; j++ {
			if v <= s[j] {
				copy(s[j+1:i+1], s[j:i])
				s[j] = v
				break
			}
		}
	}
}
