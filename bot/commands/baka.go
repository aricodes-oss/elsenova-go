package commands

import (
	"github.com/bwmarrin/discordgo"
)

var bakaCmd = &Definition{
	"baka",
	&cmd{
		Description: "It's not like I like you or anything!",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "B-b-b-baka!",
			},
		})
	},
}

func init() {
	Register(bakaCmd)
}
