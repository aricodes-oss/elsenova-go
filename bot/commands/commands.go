package commands

import (
	"github.com/bwmarrin/discordgo"
)

type handler = func(s *discordgo.Session, i *discordgo.InteractionCreate)
type handlerMap = map[string]handler
type cmd = discordgo.ApplicationCommand

type definition struct {
	name    string
	base    *discordgo.ApplicationCommand
	handler handler
}

var (
	commands = []*discordgo.ApplicationCommand{}
	handlers = handlerMap{}
)

func register(cmd *definition) {
	cmd.base.Name = cmd.name
	commands = append(commands, cmd.base)
	handlers[cmd.name] = cmd.handler
}

func All() ([]*discordgo.ApplicationCommand, handlerMap) {
	return commands, handlers
}
