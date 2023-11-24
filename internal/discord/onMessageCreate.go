package discord

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/quo0001/Discord-Telegram-Mirror/internal/telegram"
	"github.com/tidwall/gjson"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	f, err := os.ReadFile("data/config.json")
	if err != nil {
		log.Panic(err)
	}

	guilds := gjson.Get(string(f), "discord.guilds").Array()
	channels := gjson.Get(string(f), "discord.channels").Array()

	if len(guilds) != 0 || len(channels) != 0 {
		var resume bool
		for _, id := range append(guilds, channels...) {
			if id.Value() == m.GuildID || id.Value() == m.ChannelID {
				resume = true
				break
			}
		}

		if !resume {
			return
		}
	}

	msg := parse(m)
	text := telegram.Format(msg)
	telegram.Send(text)
}
