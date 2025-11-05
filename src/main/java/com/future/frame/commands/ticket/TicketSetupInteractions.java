package com.future.frame.commands.ticket;

import net.dv8tion.jda.api.events.interaction.component.ButtonInteractionEvent;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public record TicketSetupInteractions(TicketSetup setup) {

    private static final Logger LOGGER = LoggerFactory.getLogger(TicketSetupInteractions.class);

    public void onSetupNext(ButtonInteractionEvent event) {
        switch (setup.currentStep) {
            case 1:
                setup.sendChannelNameConfig(event);
                break;
            case 2:
                setup.sendRoleConfig(event);
                break;
            case 3:
                setup.sendCategoryConfig(event);
                break;
            case 4:
                setup.finalizeSetup(event);
                break;
            default:
                LOGGER.warn("Invalid setup step: {}. Can't go forward", setup.currentStep);
                break;
        }
    }

    public void onSetupBack(ButtonInteractionEvent event) {
        switch(setup.currentStep) {
            case 1, 2:
                setup.startSetup(event);
                break;
            case 3:
                setup.sendChannelNameConfig(event);
                break;
            case 4:
                setup.sendRoleConfig(event);
                break;
            default:
                LOGGER.warn("Invalid setup step: {}. Can't go back", setup.currentStep);
                break;
        }
    }
}
