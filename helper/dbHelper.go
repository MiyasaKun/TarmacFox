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
        dbInstance, err = sql.Open("postgres", "postgres://"+credentials.Username+":"+credentials.Password+"@"+url+"/tarmac_fox"+"?sslmode=disable")
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

func GenerateTables() {
    _, err := dbInstance.Exec(`
        CREATE TABLE IF NOT EXISTS tb_tickets (
            id SERIAL PRIMARY KEY,
            title TEXT NOT NULL,
            channel_id TEXT NOT NULL,
            guild_id TEXT NOT NULL,
            created_by TEXT NOT NULL
        )
    `)
    if err != nil {
        log.Printf("Error creating tables: %v", err)
    }
}

type Ticket struct {
    ID int
    Title string
    ChannelID string
    GuildID string
    CreatedBy string
}
func getAllTickets() ([]Ticket,error) {
    tickets,err := dbInstance.Query("SELECT * FROM tb_tickets")
   if err != nil {
       log.Printf("Error getting tickets: %v", err)
       return nil, err
   }
   defer tickets.Close()

   var result []Ticket
   for tickets.Next() {
       var ticket Ticket
       if err := tickets.Scan(&ticket.ID, &ticket.Title, &ticket.ChannelID, &ticket.GuildID, &ticket.CreatedBy); err != nil {
           log.Printf("Error scanning ticket: %v", err)
           continue
       }
       result = append(result, ticket)
   }
   return result, nil
}

func SendTicketToDB(ticket Ticket) (b bool, err error) {
    res, err := dbInstance.Exec("INSERT INTO tb_tickets (title, channel_id, guild_id, created_by) VALUES ($1, $2, $3, $4)", ticket.Title, ticket.ChannelID, ticket.GuildID, ticket.CreatedBy)

    if err != nil {
        log.Printf("Error inserting ticket: %v", err)
        return false, err
    }

    rowsAffected, err := res.RowsAffected()

    if err != nil {
        log.Printf("Error getting rows affected: %v", err)
        return false, err
    }

    return rowsAffected > 0, nil
}