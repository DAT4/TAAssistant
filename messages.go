package main

import (
	"errors"
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
	regChan, err := getChannel(s, m, conf.RegChan)
	if err != nil {
		fmt.Println("Looking for channel Registration:", err)
	}
	if m.ChannelID == regChan.ID {
		// Indledende programmering
	} else {
		return
	}

	go deleteMessage(s, m)

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
	studRole, err := getRole(s, m, conf.StudRole)
	if err != nil {
		fmt.Println("Looking for role Student:", err)
		return
	}
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
			registerStudent(s, m, c, studRole)
			return
		}
		s.ChannelMessageSend(c.ID, "Velkommen "+m.Author.Username+", Vil du venligst identificere dig selv med dit studienummer.\n**Eksempel:**\n```s195469```")
		return
	}

	//Logs that registered user writes message
	fmt.Println("Registerd user:", author.FirstName, "wrote", m.Content, "in Channel(", "\""+c.Name+"\"", ")")
	fmt.Println("His ID:", author.Discord)

	if author.Role == "TA" {
		del, _ := regexp.MatchString(`delete\(s\d{6}\)`, strings.ToLower(m.Content))
		if del {
			fmt.Println("Deleting user", m.Content, "with ID:", del)
			unRegisterStudent(s, m, c, studRole)
		}
		return
	}
	s.ChannelMessageSend(c.ID, "```"+author.FirstName+" du er allerede registreret, med ID: "+author.ID+"```")
	return
}

func getRole(s *discordgo.Session, m *discordgo.MessageCreate, prefix string) (roleID string, err error) {
	roles, err := s.GuildRoles(m.GuildID)
	if err != nil {
		return "", err
	}
	for _, e := range roles {
		if strings.HasPrefix(e.Name, prefix) {
			fmt.Println("Role permission", e.Permissions)
			return e.ID, nil
		}
	}
	return "", errors.New("Something went wrong...")
}

func getChannel(s *discordgo.Session, m *discordgo.MessageCreate, prefix string) (*discordgo.Channel, error) {
	channels, err := s.GuildChannels(m.GuildID)
	if err != nil {
		return nil, err
	}
	for _, e := range channels {
		if strings.HasPrefix(e.Name, prefix) {
			return e, nil
		}
	}
	return nil, errors.New("No channel with that name.")
}
