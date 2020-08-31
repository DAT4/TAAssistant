package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func registerStudent(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel, channel Channel) {
	str := strings.Trim(strings.ToLower(m.Content), "%")
	person, err := findStudent(str, false)
	if err != nil {
		s.ChannelMessageSend(c.ID, "```Jeg kunne ikke finde dit studienummer på listen...```")
		return
	}

	if person.Discord != "" && person.Discord != m.Author.ID {
		user, err := s.User(person.Discord)
		if err != nil {
			fmt.Println("Looking for user in Guild:", err)
		}
		s.ChannelMessageSend(c.ID, "```En Discord bruger med brugernavn "+user.Username+" er allerede registreret med dette studienummer.```")
		return
	}

	s.ChannelMessageSend(c.ID, "```Hejsa "+person.fullName()+"\nDu registreres og får nu adgang til diverse tekst og talekanaler!```")
	updateStudent(str, m.Author.ID, channel.ID, channel.I, true)

	// Changing name and adding student to Student role
	if person.Role == "S" {
		err = s.GuildMemberNickname(m.GuildID, m.Author.ID, person.FirstName+" "+person.LastName+" - "+person.ID)
		if err != nil {
			fmt.Println("Changing", m.Author.Username, "'s nickname:", err)
			return
		}
		err = s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, channel.RoleID)
		if err != nil {
			fmt.Println("Adding", m.Author.Username, "to role Student:", err)
			return
		}
	}
	return
}

func unRegisterStudent(s *discordgo.Session, m *discordgo.MessageCreate, c *discordgo.Channel, channel Channel) {

	str := strings.Trim(strings.Trim(strings.ToLower(m.Content), "delete("),")")
	fmt.Println(str)
	stud, err := findStudent(str, false)
	if err != nil {
		fmt.Println("looking for student", str, ":", err)
		return
	}
	user, err := s.User(stud.Discord)
	if err != nil {
		fmt.Println("Looking for user in Guild:", err)
		return
	}
	err = s.GuildMemberNickname(m.GuildID, user.ID, user.Username)
	if err != nil {
		fmt.Println("Changing", user.Username, "'s nickname:", err)
		return
	}

	err = s.GuildMemberRoleRemove(m.GuildID, stud.Discord, channel.RoleID)
	if err != nil {
		fmt.Println("Removing", user.Username, "from role Student:", err)
		return
	}
	s.ChannelMessageSend(c.ID, "```"+stud.ID+" "+"fjernes fra discord bruger : "+user.Username+"```")

	//Connectiong to MongoDB to Update
	fmt.Println(str, "slettes.")
	err = updateStudent(str, "", channel.ID, channel.I, false)
	if err != nil {
		fmt.Println("Deleting student", str, ":", err)
	}
	return
}
