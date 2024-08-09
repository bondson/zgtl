/**
  @Description:
  @Author: ZPS
**/

package mapx

import (
	"errors"

	"github.com/bondson/zgtl"
	"github.com/bondson/zgtl/internal/tree"
)

var (
	errTreeMapComparatorIsNull = errors.New("zgtl: Comparator不能为nil")
)

// TreeMap 是基于红黑树实现的Map
type TreeMap[K any, V any] struct {
	tree *tree.RBTree[K, V]
}

// NewTreeMapWithMap TreeMap构造方法
// 支持通过传入的map构造生成TreeMap
func NewTreeMapWithMap[K comparable, V any](compare zgtl.Comparator[K], m map[K]V) (*TreeMap[K, V], error) {
	treeMap, err := NewTreeMap[K, V](compare)
	if err != nil {
		return treeMap, err
	}
	putAll(treeMap, m)
	return treeMap, nil
}

// NewTreeMap TreeMap构造方法,创建一个的TreeMap
// 需注意比较器compare不能为nil
func NewTreeMap[K any, V any](compare zgtl.Comparator[K]) (*TreeMap[K, V], error) {
	if compare == nil {
		return nil, errTreeMapComparatorIsNull
	}
	return &TreeMap[K, V]{
		tree: tree.NewRBTree[K, V](compare),
	}, nil
}

// putAll 将map传入TreeMap
// 需注意如果map中的key已存在,value将被替换
func putAll[K comparable, V any](treeMap *TreeMap[K, V], m map[K]V) {
	for k, v := range m {
		_ = treeMap.Put(k, v)
	}
}

// Put 在TreeMap插入指定值
// 需注意如果TreeMap已存在该Key那么原值会被替换
func (treeMap *TreeMap[K, V]) Put(key K, value V) error {
	err := treeMap.tree.Add(key, value)
	if err == tree.ErrRBTreeSameRBNode {
		return treeMap.tree.Set(key, value)
	}
	return nil
}

// Get 在TreeMap找到指定Key的节点,返回Val
// TreeMap未找到指定节点将会返回false
func (treeMap *TreeMap[K, V]) Get(key K) (V, bool) {
	v, err := treeMap.tree.Find(key)
	return v, err == nil
}

// Delete TreeMap中删除指定key的节点
func (treeMap *TreeMap[T, V]) Delete(k T) (V, bool) {
	return treeMap.tree.Delete(k)
}

// Keys 返回了全部的键
// 目前我们是按照中序遍历来返回的数据，但是你不能依赖于这个特性
func (treeMap *TreeMap[T, V]) Keys() []T {
	keys, _ := treeMap.tree.KeyValues()
	return keys
}

// Values 返回了全部的值
// 目前我们是按照中序遍历来返回的数据，但是你不能依赖于这个特性
func (treeMap *TreeMap[T, V]) Values() []V {
	_, vals := treeMap.tree.KeyValues()
	return vals
}

var _ mapi[any, any] = (*TreeMap[any, any])(nil)
