package models

import (
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

type DiscordGuild struct {
	gorm.Model
	ID          uint           `gorm:"primaryKey; not null"`
	DiscordID   string         `gorm:"type:varchar(255);unique_index"`
	GuildName   string         `gorm:"type:varchar(255)"`
	GuildPrefix string         `gorm:"type:varchar(16)"`
	GuildOwner  DiscordUser    `gorm:"foreignkey:ID"`
	Members     []*DiscordUser `gorm:"many2many:discord_guild_members"`
}

func GetOrCreateGuild(gu discordgo.Guild, s *discordgo.Session) (*DiscordGuild, error) {
	var guild DiscordGuild
	if err := db.Where("discord_id = ?", gu.ID).First(&guild).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			usr, _ := s.User(gu.OwnerID)
			ov, _ := GetOrCreateUser(usr)
			guild = DiscordGuild{
				DiscordID:   gu.ID,
				GuildName:   gu.Name,
				GuildPrefix: "!",
				GuildOwner:  *ov,
			}
			if err := db.Create(&guild).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return &guild, nil
}
