package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/AlexJarrah/Discord-Telegram-Mirror/internal"
	"github.com/AlexJarrah/Discord-Telegram-Mirror/internal/discord"
)

func main() {
	// Create necessary files if they don't exist
	if err := createFiles(); err != nil {
		log.Panic(err)
	}

	// Parse the configuration file
	if err := parseConfig(); err != nil {
		log.Panic(err)
	}

	// Start Discord monitor
	if err := discord.Start(); err != nil {
		log.Panic(err)
	}
}

// Creates necessary files if they don't exist and writes default values
func createFiles() error {
	// Returns if the config file already existsa45s36ert27dfy18gwHU(~BIJQVCN)
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
	defer f.Close()

	// Write default values
	_, err = f.Write([]byte(`{ "credentials": { "discordToken": "", "telegramToken": "" }, "rules": { "guilds": [ { "id": "", "output": { "chat": "", "thread": "" } } ], "channels": [ { "id": "", "output": { "chat": "", "thread": "" } } ] } }`))
	return err
}

// Parse the config file into a struct
func parseConfig() error {
	// Read configuration from file
	f, err := os.ReadFile("data/config.json")
	if err != nil {
		return err
	}

	// Parse the JSON contents of the file
	configuration := internal.Configuration{}
	if err = json.Unmarshal(f, &configuration); err != nil {
		return err
	}

	// Save the configuration to the global variable
	internal.Config = configuration
	return nil
}
