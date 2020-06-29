package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDB(t *testing.T) {
	assert.NotNil(t, GetDB())
}
