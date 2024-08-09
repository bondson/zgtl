/**
  @Description:
  @Author: ZPS
**/

package mapx

// Keys 返回 map 里面的所有的 key。
// 需要注意：这些 key 的顺序是随机。
func Keys[K comparable, V any](m map[K]V) []K {
	res := make([]K, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

// Values 返回 map 里面的所有的 value。
// 需要注意：这些 value 的顺序是随机。
func Values[K comparable, V any](m map[K]V) []V {
	res := make([]V, 0, len(m))
	for k := range m {
		res = append(res, m[k])
	}
	return res
}

// KeysValues 返回 map 里面的所有的 key,value。
// 需要注意：这些 (key,value) 的顺序是随机,相对顺序是一致的。
func KeysValues[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	values := make([]V, 0, len(m))
	for k := range m {
		keys = append(keys, k)
		values = append(values, m[k])
	}
	return keys, values
}
