package ticket

import (
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
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "title",
					Description: "The title of the ticket",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    false,
				},
			},
		},
		{
			Name:        "close",
			Description: "Close an existing ticket",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
	},
}
var SetupCommandTicket = &discordgo.ApplicationCommand{
	Name:        "setup",
	Description: "Setup the bot in the server",
	Type:        discordgo.ChatApplicationCommand,
}

func TicketHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Options[0].Name {
	case "create":
		HandleCreate(s, i)
	case "close":
		// Handle ticket closing
	}
}

func SetupHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	HandleSetup(s, i)
}

// Separate handler for component interactions
func SetupComponentHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.MessageComponentData().CustomID {
	case "ticket_category_select":
		HandleCategorySelect(s, i)
	case "create_new_ticket_category":
		handleCreateNewCategory(s, i)
	}
}
