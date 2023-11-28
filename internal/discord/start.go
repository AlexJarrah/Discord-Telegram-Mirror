package discord

import (
	"github.com/bwmarrin/discordgo"
)

// Initiates the Discord monitoring via the provided token
func Start(token string) error {
	// Create a new Discord session
	dg, err := discordgo.New(token)
	if err != nil {
		return err
	}

	// Specify required intents for the session
	dg.Identify.Intents = discordgo.IntentsAll

	// Add a message creation handler
	dg.AddHandler(messageCreate)

	// Open a websocket connection
	if err := dg.Open(); err != nil {
		return err
	}
	defer dg.Close()

	// Wait for signals to exit
	select {}
}
