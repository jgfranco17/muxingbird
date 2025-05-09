package errorx

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

// ErrorWithCode wraps an error with an associated exit code.
type ErrorWithCode struct {
	Err  error
	Code ShellExitCode
}

// Error satisfies the error interface.
func (e *ErrorWithCode) Error() string {
	return e.Err.Error()
}

// Unwrap enables error unwrapping.
func (e *ErrorWithCode) Unwrap() error {
	return e.Err
}

// NewErrorWithCode returns a new ErrorWithCode.
func NewErrorWithCode(err error, code ShellExitCode) error {
	if err == nil {
		return nil
	}
	return &ErrorWithCode{
		Err:  err,
		Code: code,
	}
}

// ExtractCode returns the exit code associated with the error.
// If the error does not implement ErrorWithCode, it returns
// an generic exit error.
func ExtractCode(err error) ShellExitCode {
	var ew *ErrorWithCode
	if errors.As(err, &ew) {
		return ew.Code
	}
	return ExitGenericError
}

// HandleError prints the error message to stderr and exits
// the program with the correct code.
func HandleError(logger *logrus.Logger, err error) {
	code := ExtractCode(err)
	logger.Errorf("Muxingbird error: %v", err)
	os.Exit(int(code))
}

func HandleRecovery(logger *logrus.Logger) {
	if rec := recover(); rec != nil {
		logger.Panicf("Panic recovered: %v", rec)
		os.Exit(int(ExitPanic))
	}
}
