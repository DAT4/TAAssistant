package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
	"time"
)

func deleteMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	time.Sleep(5 * time.Second)
	err := s.ChannelMessageDelete(m.ChannelID, m.ID)
	if err != nil {
		fmt.Println("Deleting message:", err)
	}
	return
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("Channel:", m.ChannelID)

	//Only works in channel registrering
	if m.ChannelID != "749103315597394020" {
		return
	}
	go deleteMessage(s, m)

	//Exclude messages from web bot
	if m.Author.ID == "748621295092105336" {
		return
	}
	//Exclude messages from client
	if m.Author.ID == s.State.User.ID {
		return
	}

	//Identify channel
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}

	//Identify author
	author, err := findStudent(m.Author.ID, true)
	if err != nil {
		fmt.Println("Trying to identify user", err)

		//Register member
		register, err := regexp.MatchString(`s\d{6}`, strings.ToLower(m.Content))
		if err != nil {
			fmt.Println("Matching studentID:", err)
			return
		}
		fmt.Println(m.Author.Username, "wrote:", m.Content)
		if register {
			registerStudent(s, m, c)
			return
		}
		s.ChannelMessageSend(c.ID, "Velkommen "+m.Author.Username+", Vil du venligst identificere dig selv med dit studienummer.\n**Eksempel:**\n```s195469```")
		return
	}

	//Logs that registered user writes message
	fmt.Println("Registerd user:", author.FirstName, "wrote", m.Content, "in Channel(", "\""+c.Name+"\"", ")")
	fmt.Println("His ID:", author.Discord)

	if author.Role == "TA" {
		del, _ := regexp.MatchString(`##s\d{6}`, m.Content)
		if del {
			fmt.Println("Deleting user", m.Content)
			unRegisterStudent(s, m, c)
		}

	}
	return
}
