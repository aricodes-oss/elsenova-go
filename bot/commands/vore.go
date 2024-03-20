package commands

import (
	"elsenova/config"
	"elsenova/models"
	"elsenova/query"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var voreCmd = &definition{
	"vore",
	&cmd{
		Description: "Increments the vore counter",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		v := query.Vore
		conf := config.Load()

		v.Create(&models.Vore{
			UserID: i.Member.User.ID,
		})

		// The number prior to migrating to the leadervoreds system
		baseCount := int64(conf.BaseVoreCount)
		recordCount, _ := v.Count()

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("We've talked about vore %d times now. Stop it.", baseCount+recordCount),
			},
		})
	},
}

func init() {
	register(voreCmd)
}
