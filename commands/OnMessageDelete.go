package commands

import (
	"GOMusicBot/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func OnDelete(s *discordgo.Session, evt *discordgo.MessageDelete) {
	if evt.Author.ID == s.State.User.ID || evt.Author.Bot {
		return
	}
	_, err := models.DeletedMessage(evt.Message)
	if err != nil {
		fmt.Println("[ERROR] Error getting message,", err)
	}
}
