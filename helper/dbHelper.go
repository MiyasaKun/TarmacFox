package helper

import (
    "database/sql"
    "log"
    "sync"
    _ "github.com/lib/pq" // PostgreSQL driver
)

type Credentials struct {
    Username string
    Password string
}

var (
    dbInstance *sql.DB
    once       sync.Once
)

// GetDB returns the singleton database instance
func GetDB() *sql.DB {
    return dbInstance
}

// InitDB initializes the database connection (call this once at startup)
func InitDB(url string, credentials *Credentials) error {
    var err error
    once.Do(func() {
        dbInstance, err = sql.Open("postgres", "postgres://"+credentials.Username+":"+credentials.Password+"@"+url)
        if err != nil {
            log.Fatalf("error connecting to database: %v", err)
            return
        }
        log.Println("Connected to database successfully")
    })
    return err
}

// CloseDB closes the database connection (call this at shutdown)
func CloseDB() {
    if dbInstance != nil {
		log.Println("Closing database connection")
        dbInstance.Close()
    }
}