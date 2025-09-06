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

    // Create the embed with markdown formatting
    embed := &discordgo.MessageEmbed{
        Title:       "🎫 Ticket System Setup",
        Description: "**Step 1:** Configure your ticket category\n\nPlease select an existing category for new tickets, or create a new one if needed.",
        Color:       0x5865F2, // Discord blurple color
        Fields: []*discordgo.MessageEmbedField{
            {
                Name:   "📋 Instructions",
                Value:  "• Select a category from the dropdown below\n• Or click **Create New Category** to make a new one\n• This category will contain all new ticket channels",
                Inline: false,
            },
        },
        Footer: &discordgo.MessageEmbedFooter{
            Text: "Only administrators can configure the ticket system",
        },
    }

    var components []discordgo.MessageComponent

    // Add select menu only if there are existing categories
    if len(options) > 0 {
        components = append(components, &discordgo.ActionsRow{
            Components: []discordgo.MessageComponent{
                &discordgo.SelectMenu{
                    CustomID:    "ticket_category_select",
                    Placeholder: "Choose an existing category...",
                    Options:     options,
                },
            },
        })
    }

    // Add the create new category button
    components = append(components, &discordgo.ActionsRow{
        Components: []discordgo.MessageComponent{
            &discordgo.Button{
                Label:    "Create New Category",
                Style:    discordgo.PrimaryButton,
                CustomID: "create_new_ticket_category",
                Emoji: &discordgo.ComponentEmoji{
                    Name: "➕",
                },
            },
        },
    })

    // If no categories exist, show a different message
    if len(options) == 0 {
        embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
            Name:   "ℹ️ No Categories Found",
            Value:  "No categories were found in this server.\nClick the button below to create your first category.",
            Inline: false,
        })
    }

    err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Embeds:     []*discordgo.MessageEmbed{embed},
            Flags:      discordgo.MessageFlagsEphemeral,
            Components: components,
        },
    })
    
    if err != nil {
        log.Printf("Error sending interaction response: %v", err)
    }
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
                Label:       ch.Name,
                Value:       ch.ID,
                Description: "Use this existing category for tickets",
                Emoji: &discordgo.ComponentEmoji{
                    Name: "📁",
                },
            })
        }
    }
    return options, nil
}

func HandleCategorySelect(s *discordgo.Session, i *discordgo.InteractionCreate) {
	selectedCategoryID := i.MessageComponentData().Values[0]

	// Respond to the interaction
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content:    "Category selected successfully! The ID is: " + selectedCategoryID,
			Components: []discordgo.MessageComponent{}, // This removes the buttons and dropdown from the original message
		},
	})
	if err != nil {
		log.Printf("Error sending interaction response: %v", err)
	}
}
func handleCreateNewCategory(s *discordgo.Session, i *discordgo.InteractionCreate) {
	
}
