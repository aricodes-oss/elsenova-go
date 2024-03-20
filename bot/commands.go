package bot

import (
	"elsenova/bot/commands"
	"time"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate)

func (b *bot) slashCommandRouter(s *discordgo.Session, i *discordgo.InteractionCreate) {
	name := i.ApplicationCommandData().Name
	_, handlers := commands.All()

	if handler, ok := handlers[name]; ok {
		if b.commandOnCooldown(name) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "This command is still on cooldown!",
				},
			})

			return
		}

		b.lastRunTime[name] = time.Now()
		handler(s, i)
	}
}

var (
	// Stores command registration info for deletion on shutdown
	registeredCommands []*discordgo.ApplicationCommand
)
