package commands

import (
	"github.com/bwmarrin/discordgo"
)

var leaderboardsCmd = &Definition{
	"leadervoreds",
	&cmd{
		Description: "Sends a link to the leadervoreds",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "https://elsenova.sudra-routes.com/leaderboards",
			},
		})
	},
}

func init() {
	Register(leaderboardsCmd)
}
