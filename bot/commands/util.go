package commands

import "github.com/bwmarrin/discordgo"

type option = *discordgo.ApplicationCommandInteractionDataOption

func optionMap(i *discordgo.InteractionCreate) (ret map[string]option) {
	options := i.ApplicationCommandData().Options
	ret = make(map[string]option, len(options))
	for _, opt := range options {
		ret[opt.Name] = opt
	}

	return
}
