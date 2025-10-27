package com.future.frame.handler;

import com.future.frame.commands.ticket.TicketSetup;
import net.dv8tion.jda.api.events.interaction.component.ButtonInteractionEvent;
import net.dv8tion.jda.api.hooks.ListenerAdapter;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class ComponentHandler extends ListenerAdapter {
    private static final Logger logger = LoggerFactory.getLogger(ComponentHandler.class);

    public void onButtonInteraction(ButtonInteractionEvent event) {
        logger.info("ButtonInteractionEvent fired {}", event.getCustomId());
        switch (event.getCustomId()) {
            case "ticket_setup_start":
                new TicketSetup().startSetup(event);
                break;
            case "ticket_setup_back":
                new TicketSetup().setupBack(event);
                break;
            case "ticket_setup_next":
                new TicketSetup().setupNext(event);
                break;
            case "ticket_setup_finish":
                new TicketSetup().setupFinish(event);
                break;
            default:
                break;
        }
    }

}
