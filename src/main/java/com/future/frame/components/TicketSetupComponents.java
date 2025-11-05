package com.future.frame.components;

import lombok.Data;
import lombok.Getter;
import net.dv8tion.jda.api.components.actionrow.ActionRow;
import net.dv8tion.jda.api.components.buttons.Button;

@Getter
@Data
public class TicketSetupComponents {
    public final Button backButton = Button.secondary("ticket_setup_back_step", "← Back");
    public final Button nextButton = Button.secondary("ticket_setup_next_step", "→ Save & Continue ");
    public final ActionRow actionRowButtons = ActionRow.of(backButton,nextButton);
}
