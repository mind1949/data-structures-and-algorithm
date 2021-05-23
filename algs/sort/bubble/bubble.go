package bubble

func Sort(s []int) {
	if len(s) < 2 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			s[i], s[i+1] = s[i+1], s[i]
		}
	}
	Sort(s[:len(s)-1])
}
