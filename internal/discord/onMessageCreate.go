package discord

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/quo0001/Discord-Telegram-Mirror/internal/telegram"
	"github.com/tidwall/gjson"
)

// Handler function for new Discord messages
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the signed in user
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Read configuration from file
	f, err := os.ReadFile("data/config.json")
	if err != nil {
		log.Println(err)
	}

	// Extract Discord settings from config
	guilds := gjson.Get(string(f), "discord.guilds").Array()
	channels := gjson.Get(string(f), "discord.channels").Array()

	// Check if the message is from a monitored guild or channel. If no guilds
	// or channels are specified, all channels and guilds will be monitored.
	if len(guilds) != 0 || len(channels) != 0 {
		var resume bool
		for _, id := range append(guilds, channels...) {
			if id.Value() == m.GuildID || id.Value() == m.ChannelID {
				resume = true
				break
			}
		}

		// Return if the message is not from a monitored guild or channel
		if !resume {
			return
		}
	}

	// Parse & format the message
	msg := parse(m)
	text := telegram.Format(msg)

	// Send the formatted message to Telegram
	if err = telegram.Send(text); err != nil {
		log.Println(err)
	}
}
