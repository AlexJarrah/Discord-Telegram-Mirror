package telegram

import (
	"fmt"

	"github.com/quo0001/Discord-Telegram-Mirror/internal"
)

// Formats a message for Telegram using HTML markup
func Format(msg internal.Message) string {
	// Message URL
	url := fmt.Sprintf("https://discord.com/channels/%s/%s/%s", msg.Guild, msg.Channel, msg.Message)

	// Add profile details
	res := fmt.Sprintf("<a href='%s'>%s</a>\n", url, msg.Profile.Name)

	// Add message content
	res += msg.Content

	// Return if message has no embeds
	if len(msg.Embeds) == 0 {
		return res
	}

	// Add embeds
	for _, embed := range msg.Embeds {
		// Skip if invalid embed
		if embed.Body.Title == "" {
			continue
		}

		res += "\n\n-------------------------------------------------"
		res += fmt.Sprintf("\n<a href='%s'><b>%s</b>:</a>\n", embed.Body.URL, embed.Body.Title)
		res += embed.Body.Description + "\n\n"

		// Add embed fields
		for _, field := range embed.Fields {
			res += fmt.Sprintf("<b>%s</b>: %s\n", field.Name, field.Value)
		}
	}

	// Add image & footer details
	res += "-------------------------------------------------\n"
	res += fmt.Sprintf("<a href='%s'><b>Thumbnail</b></a> - ", msg.Embeds[0].Image.ThumbnailURL)
	res += fmt.Sprintf("<a href='%s'><b>Image</b></a>\n", msg.Embeds[0].Image.URL)
	res += fmt.Sprintf("<a href='%s'><b>%s</b>:</a> - %s\n", msg.Embeds[0].Footer.FooterIconURL, msg.Embeds[0].Footer.Text, msg.Embeds[0].Footer.Timestamp)

	return string(res)
}
