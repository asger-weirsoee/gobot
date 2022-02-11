package CustomCommands

import (
	"GOMusicBot/CustomTypes"
	"github.com/bwmarrin/discordgo"
)

func GetCalculate() CustomCommands {
	return CustomCommands{
		Name: "calculate",
		Description: "Calculate the given expression.\n" +
			"Example: `calculate 2 + 2`",
		Category:  "Maths",
		Aliases:   CustomTypes.NewSet("add"),
		GuildOnly: true,
		Command: func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
			if len(args) == 0 {
				s.ChannelMessageSend(m.ChannelID, "Please provide an expression to calculate.")
				return
			}

			result, err := Calculate(args)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
				return
			}

			s.ChannelMessageSend(m.ChannelID, "Result: "+result)
		},
	}
}

func Calculate(args []string) (interface{}, interface{}) {
	return nil, nil
}
