package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var token = os.Getenv("DISCORD_TOKEN")
func main() {
	discordConnect()
}

func discordConnect(){
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Connect to the bot", err)
		return
	}
	discord.AddHandler(ready)
	discord.AddHandler(messageCreate)

	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages)
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening discord:", err)
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	discord.Close()
}


func ready(s *discordgo.Session, even *discordgo.Ready){
	s.UpdateStatus(0,"Tower of Hanoi")
}

