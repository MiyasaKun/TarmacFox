package helper

import "log/slog"

func SetTicketChannel(guildID string, ticketChannel string, logChannel string, guildName string) error {

	_, err := GetDatabaseInstance().Exec("INSERT INTO guilds (guild_id, ticket_category_id, log_channel_id, guild_name) VALUES ($1, $2, $3, $4) ON CONFLICT (guild_id) DO UPDATE SET ticket_category_id = $2, log_channel_id = $3, guild_name = $4", guildID, ticketChannel, logChannel, guildName)

	return err
}

func CreateDatabaseTables() {

	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS guilds (id SERIAL PRIMARY KEY, guild_id VARCHAR(20) UNIQUE NOT NULL, ticket_category_id VARCHAR(20), log_channel_id VARCHAR(20), guild_name VARCHAR(255));")

	if err != nil {
		slog.Warn("Failed to create guilds table: " + err.Error())
	}

	slog.Info("Guild Table checked/created")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS tickets (id SERIAL PRIMARY KEY, guild_id VARCHAR(20) NOT NULL, channel_id VARCHAR(20) UNIQUE NOT NULL, user_id VARCHAR(20) NOT NULL, status VARCHAR(20) NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);")

	if err != nil {
		slog.Warn("Failed to create tickets table: " + err.Error())
	}
	slog.Info("Tickets Table checked/created")

}
