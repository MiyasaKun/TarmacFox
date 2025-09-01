package commands

import (
	"tarmac-fox/commands/ping"
	"tarmac-fox/commands/ticket"
)

func init() {
	RegisterCommand(ping.Command, ping.Handler)
	RegisterCommand(ticket.CommandTicket, ticket.TicketHandler)
	RegisterCommand(ticket.SetupCommandTicket, ticket.SetupHandler)
	ProcessCommands()
}