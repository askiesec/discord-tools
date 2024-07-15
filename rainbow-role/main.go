package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	token      string
	guildID    string
	roleID     string
	interval   int
	colors     = []int{0xff0000, 0xff7f00, 0xffff00, 0x00ff00, 0x0000ff, 0x4b0082, 0x8b00ff}
	colorIndex = 0
)

func init() {
	flag.StringVar(&token, "token", "", "Discord Bot Token")
	flag.StringVar(&guildID, "guild", "", "Discord Guild ID")
	flag.StringVar(&roleID, "role", "", "Discord Role ID")
	flag.IntVar(&interval, "interval", 10, "Interval in seconds to change the role color")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
}

func main() {
	if token == "" || guildID == "" || roleID == "" {
		log.Fatalf("Token, Guild ID, and Role ID are required.")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %v", err)
	}
	defer dg.Close()

	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-ticker.C:
				updateRoleColor(dg, guildID, roleID)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	<-quit
	fmt.Println("Shutting down bot.")
}

func updateRoleColor(s *discordgo.Session, guildID, roleID string) {
	colorIndex = (colorIndex + 1) % len(colors)
	newColor := colors[colorIndex]

	role, err := s.State.Role(guildID, roleID)
	if err != nil {
		log.Printf("Error getting role: %v", err)
		return
	}

	roleParams := &discordgo.RoleParams{
		Color:       &newColor,
		Name:        role.Name,
		Hoist:       &role.Hoist,
		Permissions: &role.Permissions,
		Mentionable: &role.Mentionable,
	}

	_, err = s.GuildRoleEdit(guildID, roleID, roleParams)
	if err != nil {
		log.Printf("Error updating role color: %v", err)
		return
	}

	log.Printf("Updated role %s color to %06x", role.Name, newColor)
}
