package CustomCommands

import (
	"GOMusicBot/CustomTypes"
	"github.com/bwmarrin/discordgo"
)

func GetPing() CustomCommands {
	return CustomCommands{

		Name:        "ping",
		Description: "You write ping, bot responds with pong!",
		Usage:       "ping",
		Category:    "General",
		RequiredPermissions: []string{
			"SEND_MESSAGES",
		},
		Aliases:   CustomTypes.NewSet("beep"),
		GuildOnly: false,
		OwnerOnly: false,
		Hidden:    false,
		Enabled:   true,
		Command: func(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) {
			_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
			if err != nil {
				println(err.Error())
			}
		},
	}
}
