package CustomCommands

import (
	"GOMusicBot/CustomTypes"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type CustomCommands struct {
	Name                string
	Description         string
	Usage               string
	Category            string
	Aliases             *CustomTypes.Set
	RequiredPermissions []string
	Enabled             bool
	GuildOnly           bool
	OwnerOnly           bool
	Hidden              bool
	Command             func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
}

func (c *CustomCommands) GetName() string {
	return c.Name
}

func (c *CustomCommands) GetDescription() string {
	return c.Description
}

func (c *CustomCommands) GetUsage() string {
	return c.Usage
}

func (c *CustomCommands) GetCategory() string {
	return c.Category
}

func (c *CustomCommands) GetAliases() *CustomTypes.Set {
	if c.Aliases == nil {
		c.Aliases = CustomTypes.NewSet()
	}
	c.Aliases.Add(c.Name)
	return c.Aliases
}

func (c *CustomCommands) GetRequiredPermissions() []string {
	return c.RequiredPermissions
}

func (c *CustomCommands) IsEnabled() bool {
	return c.Enabled
}

func (c *CustomCommands) IsGuildOnly() bool {
	return c.GuildOnly
}

func (c *CustomCommands) IsOwnerOnly() bool {
	return c.OwnerOnly
}

func (c *CustomCommands) IsHidden() bool {
	return c.Hidden
}

func (c *CustomCommands) RunCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	c.Command(s, m, parseArgs(m.Content))
}

func parseArgs(content string) []string {
	args := make([]string, 0)
	for _, arg := range strings.Split(content, " ") {
		if len(arg) > 0 {
			args = append(args, arg)
		}
	}
	return args[1:]
}
