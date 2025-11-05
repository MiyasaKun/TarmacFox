package com.future.frame.config;

import io.github.cdimascio.dotenv.Dotenv;

public class Config {
    public static final String TOKEN = getEnv("BOT_TOKEN");
    public static final String DB_USER = getEnv("DB_USER");
    public static final String DB_PASSWORD = getEnv("DB_PASSWORD");

    private static String getEnv(String name){
        String variable = Dotenv.load().get(name);
        if(variable == null) {
            return "";
        }
        return variable;
    }
}
