/*
 * @Description: TODO
 * @Author: ZPS
 */

package structx

import (
	"fmt"
	"testing"
)

func TestTrimSpace(t *testing.T) {
	type Person struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}
	type args struct {
		obj Person
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Alice",
			args: args{obj: Person{
				Name: " Alice ", Email: "alice@example.com",
			}},
		},
		{
			name: "Bob",
			args: args{obj: Person{
				Name: " Bob   ", Email: " bob@example.com ",
			}},
		},
		{
			name: "Charlie",
			args: args{obj: Person{
				Name: "Charlie", Email: "  charlie@example.com ",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("tt.args.obj: %+v \n", tt.args.obj)
			TrimSpace(&tt.args.obj)
			fmt.Printf("tt.args.obj: %+v \n", tt.args.obj)
			//fmt.Printf("data: %+v \n", data)
			fmt.Print("*********")
		})
	}
}
