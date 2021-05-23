package bucket

import (
	"algorithm/algs/sort/quick"
	"math"
)

func Sort(s []int) {
	if len(s) <= 1 {
		return
	}

	var (
		max = s[0]
		min = s[0]
	)
	for _, v := range s {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}

	bucketCap := 2                                                         // 桶的容量
	bucketCount := int(math.Ceil(float64(max-min+1) / float64(bucketCap))) // 桶的数量
	buckets := make([][]int, bucketCount)
	for _, v := range s {
		i := (v - min) / bucketCap
		buckets[i] = append(buckets[i], v)
	}
	for _, bucket := range buckets {
		quick.Sort(bucket)
	}
	s = s[:0]
	for _, bucket := range buckets {
		s = append(s, bucket...)
	}
}
