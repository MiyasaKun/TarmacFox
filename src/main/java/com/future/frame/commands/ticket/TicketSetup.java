package com.future.frame.commands.ticket;

import net.dv8tion.jda.api.EmbedBuilder;
import net.dv8tion.jda.api.components.actionrow.ActionRow;
import net.dv8tion.jda.api.components.buttons.Button;
import net.dv8tion.jda.api.events.interaction.command.SlashCommandInteractionEvent;
import net.dv8tion.jda.api.events.interaction.component.ButtonInteractionEvent;

import javax.swing.*;
import java.util.List;

public class TicketSetup {
    Button backButton = Button.secondary("ticket_setup_back_step_1", "← Back");
    Button nextButton = Button.secondary("ticket_setup_next_step_1", "→ Save & Continue ");

    private final List<Integer> steps = List.of(
            1, // Configuration Name
            2, // Ticket Channel
            3, // Category Selection
            4, // Role Selection
            5  // Confirmation
    );
    private int currentStep = 0;

    public void execute(SlashCommandInteractionEvent event, boolean deferred) {
        EmbedBuilder builder = new EmbedBuilder();
        builder.setTitle("Welcome to the Ticket System Setup");
        builder.setDescription("This Setup will guide you through the process of Configuration\n" +
                "\n **Please follow the steps carefully to ensure proper setup of the ticket system **." +
                "\n\n*If you need assistance, feel free to reach out to the support team.*");
        builder.setColor(0x1F8B4C);
        Button button = Button.secondary("ticket_setup_start", "Create a new Configuration");
        ActionRow actionRow = ActionRow.of(
                button
        );
        if (deferred) {
            event.getHook().sendMessageEmbeds(builder.build()).complete();
            return;
        }
        event.replyEmbeds(builder.build()).addComponents(actionRow).complete();
    }

    public void startSetup(ButtonInteractionEvent event) {
        currentStep = steps.getFirst();

        Button configNameButton = Button.primary("ticket_setup_config_name", "Set Configuration Name");
        EmbedBuilder configNameEmbed = new EmbedBuilder();
        configNameEmbed.setTitle("Configuration Name: `Unknown`");

        ActionRow actionRow1 = ActionRow.of(
                configNameButton
        );

        ActionRow actionRow = ActionRow.of(
            backButton, nextButton
        );

        EmbedBuilder builder = new EmbedBuilder();
        builder.setTitle("Step 1: Configuration Name");
        builder.setDescription("Please provide a name for this ticket configuration. This name will help you identify this setup later on.\n\n" +
                "**Click the button below to set the configuration name.**");
        builder.setColor(0x1F8B4C);


        event.editMessageEmbeds(builder.build(),configNameEmbed.build()).setComponents(actionRow1,actionRow).queue();
    }

    public void setupBack(ButtonInteractionEvent event) {

    }

    public void setupNext(ButtonInteractionEvent event) {

    }

    public void setupFinish(ButtonInteractionEvent event) {

    }
}
