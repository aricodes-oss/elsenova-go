package bot

import (
	"elsenova/bot/commands"
	"elsenova/config"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

var conf *config.Config

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

	// Command cooldowns
	lastRunTime map[string]time.Time

	// Cron instance for scheduling
	scheduler *cron.Cron

	initOnce    sync.Once
	destroyOnce sync.Once
}

func (b *bot) init() {
	b.initOnce.Do(func() {
		conf := config.Load()

		b.dg.Identify.Intents = discordgo.IntentsGuildMessages
		b.dg.Open()

		// Initialize fields
		b.sc = make(chan os.Signal, 1)
		b.lastRunTime = make(map[string]time.Time)
		b.scheduler = cron.New()

		// Attach handlers
		b.dg.AddHandler(b.messageCreate)      // Incoming message
		b.dg.AddHandler(b.slashCommandRouter) // Slash command (route to map, see commands.go and ./commands)

		// Attach scheduled jobs
		_, err := b.scheduler.AddFunc(conf.DailySeed.Schedule, b.postDailySeed)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to start scheduler")
		}

		// We have to be able to address this so the allocation is necessary
		managePermissions := int64(discordgo.PermissionManageMessages)
		commands.Register(&commands.Definition{
			Name: "trigger-seed",
			Base: &discordgo.ApplicationCommand{
				Description:              "Posts a new daily seed thread, off-schedule",
				DefaultMemberPermissions: &managePermissions,
			},
			Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
				b.postDailySeed()
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Seed posted!",
					},
				})
			},
		})

		// Register application commands
		cmdList, _ := commands.All()
		registeredCommands = make([]*discordgo.ApplicationCommand, len(cmdList))
		for idx, rawCmd := range cmdList {
			cmd, err := b.dg.ApplicationCommandCreate(b.dg.State.User.ID, conf.GuildID, rawCmd)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to register application commands")
			}
			registeredCommands[idx] = cmd
		}

		// Start the scheduler
		b.scheduler.Start()
	})
}

func (b *bot) destroy() {
	b.destroyOnce.Do(func() {
		conf = config.Load()

		// Stop the scheduler
		b.scheduler.Stop()

		// Detach application commands
		for _, cmd := range registeredCommands {
			err := b.dg.ApplicationCommandDelete(b.dg.State.User.ID, conf.GuildID, cmd.ID)
			if err != nil {
				log.Error().Err(err).Msg("Failed to deregister command!")
			}
		}
	})
}

func (b *bot) Start() error {
	b.init()

	signal.Notify(b.sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT, os.Interrupt)
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
