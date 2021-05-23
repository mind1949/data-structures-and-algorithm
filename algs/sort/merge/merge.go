package merge

func Sort(s []int) {
	if len(s) <= 1 {
		return
	}
	// 分解
	m := len(s) / 2
	l := s[:m]
	r := s[m:]
	// 解决
	Sort(l)
	Sort(r)
	// 合并
	merge(l, r)
}

func merge(l, r []int) {
	tmp := make([]int, 0, len(l)+len(r))
	var (
		li, ri int
	)
	for {
		if l[li] < r[ri] {
			tmp = append(tmp, l[li])
			li++
		} else {
			tmp = append(tmp, r[ri])
			ri++
		}
		if li >= len(l) {
			tmp = append(tmp, r[ri:]...)
			break
		}
		if ri >= len(r) {
			tmp = append(tmp, l[li:]...)
			break
		}
	}
	// 每次合并时，将原始切片覆盖，这样返回时会自动回收临时切片tmp
	copy(l[:], tmp[:len(l)])
	copy(r[:], tmp[len(l):])
}
