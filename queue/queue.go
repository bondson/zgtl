/*
 * @Description: TODO
 * @Author: ZPS
 */

package queue

import (
	"container/list"
)

// 使用list实现queue
type ListQueue[T any] struct {
	data *list.List
}

func NewListQueue[T any]() *ListQueue[T] {
	return &ListQueue[T]{data: list.New()}
}

// 入队
func (q *ListQueue[T]) Enqueue(value T) {
	q.data.PushBack(value)
}

// 出队
func (q *ListQueue[T]) Dequeue() (T, error) {
	front := q.data.Front()
	var res T
	if front != nil {
		q.data.Remove(front)
		res = front.Value.(T)
		return res, nil
	}
	return res, ErrEmptyQueue
}
