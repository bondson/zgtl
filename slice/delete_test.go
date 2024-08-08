/**
  @Description: 差集
  @Author: ZPS
**/

package slice

import (
	"fmt"
	"testing"
)

func TestDelete(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 12, 13}
	fmt.Printf("len: %+v, len: %d, cap: %d \n", list, len(list), cap(list))
	list = Delete[int](list, 8)
	fmt.Printf("len: %+v, len: %d, cap: %d \n", list, len(list), cap(list))
}
