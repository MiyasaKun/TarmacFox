package setup

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

// TODO: Handle Command use
func HandleSetup(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Check if the user has admin permissions on the guild
	// If not respond with an error message
	if (i.Member.Permissions & discordgo.PermissionAdministrator) == 0 {
		slog.Warn("No Permission")
		// Respond with embeded error message
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Setup Error - Permission Denied",
						Description: "You do not have the required permissions to use this command. Administrator permission is required.",
						Color:       0xff0000,
					},
				},
			},
		})

		if err != nil {
			slog.Error("Failed to send permission error message", "error", err)
		}

		return
	}

	// Proceed with the setup process

	// Send Message to set Name for the Ticket Channel

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Ticket Setup - Step 1",
					Description: "Please enter the name for your ticket channel.\n\n This Channel will be used to create new support tickets.\n\nExample: `support-tickets`",
					Color:       0x00ff00,
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Set Ticket Channel Name",
							CustomID: "ticket_btn_set_name",
							Style:    discordgo.PrimaryButton,
						},
					},
				},
			},
		},
	})

	if err != nil {
		slog.Error("Failed to send ticket name prompt", "error", err)
	}

}

func HandleSetName(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: "ticket_modal_set_name",
			Title:    "Set Ticket Channel Name",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "ticket_field_name_input",
							Label:       "Ticket Channel Name",
							Style:       discordgo.TextInputShort,
							Placeholder: "Input the ticket channel name",
							Required:    true,
						},
					},
				},
			},
		},
	})

	if err != nil {
		slog.Error("Failed to send ticket name modal", "error", err)
	}
}

//TODO: Handle Ticket Name
//TODO: Handle Support Role Selection
//TODO: Handle Config Save to DB
//TODO: Handle Confirmation Message
