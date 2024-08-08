/**
  @Description: 删除切片的某个元素
  @Author: ZPS
**/

package slice

import "github.com/bondson/zgtl/internal/slice"

// Delete 删除 index 处的元素
// 如果下标不是合法的下标，返回原数据
func Delete[Src any](src []Src, index int) []Src {
	res, _, err := slice.Delete[Src](src, index)
	if err != nil {
		return src
	}
	return res
}

// FilterDelete 删除符合条件的元素
// 考虑到性能问题，所有操作都会在原切片上进行
// 被删除元素之后的元素会往前移动，有且只会移动一次
func FilterDelete[T any](list []T, f func(key int, value T) bool) []T {
	// 记录被删除的元素位置，也称空缺的位置
	pos := 0
	for key := range list {
		// 判断是否满足删除的条件
		if f(key, list[key]) {
			continue
		}
		// 移动元素
		list[pos] = list[key]
		pos++
	}
	return list[:pos]
}

// Shrink 缩容
func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	n, changed := calCapacity(c, l)
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}

func calCapacity(c, l int) (int, bool) {
	// 容量 <=64 缩不缩都无所谓，因为浪费内存也浪费不了多少
	// 你可以考虑调大这个阈值，或者调小这个阈值
	if c <= 64 {
		return c, false
	}
	// 如果容量大于 2048，但是元素不足一半，
	// 降低为 0.625，也就是 5/8
	// 也就是比一半多一点，和正向扩容的 1.25 倍相呼应
	if c > 2048 && (c/l >= 2) {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}
	// 如果在 2048 以内，并且元素不足 1/4，那么直接缩减为一半
	if c <= 2048 && (c/l >= 4) {
		return c / 2, true
	}
	// 整个实现的核心是希望在后续少触发扩容的前提下，一次性释放尽可能多的内存
	return c, false
}
