package errors

import (
	stderr "errors"

	"github.com/pkg/errors"
)

var (
	Wrap   = errors.Wrap
	Wrapf  = errors.Wrapf
	Unwrap = stderr.Unwrap
	New    = stderr.New
	Is     = stderr.Is
	As     = stderr.As
)
