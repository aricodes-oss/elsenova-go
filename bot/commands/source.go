package commands

import (
	"github.com/bwmarrin/discordgo"
)

var sourceCmd = &Definition{
	"source",
	&cmd{
		Description: "Sends a link to this bot's GitHub repo",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "https://github.com/aricodes-oss/elsenova-go",
			},
		})
	},
}

func init() {
	Register(sourceCmd)
}
