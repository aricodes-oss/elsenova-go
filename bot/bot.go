package bot

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Bot interface {
	// Start brings the bot online.
	Start() error
	// Wait blocks until it receives a shutdown signal.
	Wait()
	// Stop takes the bot offline.
	Stop()
}

func New(token string) (Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	b := &bot{dg: dg}
	b.init()

	return b, err
}

// bot stores values for the bot application.
type bot struct {
	// The raw session pointer
	dg *discordgo.Session

	// Signal handler for receiving shutdown requests
	sc chan os.Signal

	initOnce    sync.Once
	destroyOnce sync.Once
}

func (b *bot) init() {
	b.initOnce.Do(func() {
		b.dg.Identify.Intents = discordgo.IntentsGuildMessages
		b.dg.Open()
		b.sc = make(chan os.Signal, 1)

		// Incoming message event handler
		b.dg.AddHandler(b.messageCreate)

		// Slash command base handler (route to map)
		b.dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if handler, ok := handlers[i.ApplicationCommandData().Name]; ok {
				handler(s, i)
			}
		})

		registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))
		for idx, rawCmd := range commands {
			cmd, err := b.dg.ApplicationCommandCreate(b.dg.State.User.ID, viper.GetString("guild_id"), rawCmd)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to register application commands")
			}
			registeredCommands[idx] = cmd
		}
	})
}

func (b *bot) destroy() {
	b.destroyOnce.Do(func() {
		for _, cmd := range registeredCommands {
			err := b.dg.ApplicationCommandDelete(b.dg.State.User.ID, viper.GetString("guild_id"), cmd.ID)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to deregister command!")
			}
		}
	})
}

func (b *bot) Start() error {
	b.init()

	signal.Notify(b.sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	return nil
}

func (b *bot) Wait() {
	b.init()

	<-b.sc
}

func (b *bot) Stop() {
	b.init()

	log.Info().Msg("Shutting down politely! This might take a moment.")
	b.destroy()
	b.dg.Close()
	log.Info().Msg("Goodbye! 💖")
}