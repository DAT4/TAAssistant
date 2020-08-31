package main

func (student student) fullName() string {
	if student.MiddleName != "" {
		return student.FirstName + " " + student.MiddleName + " " + student.LastName
	} else {
		return student.FirstName + " " + student.LastName
	}
}

type Role struct {
	ID 			string	`bson:"ID"`
	OnDiscord 	bool	`bson:"OnDiscord"`
}

type Channel struct {
	I 		int
	ID 		string
	RoleID	string
}

type student struct {
	FirstName  string 	`bson:"FirstName"`
	MiddleName string 	`bson:"MiddleName"`
	LastName   string 	`bson:"LastName"`
	ID         string 	`bson:"ID"`
	Role       string 	`bson:"Role"`
	Discord    string 	`bson:"DiscordID"`
	Courses	   []Role 	`bson:"Courses"`
}
