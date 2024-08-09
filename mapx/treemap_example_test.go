/**
  @Description:
  @Author: ZPS
**/

package mapx_test

import (
	"fmt"

	"github.com/bondson/zgtl"
	"github.com/bondson/zgtl/mapx"
)

func ExampleNewTreeMap() {
	m, _ := mapx.NewTreeMap[int, int](zgtl.ComparatorRealNumber[int])
	_ = m.Put(1, 11)
	val, _ := m.Get(1)
	fmt.Println(val)
	// Output:
	// 11
}
