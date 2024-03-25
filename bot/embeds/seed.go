package embeds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func DailySeed(seed string) *discordgo.MessageEmbed {
	timestamp := time.Now().Format("2 January 2006")

	return &discordgo.MessageEmbed{
		Color:  0x0099ff,
		Author: &discordgo.MessageEmbedAuthor{},
		Title:  "Daily Randomizer Seed!",

		// TODO: fetch this from config
		Description: "<@&823674572855181332>",

		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Seed",
				Value:  seed,
				Inline: true,
			},
			{
				Name:   "Date",
				Value:  timestamp,
				Inline: true,
			},
		},

		// Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
