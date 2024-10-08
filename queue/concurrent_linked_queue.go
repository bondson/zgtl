/*
 * @Description: TODO
 * @Author: ZPS
 */

package queue

import (
	"sync/atomic"
	"unsafe"
)

// ConcurrentLinkedQueue 无界并发安全队列
type ConcurrentLinkedQueue[T any] struct {
	// *node[T]
	head unsafe.Pointer
	// *node[T]
	tail unsafe.Pointer
}

func NewConcurrentLinkedQueue[T any]() *ConcurrentLinkedQueue[T] {
	head := &node[T]{}
	ptr := unsafe.Pointer(head)
	return &ConcurrentLinkedQueue[T]{
		head: ptr,
		tail: ptr,
	}
}

func (c *ConcurrentLinkedQueue[T]) Enqueue(t T) error {
	newNode := &node[T]{val: t}
	newPtr := unsafe.Pointer(newNode)
	for {
		tailPtr := atomic.LoadPointer(&c.tail)
		tail := (*node[T])(tailPtr)
		tailNext := atomic.LoadPointer(&tail.next)
		if tailNext != nil {
			// 已经被人修改了，我们不需要修复，因为预期中修改的那个人会把 c.tail 指过去
			continue
		}
		if atomic.CompareAndSwapPointer(&tail.next, tailNext, newPtr) {
			// 如果失败也不用担心，说明有人抢先一步了
			atomic.CompareAndSwapPointer(&c.tail, tailPtr, newPtr)
			return nil
		}
	}
}

func (c *ConcurrentLinkedQueue[T]) Dequeue() (T, error) {
	for {
		headPtr := atomic.LoadPointer(&c.head)
		head := (*node[T])(headPtr)
		tailPtr := atomic.LoadPointer(&c.tail)
		tail := (*node[T])(tailPtr)
		if head == tail {
			// 不需要做更多检测，在当下这一刻，我们就认为没有元素，即便这时候正好有人入队
			// 但是并不妨碍我们在它彻底入队完成——即所有的指针都调整好——之前，
			// 认为其实还是没有元素
			var t T
			return t, ErrEmptyQueue
		}
		headNextPtr := atomic.LoadPointer(&head.next)
		if atomic.CompareAndSwapPointer(&c.head, headPtr, headNextPtr) {
			headNext := (*node[T])(headNextPtr)
			return headNext.val, nil
		}
	}
}

type node[T any] struct {
	val T
	// *node[T]
	next unsafe.Pointer
}
