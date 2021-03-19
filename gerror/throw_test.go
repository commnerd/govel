package gerror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThrowError(t *testing.T) {
	err := errors.New("tested")

	defer func() {
		if err := recover(); err != nil {
			assert.Equal(t, errors.New("tested"), err)
		}
	}()

	Throw(err)
}

func TestThrowString(t *testing.T) {
	err := "tested"

	defer func() {
		if err := recover(); err != nil {
			assert.Equal(t, errors.New("tested"), err)
		}
	}()

	Throw(err)
}

type customThrowError struct{}

func (e customThrowError) Error() error {
	return errors.New("tested")
}

func TestThrowGError(t *testing.T) {
	err := customThrowError{}

	defer func() {
		if err := recover(); err != nil {
			assert.Equal(t, errors.New("tested"), err)
		}
	}()

	Throw(err)
}
