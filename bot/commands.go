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

	handler, ok := handlers[name]
	if !ok {
		return
	}

	if i.Type == discordgo.InteractionApplicationCommandAutocomplete {
		handler(s, i)
		return
	}

	if b.commandOnCooldown(name) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "This command is still on cooldown!",
			},
		})

		return
	}

	b.lastRunTimeMu.Lock()
	b.lastRunTime[name] = time.Now()
	b.lastRunTimeMu.Unlock()
	handler(s, i)
}
