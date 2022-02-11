package models

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Conf struct {
	Token                string
	Prefix               string
	Owner                string
	Authors              []string
	DatabaseDir          string
	DiscordDb            string
	LogDir               string
	DiscordLogFile       string
	DiscordLogLevel      string
	DiscordLogMaxSizeMb  int
	DiscordLogMaxAgeDays int
	DiscordLogMaxBackups int
}

func (c Conf) GetDiscordDb() string {
	// If directory DatabaseDir does not exist, create it
	if _, err := os.Stat(c.DatabaseDir); os.IsNotExist(err) {
		err := os.Mkdir(c.DatabaseDir, os.ModePerm)
		if err != nil {
			panic(err.Error())
		}
	}
	return filepath.Join(filepath.Join(c.DatabaseDir, c.DiscordDb))
}

func GenerateDefaultConfiguration() {
	viper.SetDefault("Token", "Insert Token Here")
	viper.SetDefault("Prefix", "!")
	viper.SetDefault("Owner", "Insert Owner ID Here")
	viper.SetDefault("Authors", []int64{143662841809862656})
	viper.SetDefault("DatabaseDir", "./data")
	viper.SetDefault("DiscordDb", "discord.DB")
	viper.SetDefault("LogDir", "./logs")
	viper.SetDefault("DiscordLogFile", "bot.log")
	viper.SetDefault("DiscordLogLevel", "info")
	viper.SetDefault("DiscordLogMaxSizeMb", 10)
	viper.SetDefault("DiscordLogMaxAgeDays", 7)
	viper.SetDefault("DiscordLogMaxBackups", 3)
}

// Load configuration
func LoadConfiguration(root string) Conf {
	viper.SetConfigName("config")
	viper.AddConfigPath(root)
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		println("Error reading config file: ", err.Error())
		err = nil
	}
	conf := Conf{}
	GenerateDefaultConfiguration()
	err = viper.WriteConfig()
	if err != nil {
		err = nil
		err = viper.SafeWriteConfig()
		if err != nil {
			println("Error writing config file: ", err.Error())
		}
		err = nil
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		println("Error unmarshalling config file: ", err.Error())
	}
	return conf
}
