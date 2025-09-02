package ticket

import (
	"log"

	"github.com/bwmarrin/discordgo"
)


func HandleSetup(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Check User has the PermissionsLevel to use the Commands

	if i.Member.Permissions&discordgo.PermissionAdministrator == 0 {
		log.Println("User does not have permission to use the command.")
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You do not have permission to use this command.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}
	log.Println("User has permission to use the command.")

	options, err := categoriesToSelectMenu(s, i.GuildID)
	if err != nil {
		log.Printf("Error getting categories: %v", err)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Error getting categories for setup.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Please select a category for new tickets:",
			Flags:   discordgo.MessageFlagsEphemeral, // Keep it private to the user
			Components: []discordgo.MessageComponent{
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						&discordgo.SelectMenu{
							CustomID:    "ticket_category_select",
							Placeholder: "Choose a category",
							Options:     options,
						},
					},
				},
			},
		},
	})
}

func categoriesToSelectMenu(s *discordgo.Session, guildID string) ([]discordgo.SelectMenuOption, error) {
	channels, err := s.GuildChannels(guildID)
	if err != nil {
		return nil, err
	}

	var options []discordgo.SelectMenuOption
	for _, ch := range channels {
		if ch.Type == discordgo.ChannelTypeGuildCategory {
			options = append(options, discordgo.SelectMenuOption{
				Label: ch.Name,
				Value: ch.ID,
			})
		}
	}
	return options, nil
}
