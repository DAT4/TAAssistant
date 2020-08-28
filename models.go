package main

func (student student) fullName() string {
	if student.MiddleName != "" {
		return student.FirstName+" "+student.MiddleName+" "+student.LastName
	} else {
		return student.FirstName+" "+student.LastName
	}
}
type student struct {
	FirstName string	`bson:"firstName"`
	MiddleName string	`bson:"middleName"`
	LastName string		`bson:"lastName"`
	ID string			`bson:"id"`
	Role string			`bson:"role"`
	Discord string		`bson:"discord"`
}

type question struct {
	Student student		`bson:"student"`
	ChannelID string	`bson:"channelId"`
	Timestamp int64		`bson:"timestamp"`
	Topic []string		`bson:"topic"`
	Question string		`bson:"question"`
	Active bool			`bson:"active"`
}
