package shell

import (
	//pkgerrors "github.com/pkg/errors"
	"errors"
)

const InternalIgnoreShellScript = "IgnoreShellScript"

var InternalIgnoreShellScriptError = errors.New(InternalIgnoreShellScript)
