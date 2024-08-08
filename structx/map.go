/*
 * @Description: 结构体转Map
 * @Author: ZPS
 */

package structx

import (
	"reflect"

	"github.com/bondson/zgtl/internal/structx"
	"github.com/bondson/zgtl/slice"
)

// ToMap 基于tag的json转Map, filters需要过滤的key
func ToMap(obj any, filters ...string) map[string]any {
	result := make(map[string]any)
	structx.ForValue(obj, func(field reflect.StructField, value any) {
		// 获取 json tag
		jsonTag := field.Tag.Get("json")
		// 如果有 json tag，则将 json tag 也加入结果 map
		if jsonTag != "" && !slice.Contains[string](filters, jsonTag) {
			result[jsonTag] = value
		}
	})
	return result
}

// ToGormMap 基于tag的gorm:column,json 转 Map， filters需要过滤的key
func ToGormMap(obj interface{}, filters ...string) map[string]any {
	result := make(map[string]any)
	structx.ForValue(obj, func(field reflect.StructField, value any) {
		// 获取 gorm 的 column 和 json tag
		gormTag := field.Tag.Get("gorm")
		// 解析 gormTag 获取 column 名称
		columnName := structx.ParseTag(gormTag, "column")
		// 如果 column 名称为空，则使用json tag
		if columnName == "" {
			columnName = field.Tag.Get("json")
			// 如果 columnName 名称为空，则使用字段名
			if columnName == "" {
				columnName = field.Name
			}
		}
		if slice.Contains[string](filters, columnName) {
			return
		}
		// 将字段名和对应的值加入结果 map
		result[columnName] = value
	})
	return result
}
