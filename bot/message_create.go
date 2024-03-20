package bot

import (
	"elsenova/config"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func (b *bot) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Drop all messages from ourselves, to avoid recursion
	if m.Author.ID == s.State.User.ID {
		return
	}

	conf := config.Load()

	if conf.DebugLogging {
		msg := fmt.Sprintf("RCV [%v] - %v", m.ChannelID, m.Content)
		log.Info().Msg(msg)
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
