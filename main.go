package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

var conf Conf

func main() {
	var filePath string
	flag.StringVar(&filePath, "c", "conf.json", "The path for the config file")
	flag.Parse()
	fmt.Println("TAAssistant is Running!")
	err := configuration(filePath)
	if err != nil {
		fmt.Println("Loading configuration file:", err)
		return
	}
	discordConnect()
}

func configuration(path string) error {
	byteValue, err := ioutil.ReadFile(path)
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
