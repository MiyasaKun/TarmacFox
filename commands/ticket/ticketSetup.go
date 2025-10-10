package ticket

import (
	"log"
	"log/slog"
	"tarmac-fox/helper"

	"github.com/bwmarrin/discordgo"
)

var ticketCategory *discordgo.Channel
var logChannel *discordgo.Channel

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
	var err error

	selectedCategoryID := i.MessageComponentData().Values[0]
	// Jump to next step - Log Channel Selection
	handleLogChannel(s, i)
	// Store the ticketCategory globally in this package
	ticketCategory, err = s.State.Channel(selectedCategoryID)
	//TODO Safe the Category to the Database

	if err != nil {
		log.Printf("Error fetching selected category channel: %v", err)
	}

	log.Println("Ticket Category set to:", ticketCategory.Name)
}

func handleLogChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {

	embed := &discordgo.MessageEmbed{
		Title:       "🎫 Ticket System Setup",
		Description: "Step 2 - Please select a log channel for the ticket system.",
		Color:       0x5865F2, // Discord blurple color
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "📋 Instructions",
				Value:  "• The Bot will now create a log channel for ticket events\n• This channel will be used to log ticket creations, closures, and other important events\n• Please ensure the bot has permission to send messages in this channel",
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Please Confirm that the Bot is allowed to Create a Log Channel",
		},
	}
	components := []discordgo.MessageComponent{
		&discordgo.Button{
			CustomID: "ticket_log_channel_confirm",
			Label:    "Confirm",
			Style:    discordgo.SuccessButton,
		},
		&discordgo.Button{
			CustomID: "ticket_log_channel_cancel",
			Label:    "Cancel",
			Style:    discordgo.DangerButton,
		},
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
			Flags:  discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				&discordgo.ActionsRow{
					Components: components,
				},
			},
		},
	})

	if err != nil {
		log.Printf("Error sending interaction response: %v", err)
	}

}

func HandleLogChannelConfirm(s *discordgo.Session, i *discordgo.InteractionCreate) {

	// Final Step - Setup Complete
	embed := &discordgo.MessageEmbed{
		Title:       "🎫 Ticket System Setup",
		Description: "Thanks for confirming the log channel setup. The bot will now create the log channel and finalize the ticket system configuration.",
		Color:       0x5865F2, // Discord blurple color
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "📋 Instructions",
				Value:  "The Setup is complete! You can now start using the ticket system with the `/ticket create` command.\n\nIf you need to change any settings, use the `/setup` command again.",
				Inline: false,
			},
		},
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})

	if err != nil {
		log.Printf("Error sending interaction response: %v", err)
	}

	// Fetch existing Channels for the current Guild
	channels, err := s.GuildChannels(i.GuildID)

	if err != nil {
		log.Printf("Couldn't fetch guild channels: %v", err)
	}

	//Check if the log channel already exists

	for _, ch := range channels {
		if ch.Name == helper.GetEnvOrDefault("LOG_CHANNEL_NAME", "behind-the-scenes") && ch.Type == discordgo.ChannelTypeGuildText {
			log.Println("Log channel already exists. Skipping creation.")
			logChannel = ch
			//Check if the channel is in the correct category
			if ch.ParentID != ticketCategory.ID {
				s.ChannelEditComplex(logChannel.ID, &discordgo.ChannelEdit{
					Topic:    helper.GetEnvOrDefault("LOG_CHANNEL_TOPIC", "This channel is used for logging ticket events. Please do not delete or modify this channel."),
					ParentID: ticketCategory.ID,
				})
				return
			}
		}
	}

	//Create the log channel
	logChannel, err = s.GuildChannelCreate(i.GuildID, helper.GetEnvOrDefault("LOG_CHANNEL_NAME", "behind-the-scenes"), discordgo.ChannelTypeGuildText)

	if err != nil {
		log.Printf("Couldn't create the log channel: %v", err)
	}
	// Set the topic of the log Channel
	_, err = s.ChannelEditComplex(logChannel.ID, &discordgo.ChannelEdit{
		Topic:    helper.GetEnvOrDefault("LOG_CHANNEL_TOPIC", "This channel is used for logging ticket events. Please do not delete or modify this channel."),
		ParentID: ticketCategory.ID,
	})

	if err != nil {
		log.Printf("Couldn't set the topic of the log channel: %v", err)
	}

	s.GuildChannelsReorder(i.GuildID, []*discordgo.Channel{
		{
			ID:       logChannel.ID,
			ParentID: ticketCategory.ID,
		},
	})
	// Save the ticket channel and log channel to the database
	// Get the guild Info
	guild, err := s.State.Guild(i.GuildID)

	if err != nil {
		slog.Error("Failed to get guild from state: " + err.Error())
		return
	}
	// Save the information to the database
	err = helper.SetTicketChannel(i.GuildID, ticketCategory.ID, logChannel.ID, guild.Name)

	if err != nil {
		slog.Error("Failed to save ticket and log channel to the database: " + err.Error())
	}
	// Send a message to the log channel to indicate that it has been set up
	_, err = s.ChannelMessageSendEmbed(logChannel.ID, &discordgo.MessageEmbed{
		Title: "🎉 Log Channel Initialized! 📜✨",
		Description: `
				Heyo! 👋 This channel is where **all the logging magic happens** 🎟️🔮  
				Every event, every update, all logged right here! 📝📊  

				⚠️ Please **don’t yeet** 🚫🗑️ or **tinker** 🛠️ with this channel —  
				our log goblins 🧙‍♂️ are watching... 👀
				`,
		Color: 0x00ADEF, // nice blue, change if you want

	})

	if err != nil {
		log.Printf("Couldn't send message to the log channel: %v", err)
	}
}

func HandleLogChannelCancel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Final Step - Setup Complete
	embed := &discordgo.MessageEmbed{
		Title:       "🎫 Ticket System Setup",
		Description: "The bot will now finalize the ticket system configuration.",
		Color:       0x5865F2, // Discord blurple color
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "📋 Instructions",
				Value:  "The Setup is complete! You can now start using the ticket system with the `/ticket create` command.\n\nIf you need to change any settings, use the `/setup` command again.",
				Inline: false,
			},
		},
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})

	if err != nil {
		log.Printf("Error while sending Response: %v", err)
	}

}
