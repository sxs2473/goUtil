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
	// 同时，也有一个问题，因为 newS 和 s 共享底层数组，那么过滤后 s 也会被修改！
	newS := s[:0]
	// 遍历，对每个元素执行 filter，符合条件的加入新切片中
	for _, x := range s {
		if !filter(x) {
			newS = append(newS, x)
		}
	}
	return newS
}

func RemoveDuplicates(s []int64) []int64 {
	var ret []int64
	for _, v := range s {
		found := false
		for _, v2 := range ret {
			if v == v2 {
				found = true
				break
			}
		}
		if !found {
			ret = append(ret, v)
		}
	}
	return ret
}

func RemoveDuplicates2(s []int64) []int64 {
	ret := s[:0]
	// 利用 struct{}{} 减少内存占用
	assist := map[int64]struct{}{}
	for _, v := range s {
		if _, ok := assist[v]; !ok {
			assist[v] = struct{}{}
			ret = append(ret, v)
		}
	}
	return ret
}

func Reversing(s []int64) []int64 {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
	return s
}

func SliceChunk(s []int64, size int) [][]int64 {
	var ret [][]int64
	for size < len(s) {
		// s[:size:size] 表示 len 为 size，cap 也为 size，第二个冒号后的 size 表示 cap
		s, ret = s[size:], append(ret, s[:size:size])
	}
	ret = append(ret, s)
	return ret
}
