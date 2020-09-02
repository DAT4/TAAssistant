package main

func (student student) fullName() string {
	if student.MiddleName != "" {
		return student.FirstName + " " + student.MiddleName + " " + student.LastName
	} else {
		return student.FirstName + " " + student.LastName
	}
}

type Database struct {
	Uri string `json:"uri"`
	Db  string `json:"db"`
	Col string `json:"col"`
}
type Conf struct {
	Token     string   `json:"token""`
	BotStatus string   `json:"botStatus"`
	StudRole  string   `json:"studRole"`
	RegChan   string   `json:"regChan"`
	Mongo     Database `json:"mongo"`
}
type student struct {
	FirstName  string `bson:"FirstName"`
	MiddleName string `bson:"MiddleName"`
	LastName   string `bson:"LastName"`
	ID         string `bson:"ID"`
	Role       string `bson:"Role"`
	Discord    string `bson:"DiscordID"`
}
