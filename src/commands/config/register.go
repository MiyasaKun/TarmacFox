package config

import (
	"log"
	"strings"
	"tarmac-fox/src/commands/ticket/handler"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler struct {
    Command *discordgo.ApplicationCommand
    Handler func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var (
    Commands        []*discordgo.ApplicationCommand
    CommandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
    registeredCommands []CommandHandler
)

// RegisterCommand registers a command with its handler
func RegisterCommand(cmd *discordgo.ApplicationCommand, handler func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
    registeredCommands = append(registeredCommands, CommandHandler{
        Command: cmd,
        Handler: handler,
    })
}

// ProcessCommands validates commands have handlers and prepares them for registration
func ProcessCommands() {
    Commands = nil // Clear existing commands
    CommandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
    
    for _, cmdHandler := range registeredCommands {
        if cmdHandler.Handler == nil {
            log.Printf("Warning: Command '%s' has no handler, skipping registration", cmdHandler.Command.Name)
            continue
        }
        
        Commands = append(Commands, cmdHandler.Command)
        CommandHandlers[cmdHandler.Command.Name] = cmdHandler.Handler
        log.Printf("Registered command: %s", cmdHandler.Command.Name)
    }
}

// GetHandler returns the handler for a given command name
func GetHandler(commandName string) (func(s *discordgo.Session, i *discordgo.InteractionCreate), bool) {
    handler, exists := CommandHandlers[commandName]
    return handler, exists
}

// SetupHandlers registers all command handlers with the Discord session
func SetupHandlers(s *discordgo.Session) {
    // Handle slash commands
    s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
        if i.Type == discordgo.InteractionApplicationCommand {
            if i.ApplicationCommandData().Name == "" {
                return
            }
            
            if handler, exists := GetHandler(i.ApplicationCommandData().Name); exists {
                handler(s, i)
            } else {
                log.Printf("No handler found for command: %s", i.ApplicationCommandData().Name)
            }
        }
    })
    
    // Handle component interactions (buttons, select menus, etc.)
    s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
        if i.Type == discordgo.InteractionMessageComponent {
            handleComponentInteraction(s, i)
        }
    })
}

// Handle all component interactions
func handleComponentInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
    customID := i.MessageComponentData().CustomID
    
    // Route to appropriate component handler based on CustomID
    switch {
    case strings.Split(customID, "_")[0] == "ticket":
        handler.HandleTicketComponentInteraction(s,i, customID)
    default:
        log.Printf("No handler found for component: %s", customID)
    }
}