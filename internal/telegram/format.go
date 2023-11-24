package telegram

import (
	"fmt"

	"github.com/quo0001/Discord-Telegram-Mirror/internal"
)

func Format(msg internal.Message) string {
	// url := fmt.Sprintf("https://discord.com/channels/%s/%s/%s", msg.Guild, msg.Channel, msg.Message)

	res := fmt.Sprintf("<a href='%s'>%s</a>\n", msg.Profile.AvatarURL, msg.Profile.Name)
	res += msg.Content

	if len(msg.Embeds) == 0 {
		return res
	}

	res += "\n\n-------------------------------------------------"
	for _, embed := range msg.Embeds {
		if embed.Body.Title == "" {
			continue
		}

		res += fmt.Sprintf("\n<a href='%s'><b>%s</b>:</a>\n", embed.Body.URL, embed.Body.Title)
		res += embed.Body.Description + "\n\n"
		for _, field := range embed.Fields {
			res += fmt.Sprintf("<b>%s</b>: %s\n", field.Name, field.Value)
		}
	}

	res += "-------------------------------------------------\n"
	res += fmt.Sprintf("<a href='%s'><b>Thumbnail</b></a> - ", msg.Embeds[0].Image.ThumbnailURL)
	res += fmt.Sprintf("<a href='%s'><b>Image</b></a>\n", msg.Embeds[0].Image.URL)
	res += fmt.Sprintf("<a href='%s'><b>%s</b>:</a> - %s\n", msg.Embeds[0].Footer.FooterIconURL, msg.Embeds[0].Footer.Text, msg.Embeds[0].Footer.Timestamp)

	return string(res)
}
