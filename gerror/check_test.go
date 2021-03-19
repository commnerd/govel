package gerror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckNil(t *testing.T) {
	var err error

	Check(err)

	assert.True(t, true)
}

func TestCheckErr(t *testing.T) {
	err := errors.New("tested")

	defer func() {
		if err := recover(); err != nil {
			assert.Equal(t, errors.New("tested"), err)
		}
	}()

	Check(err)
}

func TestCheckString(t *testing.T) {
	err := "tested"

	defer func() {
		if err := recover(); err != nil {
			assert.Equal(t, errors.New("tested"), err)
		}
	}()

	Check(err)
}

type customCheckError struct{}

func (e customCheckError) Error() error {
	return errors.New("tested")
}

func TestCheckGError(t *testing.T) {
	err := customCheckError{}

	defer func() {
		if err := recover(); err != nil {
			assert.Equal(t, errors.New("tested"), err)
		}
	}()

	Check(err)
}
