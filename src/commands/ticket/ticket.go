package ticket

import (
	"tarmac-fox/src/commands/ticket/setup"
	"github.com/bwmarrin/discordgo"
)

var CommandTicket = &discordgo.ApplicationCommand{
	Name:        "ticket",
	Description: "Ticket Control Command",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "create",
			Description: "Create a new ticket",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
		{
			Name:        "close",
			Description: "Close an existing ticket",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "ticket_id",
					Description: "The ID of the ticket to close",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		},
	},
}

var SetupCommandTicket = &discordgo.ApplicationCommand{
	Name:        "setup",
	Description: "Setup the bot for the current server",
	Type:        discordgo.ChatApplicationCommand,
}

func TicketHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Options[0].Name {
	case "create":
		break
	default:
		break
	}
}

func SetupHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	setup.HandleSetup(s, i)
}
