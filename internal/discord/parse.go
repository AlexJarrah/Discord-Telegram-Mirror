package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/quo0001/Discord-Telegram-Mirror/internal"
)

// Parses a Discord message into struct
func parse(m *discordgo.MessageCreate) internal.Message {
	msg := internal.Message{
		Guild:   m.GuildID,
		Channel: m.ChannelID,
		Message: m.ID,
		Content: m.Content,
		Profile: internal.Profile{
			ID:        m.Author.ID,
			Name:      m.Author.Username,
			AvatarURL: m.Author.AvatarURL(""),
		},
	}

	for _, e := range m.Embeds {
		res := internal.Embed{}

		res = internal.Embed{
			Body: internal.Body{
				Title:       e.Title,
				Description: e.Description,
				URL:         e.URL,
				Color:       e.Color,
			},
		}

		if e.Author != nil {
			res.Author = internal.Author{
				Name:    e.Author.Name,
				URL:     e.Author.URL,
				IconURL: e.Author.IconURL,
			}
		}

		if e.Image != nil {
			res.Image.URL = e.Image.URL
			if e.Thumbnail != nil {
				res.Image.ThumbnailURL = e.Image.URL
			}
		}

		if e.Footer != nil {
			res.Footer = internal.Footer{
				Text:          e.Footer.Text,
				Timestamp:     e.Timestamp,
				FooterIconURL: e.Footer.IconURL,
			}
		}

		for _, f := range e.Fields {
			res.Fields = append(res.Fields, internal.Field{
				Name:  f.Name,
				Value: f.Value,
			})
		}

		msg.Embeds = append(msg.Embeds, res)
	}

	return msg
}
