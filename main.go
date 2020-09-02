package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

var conf Conf

func main() {
	err := configuration()
	if err != nil {
		fmt.Println("Loading configuration file:", err)
		return
	}
	discordConnect()
}

func configuration() error {
	byteValue, err := ioutil.ReadFile("conf.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteValue, &conf)
	if err != nil {
		return err
	}
	return nil
}

func discordConnect() {
	discord, err := discordgo.New("Bot " + conf.Token)
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

func ready(s *discordgo.Session, even *discordgo.Event) {
	s.UpdateStatus(0, conf.BotStatus)
}
