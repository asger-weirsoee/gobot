package commands

import "GOMusicBot/commands/CustomCommands"

var listOfDictCommands = LoadShardsFromCustomCommands()

func LoadShardsFromCustomCommands() map[string]CustomCommands.CustomCommands {
	listOfCommands := []CustomCommands.CustomCommands{
		CustomCommands.GetPing(),
		CustomCommands.GetCalculate(),
	}
	CustomCommands.GetHelp(&listOfCommands)
	var dictList = make(map[string]CustomCommands.CustomCommands)
	for _, command := range listOfCommands {
		for _, alias := range command.GetAliases().Vals() {
			dictList[alias] = command
		}
	}
	return dictList
}
