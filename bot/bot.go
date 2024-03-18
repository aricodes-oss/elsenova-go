package bot

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

type Bot interface {
	Start() error
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
	sc       chan os.Signal
	initOnce sync.Once
}

func (b *bot) init() {
	b.initOnce.Do(func() {
		b.sc = make(chan os.Signal, 1)
	})
}

func (b *bot) Start() error {
	b.init()

	// TODO: add handlers
	b.dg.Identify.Intents = discordgo.IntentsGuildMessages

	err := b.dg.Open()
	if err != nil {
		return err
	}

	signal.Notify(b.sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	return nil
}
