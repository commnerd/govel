package config

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	v := viper.New()
	n := New()
	n.Viper = v
	assert.Equal(t, &Config{Viper: v}, n)
}
