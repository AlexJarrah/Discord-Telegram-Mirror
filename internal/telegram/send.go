package telegram

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/tidwall/gjson"
)

func Send(text string) {
	f, err := os.ReadFile("data/config.json")
	if err != nil {
		log.Panic(err)
	}

	botToken := gjson.Get(string(f), "telegram.botToken").String()
	outputChat := gjson.Get(string(f), "telegram.outputChat").String()
	outputThreadId := gjson.Get(string(f), "telegram.outputThreadId").String()

	if botToken == "" || outputChat == "" {
		return
	}

	url, err := url.Parse(fmt.Sprintf("https://api.telegram.org/%s/sendMessage", botToken))
	if err != nil {
		log.Println(err)
		return
	}

	q := url.Query()
	q.Add("chat_id", outputChat)
	q.Add("message_thread_id", outputThreadId)
	q.Add("parse_mode", "HTML")
	q.Add("text", text)
	url.RawQuery = q.Encode()

	r, err := http.Get(url.String())
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()
}
