# Embed library for DiscordGo

Inspired by experience with OOP and the simplified example found in the DiscordGo Wiki [here](https://github.com/bwmarrin/discordgo/wiki/FAQ#simplifying-embeds)
All credit to paul-io

#### Purpose

Provide a Builder pattern for Embeds, including a small color library and presets.

#### Usage

First, grab the library

```
go get github.com/paul-io/discordgo-embeds
```
alternatively
```
go get github.com/Coaltergeist/discordgo-embeds
```

Then, you can utilize this library by first creating an instance of the Embed struct

```
em := embed.New()
```

Then set its attributes

```
em.SetTitle("Example title").AddField("Title", "Description", true)
```

Then send it off through your session!

```
session.ChannelSendMessageEmbed(channelID, em.MessageEmbed)
```

#### Working with colors

There are two methods to help with colors. One takes 3 integers to form an RGB color, the other uses premade colors

```
em.SetColorRGB(255, 0, 0) // This creates 0xFF0000, or Red
```

```
em.SetColor(colors.Red()) // This uses the premade pattern for Red
```

If you want to create a reusable, custom struct to hold your color:

```
// This creates a color c which is encoded as 0xFF0000, red
c := colors.Color{
		R: 255,
		G: 0,
		B: 0,
	}
```

Then pass `c` into `SetColor`
