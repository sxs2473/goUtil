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
