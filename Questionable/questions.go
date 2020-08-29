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

func studentAsksQuestion(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel, author *main.student){
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
		Answer: nil,
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

func listQuestions(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel, author *main.student){
	str := ""
	questions := getQuestionList()
	for _,e := range(questions){
		str += "```\n"
		str += "Emne:\t\t\t "
		for _,a := range(e.Topic) {
			str += a
			str += ","
		}
		str += "\n"
		str += "Navn:\t\t\t "
		str += e.Student.FirstName +" "
		if e.Student.MiddleName != ""{
			str += e.Student.FirstName + " "
		}
		str += e.Student.LastName
		str += "\n"
		str += "Studienummer:\t "
		str += e.Student.ID
		str += "\n"
		time := time.Unix(e.Timestamp,0)
		day := time.Weekday().String()
		minute := strconv.Itoa(time.Minute())
		if time.Minute() < 10 {
			minute = "0"+strconv.Itoa(time.Minute())
		}
		hour := strconv.Itoa(time.Hour())+":"+minute
		str += "Dag:\t\t\t  "
		str += day
		str += "\n"
		str += "Time:\t\t\t "
		str += hour
		str += "\n"
		str += "Channel:\t\t  "
		cha, err := s.Channel(e.ChannelID)
		if err != nil {
			str += "No channel"
		} else {
			str += cha.Name
		}
		str += "```"
		str += "**Spørgsmål:**\n"
		str += e.Question
		if e.Answer != nil{
			str += "\n"
			str += "**Svar:**\n"
			str += e.Answer.Answer
		}
	}
	_, _ = s.ChannelMessageSend(c.ID, str)
	return
}

func writeQuestion(q question) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ , err := main.questions.InsertOne(ctx,q,)
	if err != nil {
		return err
	}
	return nil
}

func getQuestionList() (ret []question) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := main.questions.Find(ctx, bson.M{"active": true})
	if err != nil {
		fmt.Println("Looking for quetions in MongoDB:", err)
		return nil
	}
	var q question
	for cursor.Next(context.Background()) {
		cursor.Decode(&q)
		ret = append(ret, q)
	}
	return ret
}

type question struct {
	Student   main.student `bson:"student"`
	ChannelID string       `bson:"channelId"`
	Timestamp int64        `bson:"timestamp"`
	Topic     []string     `bson:"topic"`
	Question  string       `bson:"question"`
	Active    bool         `bson:"active"`
	Answer    *answer      `bson:"answer"`
}

type answer struct {
	Student   main.student `bson:"student"`
	Timestamp int64        `bson:"timestamp"`
	Answer    string       `bson:"answer"`
	Approved  bool         `bson:"approved"`
}
