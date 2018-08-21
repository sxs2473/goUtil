package sliceUtil

func MergeSortedSlice(s1, s2 []int64) []int64 {
	// 从末尾元素开始遍历
	i := len(s1) - 1
	j := len(s2) - 1
	// 合并后的长度
	newLen := len(s1) + len(s2)
	// 合并后的索引，也从末尾元素开始
	newIdx := newLen - 1
	// 创建一个新切片，代表合并后的
	newS := make([]int64, newLen)
	// 将 s1 的内容拷贝到新切片
	for k, v := range s1 {
		newS[k] = v
	}
	// 开始遍历
	for i >= 0 && j >= 0 {
		// 新元素
		var newNum int64
		// 将较大的值赋给新元素，同时向前移动指针
		if newS[i] > s2[j] {
			newNum = newS[i]
			i--
		} else {
			newNum = s2[j]
			j--
		}
		newS[newIdx] = newNum
		newIdx--
	}
	// 如果 s2 还有剩余元素，则剩余元素一定都是最小的，直接放到头部即可
	for j >= 0 {
		newS[newIdx] = s2[j]
		j--
		newIdx--
	}
	return newS
}

func FilterSlice(s []int64, filter func(x int64) bool) []int64 {
	// 返回的新切片
	// s[:0] 这种写法是创建了一个 len 为 0，cap 为 len(s) 即和原始切片最大容量一致的切片
	// 因为是过滤，所以新切片的元素总个数一定不大于比原始切片，这样做减少了切片扩容带来的影响
	newS := s[:0]
	// 遍历，对每个元素执行 filter，符合条件的加入新切片中
	for _, x := range s {
		if !filter(x) {
			newS = append(newS, x)
		}
	}
	return newS
}
