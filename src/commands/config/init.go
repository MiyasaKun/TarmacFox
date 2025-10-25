package config

import (
	"tarmac-fox/src/commands/ping"
	"tarmac-fox/src/commands/ticket"
)

func init() {
	RegisterCommand(ping.Command, ping.Handler)
	RegisterCommand(ticket.SetupCommandTicket, ticket.SetupHandler)
	ProcessCommands()
}

