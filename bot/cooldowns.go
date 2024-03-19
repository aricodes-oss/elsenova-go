package bot

import (
	"github.com/spf13/viper"
	"time"
)

func (b *bot) commandOnCooldown(name string) bool {
	defaultPeriod := viper.GetInt("cooldowns.default")
	specificPeriod := viper.GetInt("cooldowns." + name)

	effectiveCooldown := defaultPeriod
	if viper.IsSet("cooldowns." + name) {
		effectiveCooldown = specificPeriod
	}

	// "The last time this command was run is at least effectiveCooldown seconds ago"
	return !b.lastRunTime[name].Before(
		time.Now().Add(
			time.Duration(effectiveCooldown*int(time.Second)) * -1),
	)
}
