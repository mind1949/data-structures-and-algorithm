package quick

func Sort(s []int) {
	if len(s) <= 1 {
		return
	}
	m := partition(s)
	Sort(s[:m])
	Sort(s[m:])
}

// partition 分成左右两个区：s[:m]、s[m:]
// 左边的元素小于等于右边的任意一个元素
//
// 保证左边小于右边的思路是，从左边发现大于中间元素的，就替换到右边
// 右边发现小于中间元素的就替换到左边，最终左右两个索引指针会相遇，这个相遇的索引就是中间元素的位置
func partition(s []int) (m int) {
	p := s[len(s)-1]
	var (
		i = 0
		j = len(s) - 1
	)
	for i < j && j > 0 {
		if s[i] <= p {
			i++
			continue
		}
		s[j] = s[i]
		j--
		if s[j] >= p {
			j--
			continue
		}
		s[i] = s[j]
		i++
	}
	m = i
	s[m] = p
	return m
}
