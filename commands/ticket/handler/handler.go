package handler

import (
	"log/slog"
	"tarmac-fox/commands/ticket/setup"

	"github.com/bwmarrin/discordgo"
)

func HandleTicketComponentInteraction(s *discordgo.Session, i *discordgo.InteractionCreate, name string) {
	
	if (name == "") {
		slog.Error("Couldn't find handle for Component. Name is Empty")
		return
	}

	switch {
	case name == "ticket_category_select":
		setup.HandleCategorySelect(s, i)
	case name == "ticket_log_channel_confirm":
		setup.HandleLogChannelConfirm(s, i)
	case name == "ticket_log_channel_cancel":
		setup.HandleLogChannelCancel(s, i)
	default:
		slog.Error("Couldn't find handle for Component", "name", name)
	}
}