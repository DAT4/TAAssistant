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
	ret := student{}
	if discord{
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := students.FindOne(ctx, bson.M{"discord":id}).Decode(&ret)
		if err != nil {
			fmt.Println("Looking for student with id",id+":",err)
			return nil , err
		}
	} else {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := students.FindOne(ctx, bson.M{"id":id}).Decode(&ret)
		if err != nil {
			fmt.Println("Looking for student with id",id+":",err)
			return nil, err
		}
	}
	return &ret, nil
}

func writeQuestion(q question) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ , err := questions.InsertOne(ctx,q,)
	if err != nil {
		return err
	}
	return nil
}

func getQuestionList() (ret []question){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor , err := questions.Find(ctx, bson.M{"active":true})
	if err != nil {
		fmt.Println("Looking for quetions in MongoDB:",err)
		return nil
	}
	var q question
	for cursor.Next(context.Background()){
		cursor.Decode(&q)
		ret = append(ret, q)
	}
	return ret
}

