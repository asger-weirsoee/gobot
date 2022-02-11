package commands

import (
	"GOMusicBot/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func OnMessageEdit(s *discordgo.Session, evt *discordgo.MessageUpdate) {
	if evt.Author.ID == s.State.User.ID || evt.Author.Bot {
		return
	}
	_, err := models.EditedMessage(evt.Message)
	if err != nil {
		fmt.Println("[ERROR] Error getting message,", err)
	}
}
