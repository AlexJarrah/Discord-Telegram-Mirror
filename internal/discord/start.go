package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/quo0001/Discord-Telegram-Mirror/internal"
)

// Initiates the Discord monitoring via the provided token
func Start() error {
	// Create a new Discord session
	dg, err := discordgo.New(internal.Config.Credentials.DiscordToken)
	if err != nil {
		return err
	}

	// Specify required intents for the session
	dg.Identify.Intents = discordgo.IntentsAll

	// Add a ready handler to notify once monitoring has started
	dg.AddHandler(ready)
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

// Notifies the user once monitoring has started
func ready(s *discordgo.Session, m *discordgo.Ready) {
	fmt.Printf("Monitor running on %s... Press Ctrl-C to exit\n", s.State.User.Username)
}
