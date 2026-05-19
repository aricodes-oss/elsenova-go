package bot

import (
	"elsenova/bot/commands"
	"elsenova/config"
	"elsenova/models"
	"elsenova/query"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

// Constrained by what Discord accepts for slash command names
var nameRe = regexp.MustCompile(`^[a-z0-9_-]{1,32}$`)

type customCommandRegistry struct {
	mu     sync.Mutex
	byName map[string]*discordgo.ApplicationCommand
}

func newCustomCommandRegistry() *customCommandRegistry {
	return &customCommandRegistry{byName: map[string]*discordgo.ApplicationCommand{}}
}

func (r *customCommandRegistry) set(name string, cmd *discordgo.ApplicationCommand) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.byName[name] = cmd
}

func (r *customCommandRegistry) take(name string) *discordgo.ApplicationCommand {
	r.mu.Lock()
	defer r.mu.Unlock()
	cmd := r.byName[name]
	delete(r.byName, name)
	return cmd
}

func customCommandHandler(response string) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{Content: response},
		})
	}
}

func commandDescription(row *models.CustomCommand) string {
	if row.Description != "" {
		return row.Description
	}
	return "Custom command"
}

func loadCustomCommands() (map[string]bool, error) {
	rows, err := query.CustomCommand.Find()
	if err != nil {
		return nil, err
	}

	names := make(map[string]bool, len(rows))
	for _, row := range rows {
		commands.Register(&commands.Definition{
			Name:    row.Name,
			Base:    &discordgo.ApplicationCommand{Description: commandDescription(row)},
			Handler: customCommandHandler(row.Response),
		})
		names[row.Name] = true
	}

	log.Info().Int("count", len(rows)).Msg("Loaded custom commands")
	return names, nil
}

func (b *bot) registerManagementCommand() {
	managePermissions := int64(discordgo.PermissionManageMessages)

	commands.Register(&commands.Definition{
		Name: "command",
		Base: &discordgo.ApplicationCommand{
			Description:              "Manage custom moderator commands",
			DefaultMemberPermissions: &managePermissions,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "add",
					Description: "Create a new custom command",
					Options: []*discordgo.ApplicationCommandOption{
						{Type: discordgo.ApplicationCommandOptionString, Name: "name", Description: "Slash command name (a-z, 0-9, _ or -)", Required: true},
						{Type: discordgo.ApplicationCommandOptionString, Name: "response", Description: "What the bot should reply with", Required: true},
						{Type: discordgo.ApplicationCommandOptionString, Name: "description", Description: "Shown in Discord's command picker", Required: false},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "remove",
					Description: "Delete a custom command",
					Options: []*discordgo.ApplicationCommandOption{
						{Type: discordgo.ApplicationCommandOptionString, Name: "name", Description: "Command to remove", Required: true},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "list",
					Description: "List all custom commands",
				},
			},
		},
		Handler: b.handleManageCommand,
	})
}

func (b *bot) handleManageCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	if len(data.Options) == 0 {
		return
	}
	sub := data.Options[0]
	opts := map[string]*discordgo.ApplicationCommandInteractionDataOption{}
	for _, o := range sub.Options {
		opts[o.Name] = o
	}

	switch sub.Name {
	case "add":
		name := strings.ToLower(strings.TrimSpace(opts["name"].StringValue()))
		response := opts["response"].StringValue()
		description := ""
		if o, ok := opts["description"]; ok {
			description = o.StringValue()
		}
		respond(s, i, b.addCustomCommand(name, response, description, i.Member))

	case "remove":
		name := strings.ToLower(strings.TrimSpace(opts["name"].StringValue()))
		respond(s, i, b.removeCustomCommand(name))

	case "list":
		respond(s, i, b.listCustomCommands())
	}
}

func respond(s *discordgo.Session, i *discordgo.InteractionCreate, content string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func (b *bot) addCustomCommand(name, response, description string, member *discordgo.Member) string {
	if !nameRe.MatchString(name) {
		return "Name must be 1–32 chars, lowercase letters/digits/`_`/`-`."
	}

	_, existing := commands.All()
	if _, taken := existing[name]; taken {
		return fmt.Sprintf("A command named `%s` already exists.", name)
	}

	createdBy := ""
	if member != nil && member.User != nil {
		createdBy = member.User.ID
	}

	qcc := query.CustomCommand
	row := &models.CustomCommand{
		Name:        name,
		Description: description,
		Response:    response,
		CreatedBy:   createdBy,
	}
	if err := qcc.Create(row); err != nil {
		log.Error().Err(err).Msg("Failed to persist custom command")
		return "Failed to save the command, sorry!"
	}

	commands.Register(&commands.Definition{
		Name:    name,
		Base:    &discordgo.ApplicationCommand{Description: commandDescription(row)},
		Handler: customCommandHandler(response),
	})

	conf := config.Load()
	created, err := b.dg.ApplicationCommandCreate(b.dg.State.User.ID, conf.GuildID, &discordgo.ApplicationCommand{
		Name:        name,
		Description: commandDescription(row),
	})
	if err != nil {
		// Roll back so we don't leave a stale handler / row dangling
		commands.Unregister(name)
		qcc.Unscoped().Delete(row)
		log.Error().Err(err).Str("name", name).Msg("Failed to register custom command with Discord")
		return "Discord rejected the command — check the logs."
	}

	b.customCommands.set(name, created)
	b.registeredCommands = append(b.registeredCommands, created)
	log.Info().Str("name", name).Str("by", createdBy).Msg("Added custom command")
	return fmt.Sprintf("Added `/%s`.", name)
}

func (b *bot) removeCustomCommand(name string) string {
	qcc := query.CustomCommand
	row, err := qcc.Where(qcc.Name.Eq(name)).First()
	if err != nil {
		return fmt.Sprintf("No custom command named `%s`.", name)
	}

	live := b.customCommands.take(name)
	if live != nil {
		conf := config.Load()
		if err := b.dg.ApplicationCommandDelete(b.dg.State.User.ID, conf.GuildID, live.ID); err != nil {
			log.Error().Err(err).Str("name", name).Msg("Failed to deregister custom command")
		}
		for idx, c := range b.registeredCommands {
			if c.ID == live.ID {
				b.registeredCommands = append(b.registeredCommands[:idx], b.registeredCommands[idx+1:]...)
				break
			}
		}
	}

	commands.Unregister(name)
	qcc.Unscoped().Delete(row)
	log.Info().Str("name", name).Msg("Removed custom command")
	return fmt.Sprintf("Removed `/%s`.", name)
}

func (b *bot) listCustomCommands() string {
	qcc := query.CustomCommand
	rows, err := qcc.Order(qcc.Name).Find()
	if err != nil {
		return "Failed to list commands."
	}
	if len(rows) == 0 {
		return "No custom commands yet."
	}

	var sb strings.Builder
	sb.WriteString("Custom commands:\n")
	for _, r := range rows {
		fmt.Fprintf(&sb, "• `/%s` — %s\n", r.Name, commandDescription(r))
	}
	return sb.String()
}
