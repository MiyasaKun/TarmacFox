package com.future.frame.commands.ping;

import net.dv8tion.jda.api.events.interaction.command.SlashCommandInteractionEvent;

public class PingCommand {

    public void execute(SlashCommandInteractionEvent event) {
       event.getHook().sendMessage("Bot is alive and kicking \uD83C\uDFD3").complete();
    }
}
