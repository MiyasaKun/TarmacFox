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