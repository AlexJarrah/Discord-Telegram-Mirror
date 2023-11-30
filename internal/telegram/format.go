package telegram

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/quo0001/Discord-Telegram-Mirror/internal"
)

// Formats a message for Telegram using HTML markup
func Format(msg internal.Message) string {
	divider := fmt.Sprintf("%s\n", strings.Repeat("─", 20))

	// User profile URL
	userURL := fmt.Sprintf("https://discord.com/users/%s", msg.Profile.ID)
	// Message URL
	url := fmt.Sprintf("https://discord.com/channels/%s/%s/%s", msg.Guild, msg.Channel, msg.Message)

	// Add profile details
	res := fmt.Sprintf("<a href='https://discord.com/users/%s'>%s</a> in <a href='%s'>#%s</a>\n", userURL, msg.Profile.Name, url, msg.Channel)

	// Add message content
	if msg.Content != "" {
		res += msg.Content + "\n"
	}

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

		res += divider
		res += fmt.Sprintf("<a href='%s'><b>%s</b>:</a>\n", embed.Body.URL, embed.Body.Title)
		res += fmt.Sprintf("<pre>%s</pre>\n", embed.Body.Description)

		// Add embed fields
		for _, field := range embed.Fields {
			res += fmt.Sprintf("<b>%s</b>: %s\n", field.Name, field.Value)
		}
		res += divider
	}

	// Add image & footer details
	if msg.Embeds[0].Image.ThumbnailURL != "" {
		res += fmt.Sprintf("<a href='%s'><b>Thumbnail</b></a> - ", msg.Embeds[0].Image.ThumbnailURL)
	}
	if msg.Embeds[0].Image.URL != "" {
		res += fmt.Sprintf("<a href='%s'><b>Image</b></a>\n", msg.Embeds[0].Image.URL)
	}
	if msg.Embeds[0].Footer.Text != "" {
		res += fmt.Sprintf("<a href='%s'><b>%s</b></a> - %s\n", msg.Embeds[0].Footer.FooterIconURL, msg.Embeds[0].Footer.Text, msg.Embeds[0].Footer.Timestamp)
	}

	res = strings.TrimSpace(res)

	// Replace bold formatting
	res = regexp.MustCompile(`\*\*(.*?)\*\*`).ReplaceAllString(res, "<strong>$1</strong>")
	// Replace italic formatting
	res = regexp.MustCompile(`\*(.*?)\*`).ReplaceAllString(res, "<em>$1</em>")
	// Replace code block formatting
	res = regexp.MustCompile("`([^`]+)`").ReplaceAllString(res, "<code>$1</code>")
	// Replace inline code formatting
	res = regexp.MustCompile("```([^`]+)```").ReplaceAllString(res, "<pre>$1</pre>")
	// Replace underline formatting
	res = regexp.MustCompile(`__(.*?)__`).ReplaceAllString(res, "<u>$1</u>")
	// Replace strikethrough formatting
	res = regexp.MustCompile(`~~(.*?)~~`).ReplaceAllString(res, "<del>$1</del>")

	return string(res)
}
