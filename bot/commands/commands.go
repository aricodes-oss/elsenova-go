package commands

import (
	"github.com/bwmarrin/discordgo"
)

type handler = func(s *discordgo.Session, i *discordgo.InteractionCreate)
type handlerMap = map[string]handler
type cmd = discordgo.ApplicationCommand

type Definition struct {
	Name    string
	Base    *discordgo.ApplicationCommand
	Handler handler
}

var (
	commands = []*discordgo.ApplicationCommand{}
	handlers = handlerMap{}
)

func Register(cmd *Definition) {
	cmd.Base.Name = cmd.Name
	commands = append(commands, cmd.Base)
	handlers[cmd.Name] = cmd.Handler
}

func All() ([]*discordgo.ApplicationCommand, handlerMap) {
	return commands, handlers
}
