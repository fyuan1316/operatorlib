package model

import (
	"context"
	"fmt"
	"time"
)

func loopUntil(ctx context.Context, interval time.Duration, maxRetries int, f func(oCtx *OperatorContext) (bool, error), oCtx *OperatorContext) error {
	count := 0
	for {
		if stop, err := f(oCtx); err != nil {
			if count++; count > maxRetries {
				return err
			}
		} else if stop {
			break
		}
		select {
		case <-time.After(interval):
		case <-ctx.Done():
			return fmt.Errorf("execute canceled")
		}
	}
	return nil
}
