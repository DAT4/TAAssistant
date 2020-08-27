package main

func (student student) fullName() string {
	if student.MiddleName != "" {
		return student.FirstName+" "+student.MiddleName+" "+student.LastName
	} else {
		return student.FirstName+" "+student.LastName
	}
}
type student struct {
	FirstName string
	MiddleName string
	LastName string
	ID string
	Role string
	Discord string
}

type question struct {
	Student student
	Timestamp int64
	Topic []string
	Question string
	Active bool
}
