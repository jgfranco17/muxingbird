package errorx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewErrorWithCode(t *testing.T) {
	t.Run("returns ErrorWithCode when error is non-nil", func(t *testing.T) {
		err := errors.New("test error")
		code := ExitGenericError
		customErr := NewErrorWithCode(err, code)
		require.NotNil(t, customErr)
		assert.IsType(t, &ErrorWithCode{}, customErr)
		assert.Equal(t, "test error", customErr.Error())
		assert.Equal(t, code, customErr.(*ErrorWithCode).Code)
	})

	t.Run("returns nil when error is nil", func(t *testing.T) {
		customErr := NewErrorWithCode(nil, ExitGenericError)
		assert.Nil(t, customErr)
	})
}

func TestExtractCode(t *testing.T) {
	t.Run("returns exit code from ErrorWithCode", func(t *testing.T) {
		err := errors.New("test error")
		code := ExitInvalidArgs
		customErr := NewErrorWithCode(err, code)
		extractedCode := ExtractCode(customErr)
		assert.Equal(t, code, extractedCode)
	})

	t.Run("returns generic error code if error does not implement ErrorWithCode", func(t *testing.T) {
		err := errors.New("test error")
		extractedCode := ExtractCode(err)
		assert.Equal(t, ExitGenericError, extractedCode)
	})
}
