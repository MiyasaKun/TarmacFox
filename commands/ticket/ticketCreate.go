package ticket

import (
	"log"
	"tarmac-fox/helper"

	"github.com/bwmarrin/discordgo"
)


func HandleCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ch, err := createCategoryForTicket(s, i.GuildID)
	if err != nil {
		log.Printf("Error creating ticket category: %v", err)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to create ticket category.",
			},
		})
		return
	}
	ticketChannel, err := s.GuildChannelCreate(i.GuildID, "ticket-"+i.ApplicationCommandData().Options[0].Options[0].StringValue(), discordgo.ChannelTypeGuildText)
	if err != nil {
		log.Printf("Error creating ticket channel: %v", err)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to create ticket channel.",
			},
		})
		return
	}

	// Move the ticket channel under the Tickets category
	_, err = s.ChannelEdit(ticketChannel.ID, &discordgo.ChannelEdit{
		ParentID: ch.ID,
	})
	helper.SendTicketToDB(helper.Ticket{
		Title:     ticketChannel.Name,
		ChannelID: ticketChannel.ID,
		GuildID:   i.GuildID,
		CreatedBy: i.Member.User.ID,
	})
	if err != nil {
		log.Printf("Error assigning ticket channel to category: %v", err)
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Ticket created successfully!",
		},
	})
}


func createCategoryForTicket(s *discordgo.Session, guildID string) (c *discordgo.Channel, err error) {
	var categories []*discordgo.Channel
	var cat bool
	var ch *discordgo.Channel

	var ticketCategory *discordgo.Channel = &discordgo.Channel{
		Name: "Tickets",
		Type: discordgo.ChannelTypeGuildCategory,
	}

	categories, err = s.GuildChannels(guildID)

	if err != nil {
		log.Printf("Error getting channels: %v", err)
	}

	for i := 0; i < len(categories); i++ {
		if categories[i].Type == discordgo.ChannelTypeGuildCategory && categories[i].Name == "Tickets" {
			cat = true
			ch = categories[i]
			log.Println("Ticket category already exists.")
		}
	}
	if(cat){
		return ch, nil
	}

	c,err = s.GuildChannelCreate(guildID, ticketCategory.Name, ticketCategory.Type)

	if err != nil {
		log.Printf("Error creating ticket category: %v", err)
		return nil,err
	}
	log.Println("Ticket category created successfully.")

	return ch,nil
}
