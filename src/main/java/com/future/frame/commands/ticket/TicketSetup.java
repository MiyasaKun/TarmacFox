package com.future.frame.commands.ticket;

import com.future.frame.components.TicketSetupComponents;
import com.future.frame.entity.TicketEntity;
import net.dv8tion.jda.api.EmbedBuilder;
import net.dv8tion.jda.api.components.actionrow.ActionRow;
import net.dv8tion.jda.api.components.buttons.Button;
import net.dv8tion.jda.api.components.label.Label;
import net.dv8tion.jda.api.components.selections.EntitySelectMenu;
import net.dv8tion.jda.api.components.textinput.TextInput;
import net.dv8tion.jda.api.components.textinput.TextInputStyle;
import net.dv8tion.jda.api.entities.Guild;
import net.dv8tion.jda.api.entities.IMentionable;
import net.dv8tion.jda.api.entities.channel.Channel;
import net.dv8tion.jda.api.entities.channel.ChannelType;
import net.dv8tion.jda.api.events.interaction.ModalInteractionEvent;
import net.dv8tion.jda.api.events.interaction.command.SlashCommandInteractionEvent;
import net.dv8tion.jda.api.events.interaction.component.ButtonInteractionEvent;
import net.dv8tion.jda.api.events.interaction.component.EntitySelectInteractionEvent;
import net.dv8tion.jda.api.modals.Modal;

import java.awt.*;
import java.util.List;
import java.util.Objects;


public class TicketSetup {

    private final TicketEntity ticketEntity = new TicketEntity();
    private String configName = "`Unknown`";
    private String channelName = "create-ticket";
    private final TicketSetupComponents components = new TicketSetupComponents();
    protected final List<Integer> steps = List.of(
            1, // Configuration Name
            2, // Ticket Channel
            3, // Category Selection
            4, // Role Selection
            5  // Confirmation
    );
    public int currentStep;

    public void execute(SlashCommandInteractionEvent event) {
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

        event.replyEmbeds(builder.build()).addComponents(actionRow).complete();
    }

    public void startSetup(ButtonInteractionEvent event) {
        this.currentStep = steps.getFirst();

        Button configNameButton = Button.primary("ticket_setup_config_name", "Set Configuration Name");
        EmbedBuilder configNameEmbed = new EmbedBuilder();
        configNameEmbed.setTitle("Configuration Name");
        configNameEmbed.addField(configName, "", true);

        ActionRow actionRow1 = ActionRow.of(
                configNameButton
        );

        EmbedBuilder builder = new EmbedBuilder();
        builder.setTitle("Step 1: Configuration Name");
        builder.setDescription("Please provide a name for this ticket configuration. This name will help you identify this setup later on.\n\n" +
                "**Click the button below to set the configuration name.**");
        builder.setColor(0x1F8B4C);
        // uwu by Tjark

        event.editMessageEmbeds(builder.build(), configNameEmbed.build()).setComponents(actionRow1, components.actionRowButtons).complete();
    }

    public void configureNameButtonClick(ButtonInteractionEvent event) {
        TextInput textInput = TextInput.create("ticket_setup_config_name_input", TextInputStyle.SHORT)
                .setPlaceholder("Enter configuration name").build();

        Label label = Label.of("Configuration Name", textInput);

        event.replyModal(Modal.create("ticket_setup_config_name_modal", "Set Config Name").addComponents(label).build()).complete();
    }

    public void configureChannelNameButtonClick(ButtonInteractionEvent event) {
        TextInput textInput = TextInput.create("ticket_setup_channel_name_input", TextInputStyle.SHORT)
                .setPlaceholder("Enter ticket channel name").build();

        Label label = Label.of("Ticket Channel Name", textInput);

        event.replyModal(Modal.create("ticket_setup_channel_name_modal", "Set Ticket Channel Name").addComponents(label).build()).complete();
    }

    public void handleChannelNameModal(ModalInteractionEvent event) {
        this.channelName = Objects.requireNonNull(event.getValue("ticket_setup_channel_name_input")).getAsString();

        ticketEntity.setChannelName(this.channelName);

        EmbedBuilder embed = new EmbedBuilder();
        embed.setTitle("Step 2/5: Ticket Channel Configuration");
        embed.setDescription("**Configure the name of the ticket channel**. This Channel is the Channel where the user can create new Tickets ");
        embed.setColor(Color.GREEN);

        Button configTicketChannelName = Button.secondary("ticket_setup_channel_name", "Set Channel Name");
        ActionRow actionRow = ActionRow.of(configTicketChannelName);

        EmbedBuilder embedChannelName = new EmbedBuilder();
        embedChannelName.setTitle("Ticket Channel Name");
        embedChannelName.setDescription(channelName);

        event.editMessageEmbeds(embed.build(), embedChannelName.build()).setComponents(actionRow, components.actionRowButtons).complete();
    }

    public void handleConfigNameModal(ModalInteractionEvent event) {
        this.configName = Objects.requireNonNull(event.getValue("ticket_setup_config_name_input")).getAsString();

        ticketEntity.setConfigName(this.configName);

        Button configNameButton = Button.primary("ticket_setup_config_name", "Set Configuration Name");
        EmbedBuilder configNameEmbed = new EmbedBuilder();
        configNameEmbed.setTitle("Configuration Name");
        configNameEmbed.addField(configName, "", true);

        ActionRow actionRow1 = ActionRow.of(
                configNameButton
        );

        EmbedBuilder builder = new EmbedBuilder();
        builder.setTitle("Step 1/5: Configuration Name");
        builder.setDescription("Please provide a name for this ticket configuration. This name will help you identify this setup later on.\n\n" +
                "**Click the button below to set the configuration name.**");
        builder.setColor(0x1F8B4C);

        // uwu by Tjark

        event.editMessageEmbeds(builder.build(), configNameEmbed.build()).setComponents(actionRow1, components.actionRowButtons).complete();
    }

    public void sendChannelNameConfig(ButtonInteractionEvent event) {
        this.currentStep = steps.get(1);

        EmbedBuilder embed = new EmbedBuilder();
        embed.setTitle("Step 2/5: Ticket Channel Configuration");
        embed.setDescription("**Configure the name of the ticket channel**. This Channel is the Channel where the user can create new Tickets ");
        embed.setColor(Color.GREEN);

        Button configTicketChannelName = Button.secondary("ticket_setup_channel_name", "Set Channel Name");
        ActionRow actionRow = ActionRow.of(configTicketChannelName);

        EmbedBuilder embedChannelName = new EmbedBuilder();
        embedChannelName.setTitle("Ticket Channel Name");
        embedChannelName.setDescription(channelName);

        event.editMessageEmbeds(embed.build(), embedChannelName.build()).setComponents(actionRow, components.actionRowButtons).complete();
    }

    public void sendRoleConfig(ButtonInteractionEvent event) {
        this.currentStep = steps.get(2);
        String roleName = "`Not Selected`";

        EmbedBuilder builder = new EmbedBuilder();
        builder.setTitle("Step 3/5: Role Configuration");
        builder.setDescription("**Configure the role that will have access to the tickets and also work on them.**\n\n" +
                "Currently u can only specify a single role. More options will be added in future updates.");
        builder.setColor(Color.GREEN);

        EmbedBuilder roleNameEmbed = new EmbedBuilder();
        roleNameEmbed.setTitle("Selected Role");
        roleNameEmbed.setDescription(roleName);

        EntitySelectMenu roleSelectMenu = EntitySelectMenu.create("ticket_setup_role_select", EntitySelectMenu.SelectTarget.ROLE)
                .setPlaceholder("Select Role with Ticket Access")
                .setMinValues(1)
                .setMaxValues(1)
                .build();
        ActionRow actionRow = ActionRow.of(
                roleSelectMenu
        );

        event.editMessageEmbeds(builder.build(), roleNameEmbed.build()).setComponents(actionRow, components.actionRowButtons).complete();
    }

    public void handleRoleSelectMenu(EntitySelectInteractionEvent event) {
        List<IMentionable> roleList = event.getValues();
        String roleId = roleList.isEmpty() ? "" : roleList.getFirst().getId();
        ticketEntity.setRoleId(roleId);

        Guild guild = event.getGuild();
        String roleName;

        if (guild != null) {
            roleName = Objects.requireNonNull(guild.getRoleById(roleId)).getName();
            ticketEntity.setRoleName(roleName);
        } else {
            roleName = "Unknown Role";
        }

        this.currentStep = steps.get(2);

        EmbedBuilder builder = new EmbedBuilder();
        builder.setTitle("Step 3/5: Role Configuration");
        builder.setDescription("**Configure the role that will have access to the tickets and also work on them.**\n\n" +
                "Currently u can only specify a single role. More options will be added in future updates.");
        builder.setColor(Color.GREEN);

        EmbedBuilder roleNameEmbed = new EmbedBuilder();
        roleNameEmbed.setTitle("Selected Role");
        roleNameEmbed.setDescription(roleName);

        EntitySelectMenu roleSelectMenu = EntitySelectMenu.create("ticket_setup_role_select", EntitySelectMenu.SelectTarget.ROLE)
                .setPlaceholder("Select Role with Ticket Access")
                .setMinValues(1)
                .setMaxValues(1)
                .build();
        ActionRow actionRow = ActionRow.of(
                roleSelectMenu
        );

        event.editMessageEmbeds(builder.build(), roleNameEmbed.build()).setComponents(actionRow, components.actionRowButtons).complete();
    }

    public void sendCategoryConfig(ButtonInteractionEvent event) {
        this.currentStep = steps.get(3);
        String categoryName = "`Not Selected`";

        EmbedBuilder builder = new EmbedBuilder();
        builder.setTitle("Step 4/5: Category Configuration");
        builder.setDescription("**Configure the category where the ticket channels will be created.**\n\n" +
                "Select a category from the menu below.");
        builder.setColor(Color.GREEN);

        EmbedBuilder categoryEmbed = new EmbedBuilder();
        categoryEmbed.setTitle("Selected Category");
        categoryEmbed.setDescription(categoryName);


        EntitySelectMenu categorySelectMenu = EntitySelectMenu.create("ticket_setup_select_category", EntitySelectMenu.SelectTarget.CHANNEL).setPlaceholder("Choose the Category").setChannelTypes(ChannelType.CATEGORY).build();

        ActionRow actionRow = ActionRow.of(
                categorySelectMenu
        );

        event.editMessageEmbeds(builder.build(),categoryEmbed.build()).setComponents(actionRow,components.actionRowButtons).complete();
    }
    public void handleCategorySelectMenu(EntitySelectInteractionEvent event) {
        List<IMentionable> categoryList = event.getValues();
        String categoryId = categoryList.isEmpty() ? "" : categoryList.getFirst().getId();
        ticketEntity.setCategoryId(categoryId);

        Guild guild = event.getGuild();
        String categoryName;

        if (guild != null) {
            categoryName = Objects.requireNonNull(guild.getCategoryById(categoryId)).getName();
            ticketEntity.setCategoryName(categoryName);
        } else {
            categoryName = "Unknown Category";
        }

        this.currentStep = steps.get(3);

        EmbedBuilder builder = new EmbedBuilder();
        builder.setTitle("Step 4/5: Category Configuration");
        builder.setDescription("**Configure the category where the ticket channels will be created.**\n\n" +
                "Select a category from the menu below.");
        builder.setColor(Color.GREEN);

        EmbedBuilder categoryEmbed = new EmbedBuilder();
        categoryEmbed.setTitle("Selected Category");
        categoryEmbed.setDescription(categoryName);

        EntitySelectMenu categorySelectMenu = EntitySelectMenu.create("ticket_setup_select_category", EntitySelectMenu.SelectTarget.CHANNEL).setPlaceholder("Choose the Category").setChannelTypes(ChannelType.CATEGORY).build();

        ActionRow actionRow = ActionRow.of(
                categorySelectMenu
        );

        event.editMessageEmbeds(builder.build(), categoryEmbed.build()).setComponents(actionRow, components.actionRowButtons).complete();
    }
    public void finalizeSetup(ButtonInteractionEvent event) {
        this.currentStep = steps.getLast();

        EmbedBuilder builder = new EmbedBuilder();
        builder.setTitle("Step 5/5: Setup Complete");
        builder.setDescription("You have successfully completed the ticket system setup. Your configuration has been saved and is now active.\n\n" +
                "**Thank you for using our ticket system!**");
        builder.setColor(Color.GREEN);

        event.editMessageEmbeds(builder.build()).setComponents().complete();

        //TODO Create the Channel and manage Permissions

        Guild guild = event.getGuild();
        if(guild == null) {
         return;
        }

        ticketEntity.setGuildId(guild.getId());
        // Further processing like saving the ticketEntity to a database can be done here.

       Channel ticketChannel =  guild.createTextChannel(ticketEntity.getChannelName()).complete();
       ticketEntity.setChannelId(ticketChannel.getId());

       
    }
}