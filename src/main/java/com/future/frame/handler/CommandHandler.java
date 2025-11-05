package com.future.frame.handler;

import com.future.frame.commands.ping.PingCommand;
import com.future.frame.commands.ticket.TicketSetup;
import com.future.frame.commands.ticket.TicketSetupInteractions;
import net.dv8tion.jda.api.events.interaction.command.SlashCommandInteractionEvent;
import net.dv8tion.jda.api.hooks.ListenerAdapter;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class CommandHandler extends ListenerAdapter {
    private static final Logger LOGGER = LoggerFactory.getLogger(CommandHandler.class);

    public void onSlashCommandInteraction(SlashCommandInteractionEvent event) {
        String command = event.getName();
        System.out.println("Received command: " + command);
        TicketSetup setup = new TicketSetup();
        switch (command) {
            case "ping":
                event.deferReply().queue();
                new PingCommand().execute(event);
                break;
            case "setup":
                setup.execute(event);
                ComponentHandler.userSetups.put(event.getUser().getId(),setup);
                break;
            default:
                event.deferReply().queue();

                try {
                    Thread.sleep(6000);
                } catch (InterruptedException exception) {
                    LOGGER.error("Error while handling unknown command: {}", exception.getMessage());
                }

                event.getHook().sendMessage("Unknown command. Please use a valid command.").complete();
                break;
        }
    }

}
