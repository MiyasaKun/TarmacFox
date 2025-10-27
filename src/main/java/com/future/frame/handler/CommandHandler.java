package com.future.frame.handler;

import com.future.frame.commands.ping.PingCommand;
import com.future.frame.commands.ticket.TicketSetup;
import net.dv8tion.jda.api.events.interaction.command.SlashCommandInteractionEvent;
import net.dv8tion.jda.api.hooks.ListenerAdapter;

public class CommandHandler extends ListenerAdapter {
    public void onSlashCommandInteraction(SlashCommandInteractionEvent event){
        String command = event.getName();
        System.out.println("Received command: " + command);
        switch (command) {
            case "ping":
                event.deferReply().queue();
                new PingCommand().execute(event);
                break;
            case "setup":
               new TicketSetup().execute(event,false);
               break;
            default:
                event.reply("Unknown command!").queue();
                break;
        }
    }

}
