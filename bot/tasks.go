package bot

import (
	"elsenova/bot/embeds"
	"elsenova/config"
	"elsenova/models"
	"elsenova/query"
	"elsenova/util"
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rs/zerolog/log"
)

func (b *bot) postDailySeed() {
	Seed := query.Seed
	conf := config.Load()
	if conf.DailySeed.Channel == "" {
		log.Warn().Msg("config.dailySeedChannel not provided, skipping post")
		return
	}

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

	_, err := b.dg.ForumThreadStartEmbed(
		conf.DailySeed.Channel,
		fmt.Sprintf("Daily seed - %s - %s", seed, time.Now().Format("2 January 2006")),
		4320, // Three days, in hours
		embeds.DailySeed(seed),
	)
	if err != nil {
		log.Error().Err(err).Msg("Unable to create daily seed thread")
		return
	}
	log.Info().Msgf("Created daily seed thread! [%s]", seed)
}
