package models

import "github.com/bwmarrin/discordgo"

const (
	TITLE_LIMIT = 256
	DESC_LIMIT  = 2048
	FOOT_LIMIT  = 2048
	EMBED_LIMIT = 4000
)

type Footer struct {
	*discordgo.MessageEmbedFooter
}

func NewEmbedFooter() *Footer {
	return &Footer{&discordgo.MessageEmbedFooter{}}
}
func (f *Footer) SetText(text string) *Footer {
	if f.Text = text; len(text) > FOOT_LIMIT {
		f.Text = text[:FOOT_LIMIT]
	}
	return f
}
func (f *Footer) SetIconURL(url string) *Footer {
	if f.IconURL = url; len(url) > 2048 {
		f.IconURL = url[:2048]
	}
	return f
}

type Embed struct {
	*discordgo.MessageEmbed
}

func NewEmbed() *Embed {
	return &Embed{&discordgo.MessageEmbed{}}
}

func (e *Embed) SetTitle(title string) *Embed {
	if e.Title = title; len(title) > TITLE_LIMIT {
		e.Title = title[:TITLE_LIMIT]
	}
	return e
}

func (e *Embed) SetDescription(desc string) *Embed {
	if e.Description = desc; len(desc) > DESC_LIMIT {
		e.Description = desc[:DESC_LIMIT]
	}
	return e
}

func (e *Embed) SetColor(color int) *Embed {
	e.Color = color
	return e
}

func (e *Embed) SetFooter(footer Footer) *Embed {
	e.Footer = footer.MessageEmbedFooter
	return e
}

func (e *Embed) SetFooterIcon(url string) *Embed {
	e.Footer.IconURL = url
	return e
}

func (e *Embed) SetImage(url string) *Embed {
	e.Image = &discordgo.MessageEmbedImage{URL: url}
	return e
}

func (e *Embed) SetThumbnail(url string) *Embed {
	e.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: url}
	return e
}

func (e *Embed) SetAuthor(name string, iconurl string, url string) *Embed {
	var (
		aName string
		aIcon string
		aUrl  string
	)
	if name != "" {
		aName = name
	}
	if iconurl != "" {
		aIcon = iconurl
	}
	if url != "" {
		aUrl = url
	}

	e.Author = &discordgo.MessageEmbedAuthor{
		Name:    aName,
		IconURL: aIcon,
		URL:     aUrl,
	}
	return e
}

func (e *Embed) SetTimestamp(timestamp string) *Embed {
	e.Timestamp = timestamp
	return e
}

func (e *Embed) SetURL(url string) *Embed {
	e.URL = url
	return e
}
