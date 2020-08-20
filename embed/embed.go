// Package embed contains a builder pattern for Embed building
package embed

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/Coaltergeist/discordgo-embeds/colors"
)

// An Embed is a struct that contains a pointer to a discordgo Embed object
type Embed struct {
	*discordgo.MessageEmbed
}

// New creates an empty Embed and returns the struct
func New() *Embed {
	return &Embed{&discordgo.MessageEmbed{}}
}

// SetTitle sets the title of this embed
func (embed *Embed) SetTitle(title string) *Embed {
	embed.Title = title
	return embed
}

// SetAuthor sets the author field, consisting of a Title, Direct Link, and small Icon
func (embed *Embed) SetAuthor(author, url, icon string) *Embed {
	embed.Author = &discordgo.MessageEmbedAuthor{
		URL:     url,
		Name:    author,
		IconURL: icon,
	}
	return embed
}

// SetColorRGB sets the color of the embed with given 0-256 RGB ints
func (embed *Embed) SetColorRGB(r, g, b int) *Embed {
	embed.Color = (r * 256 * 256) + (g * 256) + b
	return embed
}

// SetColor sets the color of the embed with a premade color
func (embed *Embed) SetColor(c *colors.Color) *Embed {
	embed.Color = c.Convert()
	return embed
}

// SetDescription sets the description of the embed
func (embed *Embed) SetDescription(desc string) *Embed {
	embed.Description = desc
	return embed
}

// SetDesc is an alias for SetDescription
func (embed *Embed) SetDesc(desc string) *Embed {
	return embed.SetDescription(desc)
}

// AddField adds a field with a title, content, and a flag to indicate whether or not
// it's inline with the rest of the fields on this line (3 max)
func (embed *Embed) AddField(title, content string, inline bool) *Embed {
	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   title,
		Value:  content,
		Inline: inline,
	})
	return embed
}

// SetImage sets the image of this embed to a URL
func (embed *Embed) SetImage(url string) *Embed {
	embed.Image = &discordgo.MessageEmbedImage{
		URL: url,
	}
	return embed
}

// SetThumbnail sets the thumbnail of this embed to a URL
func (embed *Embed) SetThumbnail(url string) *Embed {
	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL: url,
	}
	return embed
}

// SetTimestamp sets the timestamp to a certain Time's String() value
func (embed *Embed) SetTimestamp(time time.Time) *Embed {
	embed.Timestamp = time.UTC().Format("2006-01-02T15:04:05Z07:00")
	return embed
}

// SetFooter sets the footer of this embed
func (embed *Embed) SetFooter(footer string) *Embed {
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: footer,
	}
	return embed
}
