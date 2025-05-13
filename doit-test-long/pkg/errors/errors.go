package errors

import (
	"github.com/palantir/stacktrace"
)

type Code = stacktrace.ErrorCode

func init() {
	stacktrace.DefaultFormat = stacktrace.FormatFull
}

var ErrCode = stacktrace.GetCode
var New = stacktrace.NewError
var NewWithCode = stacktrace.NewErrorWithCode
var RootCause = stacktrace.RootCause
var Wrap = stacktrace.Propagate
var WrapWithCode = stacktrace.PropagateWithCode
var Wrapf = stacktrace.Propagate
