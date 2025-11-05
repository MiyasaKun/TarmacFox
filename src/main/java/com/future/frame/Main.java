package com.future.frame;

import com.future.frame.config.Config;
import com.future.frame.handler.CommandHandler;
import com.future.frame.handler.ComponentHandler;
import net.dv8tion.jda.api.JDA;
import net.dv8tion.jda.api.JDABuilder;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
    private static final Logger logger = LoggerFactory.getLogger(Main.class);

    public static void main(String[] args) {

        Config config = new Config();
        JDA api = JDABuilder.createDefault(config.TOKEN).build();
        try {
            api.awaitReady();
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }

        logger.info("Bot is starting...");

        api.upsertCommand("ping", "Replies with Pong!").queue();
        api.upsertCommand("setup", "Runs the Setup for the ticket system").queue();
        api.addEventListener(new CommandHandler());
        api.addEventListener(new ComponentHandler());


        logger.info("Commands and Handler registered");
    }
}