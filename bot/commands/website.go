package commands

import (
	"github.com/bwmarrin/discordgo"
)

var websiteCmd = &definition{
	"website",
	&cmd{
		Description: "Sends a link to this bot's web interface",
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
	register(websiteCmd)
}
