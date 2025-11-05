package com.future.frame.entity;

import lombok.Data;
import lombok.Getter;
import lombok.Setter;

@Data
@Getter
@Setter
public class TicketEntity {


    private int id;
    private String guildId;
    private String channelName;
    private String channelId;
    private String roleId;
    private String roleName;
    private String categoryId;
    private String categoryName;
    private String configName;

}
