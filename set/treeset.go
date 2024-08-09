/**
  @Description:
  @Author: ZPS
**/

package set

import (
	"github.com/bondson/zgtl"
	"github.com/bondson/zgtl/mapx"
)

type TreeSet[T any] struct {
	treeMap *mapx.TreeMap[T, any]
}

func NewTreeSet[T any](compare zgtl.Comparator[T]) (*TreeSet[T], error) {
	treeMap, err := mapx.NewTreeMap[T, any](compare)
	if err != nil {
		return nil, err
	}
	return &TreeSet[T]{
		treeMap: treeMap,
	}, nil
}

func (s *TreeSet[T]) Add(key T) {
	_ = s.treeMap.Put(key, nil)
}

func (s *TreeSet[T]) Delete(key T) {
	s.treeMap.Delete(key)
}

func (s *TreeSet[T]) Exist(key T) bool {
	_, isExist := s.treeMap.Get(key)
	return isExist
}

// Keys 方法返回的元素顺序不固定
func (s *TreeSet[T]) Keys() []T {
	return s.treeMap.Keys()
}

var _ Set[int] = (*TreeSet[int])(nil)
