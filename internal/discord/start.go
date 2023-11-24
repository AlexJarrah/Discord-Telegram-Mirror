package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Start(token string) {
	dg, err := discordgo.New(token)
	if err != nil {
		log.Panic(err)
	}

	dg.Identify.Intents = discordgo.IntentsAll

	dg.AddHandler(messageCreate)

	if err = dg.Open(); err != nil {
		log.Panic(err)
	}
	defer dg.Close()
}
