package errs

import (
	"fmt"
	"time"
)

// NewErrIndexOutOfRange 创建一个代表下标超出范围的错误
func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("下标超出范围，长度 %d, 下标 %d", length, index)
}

// NewErrInvalidType 创建一个代表类型转换失败的错误
func NewErrInvalidType(want, got string) error {
	return fmt.Errorf("gotool: 类型转换失败，want:%s, got:%s", want, got)
}

// NewErrInvalidIntervalValue 创建一个无效的间隔时间失败的错误
func NewErrInvalidIntervalValue(interval time.Duration) error {
	return fmt.Errorf("gotool: 无效的间隔时间 %d, 预期值应大于 0", interval)
}

// NewErrInvalidMaxIntervalValue 创建一个重试的间隔时间失败的错误
func NewErrInvalidMaxIntervalValue(maxInterval, initialInterval time.Duration) error {
	return fmt.Errorf("gotool: 最大重试间隔的时间 [%d] 应大于等于初始重试的间隔时间 [%d] ", maxInterval, initialInterval)
}
