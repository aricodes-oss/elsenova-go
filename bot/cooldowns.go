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

	b.lastRunTimeMu.Lock()
	last := b.lastRunTime[name]
	b.lastRunTimeMu.Unlock()

	// "The last time this command was run is at least effectiveCooldown seconds ago"
	return !last.Before(
		time.Now().Add(
			time.Duration(effectiveCooldown*int(time.Second)) * -1),
	)
}
