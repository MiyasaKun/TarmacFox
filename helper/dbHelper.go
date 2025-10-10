package helper

func SetTicketChannel(guildID string, ticketChannel string, logChannel string, guildName string) error {

	_, err := GetDatabaseInstance().Exec("INSERT INTO guilds (guild_id, ticket_category_id, log_channel_id, guild_name) VALUES ($1, $2, $3, $4) ON CONFLICT (guild_id) DO UPDATE SET ticket_category_id = $2, log_channel_id = $3, guild_name = $4", guildID, ticketChannel, logChannel, guildName)

	return err
}
