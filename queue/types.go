/*
 * @Description: TODO
 * @Author: ZPS
 */

package queue

import (
	"context"
	"errors"
)

var ErrEmptyQueue = errors.New("zgtl: 队列为空")

// BlockingQueue 阻塞队列
// 参考 Queue 普通队列
// 一个阻塞队列是否遵循 FIFO 取决于具体实现
type BlockingQueue[T any] interface {
	// Enqueue 将元素放入队列。如果在 ctx 超时之前，队列有空闲位置，那么元素会被放入队列；
	// 否则返回 error。
	// 在超时或者调用者主动 cancel 的情况下，所有的实现都必须返回 ctx。
	// 调用者可以通过检查 error 是否为 context.DeadlineExceeded
	// 或者 context.Canceled 来判断入队失败的原因
	// 注意，调用者必须使用 errors.Is 来判断，而不能直接使用 ==
	Enqueue(ctx context.Context, t T) error
	// Dequeue 从队首获得一个元素
	// 如果在 ctx 超时之前，队列中有元素，那么会返回队首的元素，否则返回 error。
	// 在超时或者调用者主动 cancel 的情况下，所有的实现都必须返回 ctx。
	// 调用者可以通过检查 error 是否为 context.DeadlineExceeded
	// 或者 context.Canceled 来判断入队失败的原因
	// 注意，调用者必须使用 errors.Is 来判断，而不能直接使用 ==
	Dequeue(ctx context.Context) (T, error)
}

// Queue 普通队列
// 参考 BlockingQueue 阻塞队列
// 一个队列是否遵循 FIFO 取决于具体实现
type Queue[T any] interface {
	// Enqueue 将元素放入队列，如果此时队列已经满了，那么返回错误
	Enqueue(t T) error
	// Dequeue 从队首获得一个元素
	// 如果此时队列里面没有元素，那么返回错误
	Dequeue() (T, error)
}
