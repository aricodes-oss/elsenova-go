package commands

import (
	"elsenova/models"
	"elsenova/query"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var sandwichCmd = &Definition{
	"sandwich",
	&cmd{
		Description: "Classifies something as either a sandwich or a dumpling. Results are final!",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "thing",
				Description: "The item you seek to classify",
				Required:    true,
			},
		},
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		qs := query.Sandwich
		options := optionMap(i)

		thing := strings.ToLower(options["thing"].StringValue())

		record, _ := qs.Where(qs.Name.Eq(thing)).First()
		if record == nil {
			record = &models.Sandwich{
				Name:       thing,
				IsSandwich: rand.Float64() < 0.75,
			}

			qs.Create(record)
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: record.Description(),
			},
		})
	},
}

func init() {
	Register(sandwichCmd)
}
