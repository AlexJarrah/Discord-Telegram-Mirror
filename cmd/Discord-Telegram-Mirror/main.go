package main

import (
	"log"
	"os"

	"github.com/quo0001/Discord-Telegram-Mirror/internal/discord"
	"github.com/tidwall/gjson"
)

func main() {
	if _, err := os.Stat("data/config.json"); err != nil {
		os.Mkdir("data", 0755)
		os.Create("data/config.json")
	}

	f, err := os.ReadFile("data/config.json")
	if err != nil {
		log.Panic(err)
	}

	token := gjson.Get(string(f), "discord.token").String()
	discord.Start(token)

	select {}
}
