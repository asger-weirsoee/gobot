package CustomCommands

import (
	"GOMusicBot/CustomTypes"
	"GOMusicBot/models"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"strings"
)

func GetHelp(allCommands *[]CustomCommands) {
	help := CustomCommands{}
	help.Name = "help"
	help.Description = "Displays this message"
	help.Category = "Help"
	help.Usage = "help"
	help.Command = func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
		embed := models.NewEmbed()
		embed.SetTitle("**Help**")
		description := "This message provides the commands that are available for this bot.\n\n"
		setOfPrefixes := CustomTypes.NewSet(viper.GetString("prefix"), models.GetGuildPrefix(m.GuildID))
		description += "Accepted prefix: " + strings.Join(setOfPrefixes.Vals(), " | ") + "\n\n"
		// TODO: This could be done on compile time or something idk
		catagories := make(map[string][]CustomCommands)
		for _, command := range *allCommands {
			catagories[command.Category] = append(catagories[command.Category], command)
		}
		description += "**Help:**\n"
		for _, command := range catagories["Help"] {
			description += "`" + strings.Join(setOfPrefixes.Vals(), "|") + strings.Join(command.GetAliases().Vals(), ", ") + "` - " + command.Description + "\n\n"
		}
		for key, value := range catagories {
			if key == "Help" {
				continue
			}
			description += "**" + key + ":**\n"
			for _, command := range value {
				description += "`" + strings.Join(setOfPrefixes.Vals(), "|") + strings.Join(command.GetAliases().Vals(), ", ") + "` - " + command.Description + "\n"
			}
			description += "\n"
		}

		embed.SetDescription(description)
		footer := models.NewEmbedFooter()
		footer.SetText("This bot was created by Asger Weirsøe")
		footer.SetIconURL("https://cdn.discordapp.com/avatars/143662841809862656/20626ad2527fe391bd6482bdd22ac74f.png?size=128")
		embed.SetFooter(*footer)
		embed.SetColor(0x00ff00)
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		if err != nil {
			println(err.Error())
		}
	}
	*allCommands = append(*allCommands, help)
}
