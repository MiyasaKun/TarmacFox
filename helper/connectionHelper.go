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
	CreateDatabaseTables()

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

func GetDatabaseInstance() *sql.DB {
	return DB
}

func GetCacheInstance() *redis.Client {
	return CACHE
}

func CloseDatabase() {
	if DB != nil {
		slog.Info("Closing database connection.")
		DB.Close()
	}
}
