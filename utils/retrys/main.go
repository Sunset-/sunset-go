package retrys

import (
	"context"
	"time"
)

//重试
func Retry(f func() error, times int, space time.Duration, ctx context.Context) error {
	if times <= 1 {
		return f()
	}
	var err error
	var i = 0
	for {
		if ctx != nil {
			select {
			case <-ctx.Done():
				return err
			default:
			}
		}
		err = f()
		if err == nil {
			return nil
		}
		i++
		if i >= times {
			break
		}
		time.Sleep(space)
	}
	return err
}
