package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	cfg, _ := app.Get("*config.Config").(*Config)

	assert.Equal(t, cfg, Get())
}
