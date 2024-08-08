/*
 * @Description: 解析 tag
 * @Author: ZPS
 */

package structx

import (
	"reflect"
	"regexp"
	"strings"
)

func ForValue(obj any, fn func(field reflect.StructField, value any)) {
	reflectData := reflect.ValueOf(obj)
	for reflectData.Kind() == reflect.Ptr {
		reflectData = reflectData.Elem()
	}
	for i := 0; i < reflectData.NumField(); i++ {
		field := reflectData.Type().Field(i)
		if !field.IsExported() {
			// 跳过私有字段， unexported
			continue
		}
		//reflect.Zero(field.Type).Interface() //获取类型的零值
		value := reflectData.Field(i).Interface()
		fn(field, value)
	}
}

// ParseTag 解析 tag 中的键值对
func ParseTag(tag, key string) string {
	parts := regexp.MustCompile(`\s*;\s*`).Split(tag, -1)
	for _, part := range parts {
		kv := strings.SplitN(part, ":", 2)
		if len(kv) == 2 && kv[0] == key {
			return kv[1]
		}
	}
	return ""
}
