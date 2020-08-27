package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var students mongo.Collection
var questions mongo.Collection
var mongURI = os.Getenv("MONGO_URI")

func dbConnect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongURI))

	if err != nil {
		fmt.Println("Connecting to database:",err)
	}

	//defer client.Disconnect(ctx)
	database := client.Database("dtu")
	students = *database.Collection("students")
	questions = *database.Collection("questions")
}

func findStudent(id string, discord bool) (*student, error) {
	var stud bson.M
	ret := student{}
	if discord{
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := students.FindOne(ctx, bson.M{"discord":id}).Decode(&stud)
		if err != nil {
			fmt.Println("Looking for student with id",id+":",err)
			return nil , err
		}
	} else {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := students.FindOne(ctx, bson.M{"id":id}).Decode(&stud)
		if err != nil {
			fmt.Println("Looking for student with id",id+":",err)
			return nil, err
		}
	}
	ret = student{
		FirstName:	stud["f_name"].(string),
		MiddleName:	stud["m_name"].(string),
		LastName:  	stud["l_name"].(string),
		ID:      	stud["id"].(string),
		Role:    	stud["role"].(string),
		Discord: 	stud["discord"].(string),
	}

	return &ret, nil
}


