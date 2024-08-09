/**
  @Description:
  @Author: ZPS
**/

package mapx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	testCases := []struct {
		name    string
		input   map[int]int
		wantRes []int
	}{
		{
			name:    "nil",
			input:   nil,
			wantRes: []int{},
		},
		{
			name:    "empty",
			input:   map[int]int{},
			wantRes: []int{},
		},
		{
			name: "single",
			input: map[int]int{
				1: 11,
			},
			wantRes: []int{1},
		},
		{
			name: "multiple",
			input: map[int]int{
				1: 11,
				2: 12,
			},
			wantRes: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Keys[int, int](tc.input)
			assert.ElementsMatch(t, tc.wantRes, res)
		})
	}
}
func TestValues(t *testing.T) {
	testCases := []struct {
		name    string
		input   map[int]int
		wantRes []int
	}{
		{
			name:    "nil",
			input:   nil,
			wantRes: []int{},
		},
		{
			name:    "empty",
			input:   map[int]int{},
			wantRes: []int{},
		},
		{
			name: "single",
			input: map[int]int{
				1: 11,
			},
			wantRes: []int{11},
		},
		{
			name: "multiple",
			input: map[int]int{
				1: 11,
				2: 12,
			},
			wantRes: []int{11, 12},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Values[int, int](tc.input)
			assert.ElementsMatch(t, tc.wantRes, res)
		})
	}
}

func TestKeysValues(t *testing.T) {
	testCases := []struct {
		name       string
		input      map[int]int
		wantKeys   []int
		wantValues []int
	}{
		{
			name:       "nil",
			input:      nil,
			wantKeys:   []int{},
			wantValues: []int{},
		},
		{
			name:       "empty",
			input:      map[int]int{},
			wantKeys:   []int{},
			wantValues: []int{},
		},
		{
			name: "single",
			input: map[int]int{
				1: 11,
			},
			wantKeys:   []int{1},
			wantValues: []int{11},
		},
		{
			name: "multiple",
			input: map[int]int{
				1: 11,
				2: 12,
			},
			wantKeys:   []int{1, 2},
			wantValues: []int{11, 12},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			keys, values := KeysValues[int, int](tc.input)
			assert.ElementsMatch(t, tc.wantKeys, keys)
			assert.ElementsMatch(t, tc.wantValues, values)
		})
	}
}
