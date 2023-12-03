package discord

import (
	"fmt"

	"github.com/AlexJarrah/Discord-Telegram-Mirror/internal"
	"github.com/bwmarrin/discordgo"
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

	// Update the session's user agent
	dg.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"

	dg.AddHandler(ready)         // Add a ready handler to notify once monitoring has started
	dg.AddHandler(messageCreate) // Add a message creation handler

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
