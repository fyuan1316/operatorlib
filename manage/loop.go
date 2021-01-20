package manage

import (
	"context"
	"fmt"
	"github.com/fyuan1316/operatorlib/manage/model"
	"github.com/fyuan1316/operatorlib/task/shell"
	pkgerrors "github.com/pkg/errors"
	"time"
)

var defualtLoopSetting = struct {
	Interval   time.Duration
	MaxRetries int
}{5, 3}

func loopUntil(ctx context.Context, interval time.Duration, maxRetries int, f func(oCtx *model.OperatorContext) (bool, error), oCtx *model.OperatorContext) error {
	count := 0
	for {
		// shell script不执行方法时，抛出InternalIgnoreShellScriptError
		if stop, err := f(oCtx); err != nil {
			if pkgerrors.Is(err, shell.InternalIgnoreShellScriptError) {
				return err
			}
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
