package colors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlackConversion(t *testing.T) {
	assert := assert.New(t)
	black := Black()
	assert.Equal(0, black.Convert(), "Black should be 0")
}

func TestWhiteConversion(t *testing.T) {
	assert := assert.New(t)
	white := White()
	assert.Equal(0xFFFFFF, white.Convert(), "White should be 0xFFFFFF")
}
