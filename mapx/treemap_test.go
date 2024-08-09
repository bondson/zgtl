/**
  @Description:
  @Author: ZPS
**/

package mapx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/bondson/zgtl"
)

func TestNewTreeMapWithMap(t *testing.T) {
	tests := []struct {
		name       string
		m          map[int]int
		comparable zgtl.Comparator[int]
		wantKey    []int
		wantVal    []int
		wantErr    error
	}{
		{
			name:       "nil",
			m:          nil,
			comparable: nil,
			wantKey:    nil,
			wantVal:    nil,
			wantErr:    errors.New("zgtl: Comparator不能为nil"),
		},
		{
			name:       "empty",
			m:          map[int]int{},
			comparable: compare(),
			wantKey:    nil,
			wantVal:    nil,
			wantErr:    nil,
		},
		{
			name: "single",
			m: map[int]int{
				0: 0,
			},
			comparable: compare(),
			wantKey:    []int{0},
			wantVal:    []int{0},
			wantErr:    nil,
		},
		{
			name: "multiple",
			m: map[int]int{
				0: 0,
				1: 1,
				2: 2,
			},
			comparable: compare(),
			wantKey:    []int{0, 1, 2},
			wantVal:    []int{0, 1, 2},
			wantErr:    nil,
		},
		{
			name: "disorder",
			m: map[int]int{
				1: 1,
				2: 2,
				0: 0,
				3: 3,
				5: 5,
				4: 4,
			},
			comparable: compare(),
			wantKey:    []int{0, 1, 2, 3, 5, 4},
			wantVal:    []int{0, 1, 2, 3, 5, 4},
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			treeMap, err := NewTreeMapWithMap[int, int](tt.comparable, tt.m)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			for k, v := range tt.m {
				value, _ := treeMap.Get(k)
				assert.Equal(t, true, v == value)
			}

		})

	}
}

func TestTreeMap_Get(t *testing.T) {
	var tests = []struct {
		name     string
		m        map[int]int
		findKey  int
		wantVal  int
		wantBool bool
	}{
		{
			name:     "empty-TreeMap",
			m:        map[int]int{},
			findKey:  0,
			wantVal:  0,
			wantBool: false,
		},
		{
			name: "find",
			m: map[int]int{
				1: 1,
				2: 2,
				0: 0,
				3: 3,
				5: 5,
				4: 4,
			},
			findKey:  2,
			wantVal:  2,
			wantBool: true,
		},
		{
			name: "not-find",
			m: map[int]int{
				1: 1,
				2: 2,
				0: 0,
				3: 3,
				5: 5,
				4: 4,
			},
			findKey:  6,
			wantVal:  0,
			wantBool: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			treeMap, _ := NewTreeMap[int, int](compare())
			putAll(treeMap, tt.m)
			val, b := treeMap.Get(tt.findKey)
			assert.Equal(t, tt.wantBool, b)
			assert.Equal(t, tt.wantVal, val)
		})
	}
}

func TestTreeMap_Put(t *testing.T) {

	tests := []struct {
		name    string
		k       []int
		v       []string
		wantKey []int
		wantVal []string
		wantErr error
	}{
		{
			name:    "single",
			k:       []int{0},
			v:       []string{"0"},
			wantKey: []int{0},
			wantVal: []string{"0"},
			wantErr: nil,
		},
		{
			name:    "multiple",
			k:       []int{0, 1, 2},
			v:       []string{"0", "1", "2"},
			wantKey: []int{0, 1, 2},
			wantVal: []string{"0", "1", "2"},
			wantErr: nil,
		},
		{
			name:    "same",
			k:       []int{0, 0, 1, 2, 2, 3},
			v:       []string{"0", "999", "1", "998", "2", "3"},
			wantKey: []int{0, 1, 2, 3},
			wantVal: []string{"999", "1", "2", "3"},
			wantErr: nil,
		},
		{
			name:    "same",
			k:       []int{0, 0},
			v:       []string{"0", "999"},
			wantKey: []int{0},
			wantVal: []string{"999"},
			wantErr: nil,
		},
		{
			name:    "disorder",
			k:       []int{1, 2, 0, 3, 5, 4},
			v:       []string{"1", "2", "0", "3", "5", "4"},
			wantKey: []int{0, 1, 2, 3, 4, 5},
			wantVal: []string{"0", "1", "2", "3", "4", "5"},
			wantErr: nil,
		},
		{
			name:    "disorder-same",
			k:       []int{1, 3, 2, 0, 2, 3},
			v:       []string{"1", "2", "998", "0", "3", "997"},
			wantKey: []int{0, 1, 2, 3},
			wantVal: []string{"0", "1", "3", "997"},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			treeMap, _ := NewTreeMap[int, string](compare())
			for i := 0; i < len(tt.k); i++ {
				err := treeMap.Put(tt.k[i], tt.v[i])
				if err != nil {
					assert.Equal(t, tt.wantErr, err)
					return
				}
			}
			for i := 0; i < len(tt.wantKey); i++ {
				v, b := treeMap.Get(tt.wantKey[i])
				assert.Equal(t, true, b)
				assert.Equal(t, tt.wantVal[i], v)
			}

		})
	}
	subTests := []struct {
		name    string
		k       []int
		v       []string
		wantKey []int
		wantVal []string
		wantErr error
	}{
		{
			name:    "nil",
			k:       []int{0},
			v:       nil,
			wantKey: []int{0},
			wantVal: []string(nil),
		},
		{
			name:    "nil",
			k:       []int{0},
			v:       []string{"0"},
			wantKey: []int{0},
			wantVal: []string{"0"},
		},
	}
	for _, tt := range subTests {
		t.Run(tt.name, func(t *testing.T) {
			treeMap, _ := NewTreeMap[int, []string](compare())
			for i := 0; i < len(tt.k); i++ {
				err := treeMap.Put(tt.k[i], tt.v)
				if err != nil {
					assert.Equal(t, tt.wantErr, err)
					return
				}
			}
			for i := 0; i < len(tt.wantKey); i++ {
				v, b := treeMap.Get(tt.wantKey[i])
				assert.Equal(t, true, b)
				assert.Equal(t, tt.wantVal, v)
			}

		})
	}
}

func TestTreeMap_Keys(t *testing.T) {
	testCases := []struct {
		name     string
		data     map[int]int
		wantKeys []int
	}{
		{
			name:     "no data",
			wantKeys: []int{},
		},
		{
			name: "data",
			data: map[int]int{
				1: 11,
				2: 12,
				0: 10,
				3: 13,
				5: 15,
				4: 14,
			},
			wantKeys: []int{0, 1, 2, 3, 4, 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tm, err := NewTreeMap[int, int](compare())
			require.NoError(t, err)
			for k, v := range tc.data {
				err = tm.Put(k, v)
				require.NoError(t, err)
			}
			keys := tm.Keys()
			assert.Equal(t, tc.wantKeys, keys)
		})
	}
}

func TestTreeMap_Values(t *testing.T) {
	testCases := []struct {
		name       string
		data       map[int]int
		wantValues []int
	}{
		{
			name:       "no data",
			wantValues: []int{},
		},
		{
			name: "data",
			data: map[int]int{
				1: 11,
				2: 12,
				0: 10,
				3: 13,
				5: 15,
				4: 14,
			},
			wantValues: []int{10, 11, 12, 13, 14, 15},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tm, err := NewTreeMap[int, int](compare())
			require.NoError(t, err)
			for k, v := range tc.data {
				err = tm.Put(k, v)
				require.NoError(t, err)
			}
			vals := tm.Values()
			assert.Equal(t, tc.wantValues, vals)
		})
	}
}

func TestTreeMap_Delete(t *testing.T) {
	var tests = []struct {
		name    string
		m       map[int]int
		delKey  int
		delVal  int
		deleted bool
	}{
		{
			name:   "empty-TreeMap",
			m:      map[int]int{},
			delKey: 0,
		},
		{
			name: "find",
			m: map[int]int{
				1: 1,
				2: 2,
				0: 0,
				3: 3,
				5: 5,
				4: 4,
			},
			delKey:  2,
			deleted: true,
			delVal:  2,
		},
		{
			name: "not-find",
			m: map[int]int{
				1: 1,
				2: 2,
				0: 0,
				3: 3,
				5: 5,
				4: 4,
			},
			delKey: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			treeMap, _ := NewTreeMap[int, int](compare())
			for k, v := range tt.m {
				err := treeMap.Put(k, v)
				require.NoError(t, err)
			}
			delVal, ok := treeMap.Delete(tt.delKey)
			assert.Equal(t, tt.deleted, ok)
			assert.Equal(t, tt.delVal, delVal)
			_, ok = treeMap.Get(tt.delKey)
			assert.False(t, ok)
		})
	}
}

func compare() zgtl.Comparator[int] {
	return zgtl.ComparatorRealNumber[int]
}

func BenchmarkTreeMap(b *testing.B) {
	hashmap := NewHashMap[hashInt, int](10)
	treeMap, _ := NewTreeMap[uint64, int](zgtl.ComparatorRealNumber[uint64])
	m := make(map[uint64]int, 10)
	b.Run("treeMap_put", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = treeMap.Put(uint64(i), i)
		}
	})
	b.Run("map_put", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m[uint64(i)] = i
		}
	})
	b.Run("hashMap_put", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = hashmap.Put(hashInt(uint64(i)), i)
		}
	})
	b.Run("treeMap_get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = treeMap.Get(uint64(i))
		}
	})
	b.Run("map_get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = m[uint64(i)]
		}
	})
	b.Run("hashMap_get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = hashmap.Get(hashInt(uint64(i)))
		}
	})

}
