package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	prefixSet map[string]string
)

func MigrateDiscord(path string) {
	// Append the path to the database file

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// create list of tables
	tables := []interface{}{
		&DiscordUser{},
		&DiscordGuild{},
		&DiscordUserNickname{},
		&DiscordMessage{},
	}

	// Init all tables
	for _, table := range tables {
		db.AutoMigrate(table)
	}
}

func IsPriviliged(userID string) bool {
	var user DiscordUser
	db.Where("discord_id = ?", userID).First(&user)
	return user.Privileged
}

func GetAllDiscordGuilds() []DiscordGuild {
	var guilds []DiscordGuild
	db.Find(&guilds)
	return guilds
}

func GetGuildPrefix(guildID string) string {
	var guild DiscordGuild
	db.Where("discord_id = ?", guildID).First(&guild)
	return guild.GuildPrefix
}

func GetDb(conf Conf) *gorm.DB {
	if db == nil {
		db, _ = gorm.Open(sqlite.Open(conf.GetDiscordDb()), &gorm.Config{})
	}
	return db
}

func InitCachedVars(conf Conf) {
	prefixSet = make(map[string]string)
	prefixSet[""] = conf.Prefix
	for _, guild := range GetAllDiscordGuilds() {
		prefixSet[guild.DiscordID] = guild.GuildPrefix
	}
}

func GetPrefixSet() map[string]string {
	return prefixSet
}
