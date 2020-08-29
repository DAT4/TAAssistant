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


