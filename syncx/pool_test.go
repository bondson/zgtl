/*
 * @Description: TODO
 * @Author: ZPS
 */

package syncx

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPool(t *testing.T) {
	cnt := 0
	p := NewPool[[]byte](func() []byte {
		cnt += 1
		res := make([]byte, 1, 12)
		res[0] = 'A'
		return res
	})

	res := p.Get()
	assert.Equal(t, "A", string(res))
	res = append(res, 'B')
	p.Put(res)
	res = p.Get()
	if cnt == 1 {
		assert.Equal(t, "AB", string(res))
	} else {
		assert.Equal(t, "A", string(res))
	}

}

func ExampleNew() {
	p := NewPool[[]byte](func() []byte {
		res := make([]byte, 1, 12)
		res[0] = 'A'
		return res
	})

	res := p.Get()
	fmt.Print(string(res))
	// Output:
	// A
}

func BenchmarkPool_Get(b *testing.B) {
	p := NewPool[string](func() string {
		return ""
	})

	sp := &sync.Pool{
		New: func() any {
			return ""
		},
	}
	b.Run("Pool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			p.Get()
		}
	})
	b.Run("sync.Pool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sp.Get()
		}
	})
}
