package embed

import (
	"testing"
	"time"

	"github.com/phamill/discord-testground/colors"
	"github.com/stretchr/testify/assert"
)

func TestHolistic(t *testing.T) {
	assert := assert.New(t)
	em := New()
	time := time.Now()
	em.SetTitle("Title").
		SetColor(colors.Red()).
		SetDesc("This is a description").
		SetTimestamp(time)
	assert.Equal(em.MessageEmbed.Title, "Title", "Titles do not match")
	assert.Equal(em.Color, colors.Red().Convert(), "Colors do not match")
	assert.Equal(em.Description, "This is a description", "Descriptions do not match")
	assert.Equal(em.Timestamp, time.String(), "Time does not match")
}
