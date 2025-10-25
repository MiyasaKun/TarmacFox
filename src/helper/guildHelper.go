package helper

import (
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

func GetAllGuildRoles(s *discordgo.Session, guildID string) (map[string]*discordgo.Role, error) {
	var roles []*discordgo.Role
	var err error
	roles, err = s.GuildRoles(guildID)

	if err != nil {
		return nil, err
	}

	var rolesMap = make(map[string]*discordgo.Role)

	for _, role := range roles {
		rolesMap[role.ID] = role
	}

	return rolesMap, nil
}



