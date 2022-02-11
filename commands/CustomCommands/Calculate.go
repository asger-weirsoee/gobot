package CustomCommands

import (
	"GOMusicBot/CustomTypes"
	"github.com/bwmarrin/discordgo"
	"strconv"
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
			if len(args) != 2 {
				_, err := s.ChannelMessageSend(m.ChannelID, "Please provide an expression to calculate.")
				if err != nil {
					println("[WARNING] could not send message")
				}
				return
			}

			result := Calculate(args)
			res := "Result: " + strconv.FormatInt(result, 10)

			_, err := s.ChannelMessageSend(m.ChannelID, res)
			if err != nil {
				println("[WARNING] could not send message")
			}
		},
	}
}

func Calculate(args []string) int64 {
	// Add the first and the second parameter index of args together
	// Convert to integer
	// Return the result
	var int1 int
	var int2 int
	var result int

	int1, _ = strconv.Atoi(args[0])
	int2, _ = strconv.Atoi(args[1])
	result = int1 + int2

	return int64(result)
}
