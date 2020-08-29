package Questionable

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"strings"
	"time"
)

func management(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel, author *main.student){
	//List questions
	if strings.HasPrefix(m.Content, "list") {
		listQuestions(s,m,c,author)
		return
	}

	if strings.HasPrefix(m.Content, ":students") {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		data, err := main.students.Find(ctx,bson.M{"role": "S"})
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
		err := s.GuildMemberMove(m.GuildID, author.Discord,&chID)
		if err != nil {
			fmt.Println("Moving member:", err)
		}
	}
}
