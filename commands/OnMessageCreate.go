package commands

import (
	"GOMusicBot/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// MessageCreate This function will be called (due to AddHandler above) every time a
// new  message is created on any channel that the authenticated bot has
// access to.

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself.
	// This isn't required in this specific example but it's a good
	// practice.
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	_, err := models.CreateNewMessage(m.Message)
	if err != nil {
		fmt.Println("[ERROR] Error creating message,", err)
	}
	if k, ok := models.GetPrefixSet()[m.GuildID]; ok {
		if strings.HasPrefix(m.Content, k) {

			command := strings.TrimPrefix(m.Content, k)
			command = strings.Split(command, " ")[0]
			if c, ok := listOfDictCommands[command]; ok {
				println("[INFO] Accepted command: " + k + m.Content[len(k):])
				c.RunCommand(s, m)
			}
		}
	}
}
