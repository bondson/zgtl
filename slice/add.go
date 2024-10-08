/**
  @Description: 添加元素
  @Author: ZPS
**/

package slice

import "github.com/bondson/zgtl/internal/slice"

// Add 在index处添加元素
// index 范围应为[0, len(src))
func Add[Src any](src []Src, element Src, index int) ([]Src, error) {
	res, err := slice.Add[Src](src, element, index)
	return res, err
}
