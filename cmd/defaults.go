package cmd

// defaults contains a map of config keys to their default values
var defaults = map[string]any{
	"databaseUrl": "sqlite://data.db",

	"baseVoreCount":     0,
	"sandwichFrequency": 0.75,

	"dailySeed.schedule": "0 7 * * *",

	"web.port":     4000,
	"web.frontend": "http://localhost:4001",

	"debugLogging": true,

	"cooldowns.default": 2,
	"cooldowns.vore":    60,
}
