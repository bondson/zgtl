/**
  @Description: 随机码生成
  @Author: ZPS
**/

package randx

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRandCode(t *testing.T) {
	testCases := []struct {
		name      string
		length    int
		typ       TYPE
		wantMatch string
		wantErr   error
	}{
		{
			name:      "默认类型",
			length:    8,
			typ:       0,
			wantMatch: "^[0-9]+$",
			wantErr:   nil,
		},
		{
			name:      "数字验证码",
			length:    8,
			typ:       TYPE_DIGIT,
			wantMatch: "^[0-9]+$",
			wantErr:   nil,
		}, {
			name:      "小写字母验证码",
			length:    8,
			typ:       TYPE_LETTER,
			wantMatch: "^[a-z]+$",
			wantErr:   nil,
		}, {
			name:      "大写字母验证码",
			length:    8,
			typ:       TYPE_CAPITAL,
			wantMatch: "^[A-Z]+$",
			wantErr:   nil,
		}, {
			name:      "混合验证码",
			length:    8,
			typ:       TYPE_MIXED,
			wantMatch: "^[0-9a-zA-Z]+$",
			wantErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			code := RandCode(tc.length, tc.typ)
			fmt.Println("code:", code)
			//长度检验
			if len(code) != tc.length {
				t.Errorf("expected length: %d but got length:%d  ", tc.length, len(code))
			}
			//模式检验
			matched, _ := regexp.MatchString(tc.wantMatch, code)
			if !matched {
				t.Errorf("expected %s but got %s", tc.wantMatch, code)
			}
		})
	}

}
