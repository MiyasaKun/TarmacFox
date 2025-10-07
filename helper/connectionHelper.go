package helper

import (
	"context"
	"database/sql"

	"golang.org/x/exp/slog"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var (
	DBHost = GetEnvOrDefault("DB_HOST", "localhost")
	DBPort = GetEnvOrDefault("DB_PORT", "5432")
	DBUser = GetEnvOrDefault("DB_USERNAME", "postgres")
	DBPass = GetEnvOrDefault("DB_PASSWORD", "postgres")
	DB     *sql.DB
	CACHE  *redis.Client
)

func init() {
	initializeDatabase()
	initializeCache()
}

func initializeDatabase() {
	// Connect to the database
	var err error
	DB, err = sql.Open("postgres", "host="+DBHost+" port="+DBPort+" user="+DBUser+" password="+DBPass+" dbname=tarmac_fox sslmode=disable")

	if err != nil {
		slog.Error("Failed to connect to the database: " + err.Error())
		return
	}
	// Create necessary tables if they don't exist
	createDatabaseTables()
	defer DB.Close()

	err = DB.Ping()

	if err != nil {
		slog.Error("Failed to ping the database: " + err.Error())
		return
	}

	slog.Info("Successfully connected to the database")

}

var ctx = context.Background()

func initializeCache() {
	// Connect to the cache

	CACHE = redis.NewClient(&redis.Options{
		Addr:     GetEnvOrDefault("REDIS_HOST", "localhost") + ":" + GetEnvOrDefault("REDIS_PORT", "6379"),
		Password: GetEnvOrDefault("REDIS_PASSWORD", ""),
		DB:       0,
	})

	_, err := CACHE.Ping(ctx).Result()

	if err != nil {
		slog.Error("Failed to connect to the cache: " + err.Error())
		return
	}

	slog.Info("Successfully connected to the cache")

}

func createDatabaseTables() {

	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS guilds (id SERIAL PRIMARY KEY, guild_id VARCHAR(20) UNIQUE NOT NULL, ticket_category_id VARCHAR(20), log_channel_id VARCHAR(20));")

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

func GetDatabaseInstance() *sql.DB {
	return DB
}

func GetCacheInstance() *redis.Client {
	return CACHE
}
