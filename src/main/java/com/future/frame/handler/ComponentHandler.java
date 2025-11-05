package com.future.frame.handler;

import com.future.frame.commands.ticket.TicketSetup;
import com.future.frame.commands.ticket.TicketSetupInteractions;
import net.dv8tion.jda.api.events.interaction.ModalInteractionEvent;
import net.dv8tion.jda.api.events.interaction.component.ButtonInteractionEvent;
import net.dv8tion.jda.api.events.interaction.component.EntitySelectInteractionEvent;
import net.dv8tion.jda.api.hooks.ListenerAdapter;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class ComponentHandler extends ListenerAdapter {
    private static final Logger logger = LoggerFactory.getLogger(ComponentHandler.class);
    public static final Map<String, TicketSetup> userSetups = new ConcurrentHashMap<>();

    public void onButtonInteraction(ButtonInteractionEvent event) {
        logger.info("ButtonInteractionEvent fired {}", event.getCustomId());

        switch (event.getCustomId()) {
            case "ticket_setup_start":
                userSetups.get(event.getUser().getId()).startSetup(event);
                break;
            case "ticket_setup_next_step":
                new TicketSetupInteractions(userSetups.get(event.getUser().getId())).onSetupNext(event);
                break;
            case "ticket_setup_back_step":
                new TicketSetupInteractions(userSetups.get(event.getUser().getId())).onSetupBack(event);
                break;
            case "ticket_setup_config_name":
                userSetups.get(event.getUser().getId()).configureNameButtonClick(event);
                break;
            case "ticket_setup_channel_name":
                userSetups.get(event.getUser().getId()).configureChannelNameButtonClick(event);
                break;
            default:
                break;
        }
    }

    public void onModalInteraction(ModalInteractionEvent event) {
        logger.info("ModalInteractionEvent fired {}", event.getModalId());
        switch (event.getModalId()) {
            case "ticket_setup_config_name_modal":
                userSetups.get(event.getUser().getId()).handleConfigNameModal(event);
                break;
            case "ticket_setup_channel_name_modal":
                userSetups.get(event.getUser().getId()).handleChannelNameModal(event);
                break;
            default:
                break;
        }
    }

    public void onEntitySelectInteraction(EntitySelectInteractionEvent event) {
        logger.info("EntitySelectInteractionEvent fired {}", event.getSelectMenu().getCustomId());

        switch(event.getSelectMenu().getCustomId()) {
            case "ticket_setup_select_category":
                userSetups.get(event.getUser().getId()).handleCategorySelectMenu(event);
                break;
            default:
                userSetups.get(event.getUser().getId()).handleRoleSelectMenu(event);
                break;
        }
    }

}
