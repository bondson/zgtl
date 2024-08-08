/*
 * @Description: 结构体转Map
 * @Author: ZPS
 */

package structx

import (
	"fmt"
	"testing"
)

func TestToGormMap(t *testing.T) {
	type MyStruct struct {
		Id    string `gorm:"column:id;type:varchar(9);comment:ID" json:"id"`
		Name  string `gorm:"column:name;"`
		Value string `json:"value"`
		age   int
	}
	myStruct := MyStruct{
		Id:    "123",
		Name:  "Name1",
		Value: "Value2",
		age:   9,
	}
	result := ToMap(myStruct)
	fmt.Println(result)
	result = ToGormMap(myStruct, "id")
	fmt.Println(result)
	result = ToGormMap(myStruct, "id", "name")
	fmt.Println(result)
}
