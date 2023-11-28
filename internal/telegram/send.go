package telegram

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/tidwall/gjson"
)

// Send sends the given text to the specified Telegram chat (data/config.json)
func Send(text string) error {
	// Read configuration from file
	f, err := os.ReadFile("data/config.json")
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	// Extract Telegram settings from config
	token := gjson.Get(string(f), "telegram.botToken").String()
	chat := gjson.Get(string(f), "telegram.outputChat").String()
	thread := gjson.Get(string(f), "telegram.outputThreadId").String()

	// Validate Telegram settings
	if token == "" || chat == "" {
		return fmt.Errorf("invalid bot token or output chat")
	}

	// Build Telegram API URL
	u := fmt.Sprintf("https://api.telegram.org/%s/sendMessage", token)
	url, err := url.Parse(u)
	if err != nil {
		return fmt.Errorf("error parsing url: %w", err)
	}

	// Add query parameters
	q := url.Query()
	q.Add("chat_id", chat)
	q.Add("message_thread_id", thread)
	q.Add("parse_mode", "HTML")
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
