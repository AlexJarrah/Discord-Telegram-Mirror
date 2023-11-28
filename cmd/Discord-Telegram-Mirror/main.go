package main

import (
	"log"
	"os"

	"github.com/quo0001/Discord-Telegram-Mirror/internal/discord"
	"github.com/tidwall/gjson"
)

func main() {
	// Create necessary files if they don't exist
	if err := createFiles(); err != nil {
		log.Panic(err)
	}

	// Read configuration from file
	f, err := os.ReadFile("data/config.json")
	if err != nil {
		log.Panic(err)
	}

	// Extract Discord token from config
	token := gjson.Get(string(f), "discord.token").String()

	// Start Discord monitor
	if err := discord.Start(token); err != nil {
		log.Panic(err)
	}
}

// Creates necessary files if they don't exist and writes default values
func createFiles() error {
	// Returns if the config file already exists
	if _, err := os.Stat("data/config.json"); err == nil {
		return nil
	}

	// Creates the data directory
	if err := os.Mkdir("data", 0755); err != nil {
		return err
	}

	// Creates the config file in the data directory
	f, err := os.Create("data/config.json")
	if err != nil {
		return err
	}

	// Write default values
	_, err = f.Write([]byte(`{ "discord": { "token": "", "guilds": [], "channels": [] }, "telegram": { "botToken": "", "outputChat": "", "outputThreadId": "" } }`))
	return err
}
