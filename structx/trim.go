/*
 * @Description: 删除空格
 * @Author: ZPS
 */

package structx

import (
	"reflect"
	"strings"

	"github.com/bondson/zgtl/internal/structx"
)

// TrimSpace 删除所有值的前导和尾随空格
func TrimSpace(obj any) {
	// 判断是否是指针(否不往下执行)
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return
	}
	structx.ForValue(obj, func(field reflect.StructField, value any) {
		if str, ok := value.(string); ok {
			str = strings.TrimSpace(str)
			reflect.ValueOf(obj).Elem().FieldByName(field.Name).SetString(str)
		}
	})
}
