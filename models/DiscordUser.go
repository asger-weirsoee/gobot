package models

import (
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

type DiscordUser struct {
	gorm.Model
	ID         uint                  `gorm:"primaryKey; not null"`
	DiscordID  string                `gorm:"unique_index"`
	Nicknames  []DiscordUserNickname `gorm:"foreignkey:ID"`
	Avatar     string                `gorm:"varChar(255)"`
	Ignored    bool                  `gorm:"default:false"`
	Confirmed  bool                  `gorm:"default:false"`
	Owner      bool                  `gorm:"default:false"`
	Privileged bool                  `gorm:"default:false"`
	Messages   []DiscordMessage      `gorm:"foreignkey:ID"`
	Guilds     []*DiscordGuild       `gorm:"many2many:discord_guild_members"`
}

func GetOrCreateUser(mem *discordgo.User) (*DiscordUser, error) {
	var user DiscordUser
	if err := db.Where("discord_id = ?", mem.ID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			user = DiscordUser{
				DiscordID:  mem.ID,
				Avatar:     mem.Avatar,
				Ignored:    false,
				Confirmed:  false,
				Owner:      false,
				Privileged: false,
			}
			if err := db.Create(&user).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	_, _ = GetOrCreateNicknames(&user, mem.Username)
	return &user, nil
}

func AddUserToGuild(user *DiscordUser, guild *DiscordGuild) func() string {
	user.Guilds = append(user.Guilds, guild)
	if err := db.Save(user).Error; err != nil {
		return func() string {
			return err.Error()
		}
	}
	return nil
}

type DiscordUserNickname struct {
	gorm.Model
	ID       uint        `gorm:"primaryKey; not null"`
	User     DiscordUser `gorm:"foreignkey:ID; not null; UNIQUE_INDEX:uniquenick"`
	Nickname string      `gorm:"varChar(255); UNIQUE_INDEX:uniquenick"`
}

func GetOrCreateNicknames(user *DiscordUser, nickname string) (*DiscordUserNickname, error) {
	var nick DiscordUserNickname
	if err := db.Where("nickname = ?", nickname).First(&nick).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			nick = DiscordUserNickname{
				User:     *user,
				Nickname: nickname,
			}
			if err := db.Create(&nick).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return &nick, nil
}

type DiscordMessage struct {
	gorm.Model
	ID        uint            `gorm:"primaryKey; not null"`
	DiscordID string          `gorm:"unique_index"`
	User      *DiscordUser    `gorm:"foreignkey:ID"`
	Message   string          `gorm:"type:text"`
	Edited    bool            `gorm:"default:false"`
	Deleted   bool            `gorm:"default:false"`
	Parent    *DiscordMessage `gorm:"foreignkey:ID"`
}

func CreateNewMessage(message *discordgo.Message) (*DiscordMessage, error) {
	var msg DiscordMessage
	usr, err := GetOrCreateUser(message.Author)
	if err != nil {
		return nil, err
	}
	msg = DiscordMessage{
		DiscordID: message.ID,
		User:      usr,
		Message:   message.Content,
	}
	if err := db.Create(&msg).Error; err != nil {
		return nil, err
	}
	return &msg, nil
}

func EditedMessage(message *discordgo.Message) (*DiscordMessage, error) {
	var oldMsg DiscordMessage
	if err := db.Where("discord_id = ?", message.ID).First(&oldMsg).Error; err != nil {
		return nil, err
	}
	oldMsg.Edited = true
	if err := db.Save(&oldMsg).Error; err != nil {
		return nil, err
	}

	// Create new message where history is the old message
	newMsg, err := CreateNewMessage(message)
	if err != nil {
		return nil, err
	}
	// Set old message to be the history of the new message
	newMsg.Parent = &oldMsg

	if err := db.Save(&newMsg).Error; err != nil {
		return nil, err
	}
	return newMsg, nil
}

func DeletedMessage(message *discordgo.Message) (*DiscordMessage, error) {
	var msg DiscordMessage
	if err := db.Where("discord_id = ?", message.ID).First(&msg).Error; err != nil {
		return nil, err
	}
	msg.Deleted = true
	if err := db.Save(&msg).Error; err != nil {
		return nil, err
	}
	return &msg, nil
}
