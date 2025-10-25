package main

import (
	"log"
	"os"
	"os/signal"
	"tarmac-fox/src/commands/config"
	"tarmac-fox/src/helper"

	// Import Discord Command Library
	"github.com/bwmarrin/discordgo"
	// Importing godotenv to automatically load environment variables from .env file
	_ "github.com/joho/godotenv/autoload"
)

var sess *discordgo.Session

var (
	RemoveCommands = true
)

func init() {
	var err error

	sess, err = discordgo.New("Bot " + helper.GetEnvOrDefault("BOT_TOKEN", ""))
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
	}

}

func main() {

	config.SetupHandlers(sess)

	err := sess.Open()
	if err != nil {
		log.Fatalf("error opening Discord session: %v", err)
	}

	for _, cmd := range config.Commands {
		_, err := sess.ApplicationCommandCreate(sess.State.User.ID, "", cmd)
		if err != nil {
			log.Printf("Cannot create '%s' command: %v", cmd.Name, err)
		}
	}

	defer sess.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	helper.CloseDatabase()
	log.Println("Gracefully shutting down.")
}
