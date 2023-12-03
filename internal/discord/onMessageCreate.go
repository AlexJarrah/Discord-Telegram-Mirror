package discord

import (
	"fmt"
	"log"
	"time"

	"github.com/AlexJarrah/Discord-Telegram-Mirror/internal"
	"github.com/AlexJarrah/Discord-Telegram-Mirror/internal/telegram"
	"github.com/bwmarrin/discordgo"
)

// Handler function for new Discord messages
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the signed in user
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Get the corresponding chat and thread IDs based on the guild and channel IDs
	chat, thread := getChat(m.GuildID, m.ChannelID)

	// Return if the message is not from a monitored guild or channel
	if chat == "" {
		return
	}

	// Increment the message counter & update the monitor stats
	internal.Counter++
	fmt.Printf("\r%d messages forwarded | Latest message: %s", internal.Counter, time.Now().Format("01/02 03:04 PM"))

	// Parse and format the message
	msg := parse(m)
	text := telegram.Format(msg)

	// Send the formatted message to the corresponding chat and thread in Telegram
	if err := telegram.Send(chat, thread, text); err != nil {
		log.Println(err)
	}
}

// Retrieves the corresponding chat and thread IDs based on the provided guild
// and channel IDs. If no rule is found for the provided guild or channel IDs,
// but a wildcard rule ("*") is, it will return the wildcard chat and thread IDs.
func getChat(guildId, channelId string) (chat, thread string) {
	// Initialize variables to store default chat and thread IDs in case no rules are found.
	var defaultChat, defaultThread string

	// Check if the message is from a monitored guild
	for _, rule := range internal.Config.Rules.Guilds {
		if rule.ID == guildId {
			return rule.Output.ChatID, rule.Output.ThreadID
		} else if rule.ID == "*" {
			defaultChat, defaultThread = rule.Output.ChatID, rule.Output.ThreadID
		}
	}

	// Check if the message is from a monitored channel
	for _, rule := range internal.Config.Rules.Channels {
		if rule.ID == channelId {
			return rule.Output.ChatID, rule.Output.ThreadID
		} else if rule.ID == "*" {
			defaultChat, defaultThread = rule.Output.ChatID, rule.Output.ThreadID
		}
	}

	// Return the default chat and thread IDs if no specific rules were found
	return defaultChat, defaultThread
}
