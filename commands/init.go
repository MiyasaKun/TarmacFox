package commands

import "tarmac-fox/commands/ping"

func init() {
	RegisterCommand(ping.Command, ping.Handler)


	ProcessCommands()
}