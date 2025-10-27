package com.future.frame.config;

import io.github.cdimascio.dotenv.Dotenv;

public class Config {
    public final String TOKEN = getEnv("BOT_TOKEN");
    public final String DB_USER = getEnv("DB_USER");

    private String getEnv(String name){
        String variable = Dotenv.load().get(name);
        if(variable == null) {
            return "";
        }
        return variable;
    }
}
