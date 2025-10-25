package ping

import (
	"tarmac-fox/src/helper"

	"github.com/bwmarrin/discordgo"
)

var Command = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Ping the bot to check if it's online",
}

func Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Bot is alive and kicking 🏓",
		},
	})
	helper.GetAllGuildRoles(s, i.GuildID)
}