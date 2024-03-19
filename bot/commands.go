package bot

import (
	"elsenova/models"
	"elsenova/query"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

type CommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate)

func (b *bot) slashCommandRouter(s *discordgo.Session, i *discordgo.InteractionCreate) {
	name := i.ApplicationCommandData().Name

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
			v := query.Vore

			v.Create(&models.Vore{
				UserID: i.Member.User.ID,
			})

			// The number prior to migrating to the leadervoreds system
			base_count := viper.GetInt64("base_vore_count")
			record_count, _ := v.Count()

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("We've talked about vore %d times now. Stop it.", base_count+record_count),
				},
			})
		},
	}
)
