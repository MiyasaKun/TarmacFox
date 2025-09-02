package ticket

import (
	"github.com/bwmarrin/discordgo"
)


func HandleSetup(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Check User has the PermissionsLevel to use the Commands

	if i.Member.Permissions != discordgo.PermissionManageGuild {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You do not have permission to use this command.",
			},
		})

	}

}