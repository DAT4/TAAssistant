package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

func studentAsksQuestion(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel, author *student){
	q := strings.Split(strings.Split(m.Content,"::")[1], ";;")
	if len(q) < 2{
		s.ChannelMessageSend(c.ID, "```" +
			"Spørgsmålet er ikke formatteret ordenligt.. Prøv at følg den anbefalede formattering.\n" +
			"Eksmpel:\n" +
			"hjælp:: emne;; spørgsmål\n\n" +
			"Hvis du har mere end et emne så sepperer den med komma.\n" +
			"Eksempel:\n" +
			"hjælp:: emne1, emne2, emne3;; spørgsmål\n\n" +
			"Hvis du ikke har noget emne så skriv:\n" +
			"Eksempel:\n" +
			"hjælp::;; spørgsmål"+
			"```")
		return
	}
	q[0] = strings.TrimSpace(q[0])
	emner := strings.Split(q[0],",")

	data := question{
		Student:  *author,
		ChannelID: m.ChannelID,
		Timestamp: time.Now().Unix(),
		Topic:    emner,
		Question: strings.TrimSpace(q[1]),
		Active: true,
	}


	err := writeQuestion(data)
	if err != nil {
		fmt.Println("Write question to MondoDB:",err)
		return
	}

	s.ChannelMessageSend(c.ID, m.Author.Mention()+"```SPØRGSMÅL:\n\t" +
		"Student: "+author.FirstName+"\n\t" +
		"Emne: "+strings.Join(data.Topic,",")+"\n\t" +
		"Text: "+data.Question+"```")
}

