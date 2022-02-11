package main

import (
	"GOMusicBot/commands"
	"GOMusicBot/models"
	"flag"
	"fmt"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/servusdei2018/shards"
)

var (
	Mgr *shards.Manager
)

// Token - Variables used for command line parameters.
var (
	Root      string
	conf      models.Conf
	prefixSet map[string]string
	db        *gorm.DB
)

func init() {
	Root, _ = os.Getwd()

	migrate := flag.Bool("migrate", false, "Migrate database")
	createConfig := flag.Bool("create-config", false, "Only create default config file")
	flag.Parse()
	// Load Config from deafult.yml
	conf = models.LoadConfiguration(Root)
	if createConfig != nil && *createConfig {
		os.Exit(0)
	}

	if migrate != nil && *migrate {
		models.MigrateDiscord(conf.GetDiscordDb())
		os.Exit(0)
	}
	db = models.GetDb(conf)

	prefixSet = make(map[string]string)
	// Default prefix from config file
	prefixSet[""] = conf.Prefix
	// Cache all prefix for each guild from the database
	for _, guild := range models.GetAllDiscordGuilds() {
		prefixSet[guild.DiscordID] = guild.GuildPrefix
	}

	models.InitCachedVars(conf)
}

func main() {
	var err error

	// Create a new shard manager using the provided bot token.
	Mgr, err = shards.New("Bot " + conf.Token)
	if err != nil {
		fmt.Println("[ERROR] Error creating manager,", err)
		return
	}

	// List some functions for the shard manager to call when events occur.
	subscribeEvents := []interface{}{
		commands.MessageCreate,
		onConnect,
		commands.OnMessageEdit,
		commands.OnDelete,
	}
	for _, event := range subscribeEvents {
		Mgr.AddHandler(event)
	}

	// In this example, we only care about receiving message events.
	Mgr.RegisterIntent(discordgo.IntentsAll)

	fmt.Println("[INFO] Starting shard manager...")

	// Start all of our shards and begin listening.
	err = Mgr.Start()
	if err != nil {
		fmt.Println("[ERROR] Error starting manager,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("[SUCCESS] Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Manager.
	fmt.Println("[INFO] Stopping shard manager...")
	Mgr.Shutdown()
	fmt.Println("[SUCCESS] Shard manager stopped. Bot is shut down.")
}

// This function will be called (due to AddHandler above) every time one
// of our shards connects.
func onConnect(s *discordgo.Session, evt *discordgo.Connect) {
	fmt.Printf("[INFO] Shard #%v connected.\n", s.ShardID)
	for {
		lastId := ""
		// Get all guilds the bot is in
		guilds, err := s.UserGuilds(100, "", lastId)
		if err != nil {
			fmt.Println("[ERROR] Error getting guilds,", err)
			break
		}
		// loop through all guilds
		for _, guild := range guilds {
			if guild.ID > lastId {
				lastId = guild.ID
			}
			dGuild, _ := s.Guild(guild.ID)
			if dGuild == nil {
				continue
			}
			dbGuild, err := models.GetOrCreateGuild(*dGuild, s)
			if err != nil {
				fmt.Println("[ERROR] Error getting guild,", err)
				continue
			}

			for {
				lastUserId := ""
				members, err := s.GuildMembers(guild.ID, lastUserId, 100)
				if err != nil {
					fmt.Println("[ERROR] Error getting members,", err)
					break
				}
				// loop through all members
				for _, member := range members {
					if member.User.ID > lastUserId {
						lastUserId = member.User.ID
					}
					usr, _ := models.GetOrCreateUser(member.User)
					if usr == nil {
						continue
					}
					_ = models.AddUserToGuild(usr, dbGuild)
				}
			}
		}
	}
}
