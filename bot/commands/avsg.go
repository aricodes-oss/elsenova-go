package commands

import (
	"github.com/bwmarrin/discordgo"
)

var avsgCmd = &Definition{
	"save-decryptor",
	&cmd{
		Description: "Sends a link to the save game encrypt/decrypt utility",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "https://avsg.sudra-routes.com/",
			},
		})
	},
}

func init() {
	Register(avsgCmd)
}
