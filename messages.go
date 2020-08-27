package main

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	//Exclude messages from client
	if m.Author.ID == s.State.User.ID {
		return
	}

	//Only works in registrering
	if m.ChannelID != "748525251889070100" {
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

		//Register member
		register, _ := regexp.MatchString(`s\d{6}`,strings.ToLower(m.Content))
		fmt.Println(m.Content)
		if register{
			registerStudent(s,m,c)
			return
		}
		s.ChannelMessageSend(c.ID, "```Jeg kan ikke genkende dig "+m.Author.Username+", Vil du venligst identificere dig selv med dit studienummer.```\n**Eksempel:**\n```s195469```")
		return
	}

	//Logs that registered user writes message
	fmt.Println("Registerd user:",author.FirstName,"wrote", m.Content, "in Channel(", "\""+c.Name+"\"",")")
	fmt.Println("His ID:", author.Discord)


	//Student asks question
	if strings.HasPrefix(strings.ToLower(m.Content),"hjælp::") {
		studentAsksQuestion(s,m,c,author)
	}


	if author.Role == "N/S"{
		del, _ := regexp.MatchString(`##s\d{6}`,m.Content)
		if del {
			unRegisterStudent(s,m,c)
		}

		//List questions
		if strings.HasPrefix(m.Content, "list") {
			roles, _ := s.GuildRoles(m.GuildID)
			for _,e := range(roles){
				fmt.Println(e.Name,"has ID", e.ID)
			}
			_, _ = s.ChannelMessageSend(c.ID, m.Author.Mention()+" du får din liste når jeg er færdig med at udvikle den.")
			return
		}

		if strings.HasPrefix(m.Content, ":students") {
			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			data, err := students.Find(ctx,bson.M{"role":"S"})
			if err != nil {
				fmt.Println("Finding students in database:", err)
			}
			number := 0
			for data.Next(context.Background()) {
				number += 1
			}
			s.ChannelMessageSend(c.ID, "```Der findes "+strconv.Itoa(number)+" studenter i databasen```")
		}

		if strings.HasPrefix(m.Content, "chan") {
			channels, err := s.GuildChannels(m.GuildID)
			if err != nil {
				fmt.Println("Getting channels:", err)
			}

			//Checking all the channels
			textChannels := "```TEXTCHANNELS:\n\n"
			voiceChannels := "```VOICECHANNELS:\n\n"
			for _,e := range channels{
				if e.Type == 0 {
					text := e.Name+"\n\tID: "+e.ID+"\n\n"
					textChannels += text
				}
				if e.Type == 2 {
					text := e.Name+"\n\tID: "+e.ID+"\n\n"
					voiceChannels += text
				}
				if err != nil {
					fmt.Println("Sending message:", err)
				}
			}
			textChannels += "```"
			voiceChannels += "```"
			_, err = s.ChannelMessageSend(c.ID, textChannels)
			_, err = s.ChannelMessageSend(c.ID, voiceChannels)
		}

		// Move self to voice channel Grupperum-4
		if strings.HasPrefix(m.Content, "move") {
			chID := "747752735079r23745" // Grupperum 4
			err = s.GuildMemberMove(m.GuildID, author.Discord,&chID)
			if err != nil {
				fmt.Println("Moving member:", err)
			}
		}

	}



}

