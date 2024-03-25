package commands

import (
	"github.com/bwmarrin/discordgo"
)

var wikiCmd = &Definition{
	"wiki",
	&cmd{
		Description: "Sends a link to the Axiom Verge speedrunning wiki",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "https://sudra-routes.com/",
			},
		})
	},
}

func init() {
	Register(wikiCmd)
}
