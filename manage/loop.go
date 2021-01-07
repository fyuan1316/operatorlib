package manage

import (
	"context"
	"fmt"
	"github.com/fyuan1316/operatorlib/manage/model"
	"time"
)

func loopUntil(ctx context.Context, interval time.Duration, maxRetries int, f func(oCtx *model.OperatorContext) (bool, error), oCtx *model.OperatorContext) error {
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
