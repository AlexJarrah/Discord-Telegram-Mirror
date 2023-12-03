package telegram

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/AlexJarrah/Discord-Telegram-Mirror/internal"
)

// Send sends the given text to the specified Telegram chat (data/config.json)
func Send(chat, thread, text string) error {
	// Validate Telegram settings
	if internal.Config.Credentials.TelegramToken == "" || chat == "" || text == "" {
		return fmt.Errorf("invalid telegram bot token, chat id, or text")
	}

	// Build Telegram API URL
	u := fmt.Sprintf("https://api.telegram.org/%s/sendMessage", internal.Config.Credentials.TelegramToken)
	url, err := url.Parse(u)
	if err != nil {
		return fmt.Errorf("error parsing url: %w", err)
	}

	// Add query parameters
	q := url.Query()
	q.Add("chat_id", chat)
	q.Add("message_thread_id", thread)
	q.Add("parse_mode", "HTML")
	q.Add("disable_web_page_preview", "true")
	q.Add("text", text)
	url.RawQuery = q.Encode()

	// Send the HTTP request to send the message
	r, err := http.Get(url.String())
	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}
	defer r.Body.Close()

	return nil
}
