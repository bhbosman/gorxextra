package gorxextra

import (
	"context"
	"github.com/bhbosman/goerrors"
	rxgo "github.com/reactivex/rxgo/v2"
	"time"
)

func SendContextWithTimeOutAndRetries(
	i rxgo.Item,
	ctx context.Context,
	duration time.Duration,
	retryCount int,
	ch chan rxgo.Item) error {

	newTimer := time.NewTimer(duration)
	defer newTimer.Stop()
	for r := 1; r <= retryCount; r++ {
		newTimer.Reset(duration)
		select {
		case <-newTimer.C:
			if r == retryCount {
				return goerrors.TimeOut
			}
			continue
		case <-ctx.Done():
			return ctx.Err()
		case ch <- i:
			return nil
		}
	}
	return goerrors.InvalidParam
}
