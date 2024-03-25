package commands

import (
	"elsenova/bot/embeds"
	"elsenova/models"
	"elsenova/query"
	"elsenova/util"

	"github.com/bwmarrin/discordgo"
	"github.com/deckarep/golang-set/v2"
	"github.com/rs/zerolog/log"
)

var seedCmd = &Definition{
	"seed",
	&cmd{
		Description: "Creates a new daily rando seed embed and sends it",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		Seed := query.Seed

		// We're loading all of these into memory and checking them there
		// instead of issuing queries for each to take advantage of a small dataset
		// (and comparatively high network/disk latency between bot -> db)
		knownSeeds := mapset.NewSet[string]()

		// Using IIFE to avoid polluting local namespace
		func() {
			// TODO: handle error response
			all, _ := Seed.Select(Seed.Value).Find()
			util.Map(all, func(record *models.Seed, _ int) string {
				knownSeeds.Add(record.Value)
				return record.Value
			})
		}()

		// The selected seed to send
		var seed string
		for {
			seed = util.RandoSeed()
			if !knownSeeds.ContainsOne(seed) {
				break
			}
		}

		// Save in the database so we don't get this one again
		Seed.Create(&models.Seed{Value: seed})

		// Send it!
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embeds.DailySeed(seed)},
			},
		})

		log.Info().Msgf("Posted daily seed %s", seed)
	},
}

func init() {
	Register(seedCmd)
}
