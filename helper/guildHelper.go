package helper

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func GetAllGuildCategories(s *discordgo.Session, guildID string) ([]*discordgo.Channel, error) {
	var categories []*discordgo.Channel
	var err error
	channels, err := s.GuildChannels(guildID)
	if err != nil {
		return nil, err
	}
	for _, ch := range channels {
		if ch.Type == discordgo.ChannelTypeGuildCategory {
			categories = append(categories, ch)
		}
	}
	return categories, nil
}

func getAllGuildRoles(s *discordgo.Session, guildID string) ([]*discordgo.Role, error) {
	var roles []*discordgo.Role
	var err error
	roles, err = s.GuildRoles(guildID)

	if err != nil {
		return nil, err
	}
	return roles, nil
}

func GetAllGuildAdminRoles(s *discordgo.Session, guildID string) ([]*discordgo.Role, error) {
	var roles []*discordgo.Role
	var err error
	var adminRoles []*discordgo.Role

	roles,err = getAllGuildRoles(s, guildID)

	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	for _,role := range roles {
		if(role.Permissions == discordgo.PermissionAdministrator) {
			adminRoles = append(adminRoles, role)
		}
	}
	log.Println(adminRoles)
	return adminRoles,nil
}

