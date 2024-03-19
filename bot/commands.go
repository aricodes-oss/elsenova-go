package bot

import (
	"github.com/bwmarrin/discordgo"
)

type CommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate)

var (
	// Stores command registration info for deletion on shutdown
	registeredCommands []*discordgo.ApplicationCommand

	// The command+argument specifications
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "vore",
			Description: "Increments the vore counter",
		},
	}

	// The implementation of each command
	handlers = map[string]CommandHandler{
		"vore": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "God damn it, really?",
				},
			})
		},
	}
)
