package main

import (
	"log"
	"os"

	"github.com/quo0001/Discord-Telegram-Mirror/internal/discord"
	"github.com/tidwall/gjson"
)

func main() {
	if err := createFiles(); err != nil {
		log.Panic(err)
	}

	f, err := os.ReadFile("data/config.json")
	if err != nil {
		log.Panic(err)
	}

	token := gjson.Get(string(f), "discord.token").String()
	discord.Start(token)

	select {}
}

func createFiles() error {
	if _, err := os.Stat("data/config.json"); err == nil {
		return nil
	}

	if err := os.Mkdir("data", 0755); err != nil {
		return err
	}

	f, err := os.Create("data/config.json")
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(`{ "discord": { "token": "", "guilds": [], "channels": [] }, "telegram": { "botToken": "", "outputChat": "", "outputThreadId": "" } }`))
	return err
}
