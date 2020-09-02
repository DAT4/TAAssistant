package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func updateStudent(id string, discordID string) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongo.Uri))

	if err != nil {
		return err
	}

	defer client.Disconnect(ctx)

	database := client.Database(conf.Mongo.Db)
	students := database.Collection(conf.Mongo.Col)
	_, err = students.UpdateOne(ctx,
		bson.M{"ID": id},
		bson.D{
			{"$set", bson.M{"DiscordID": discordID}},
		},
	)
	if err != nil {
		fmt.Println("Updating discord number:", err)
		return err
	}
	return nil
}
func findStudent(id string, discord bool) (student *student, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongo.Uri))
	if err != nil {
		fmt.Println("Could not connect to client:", err)
		return nil, err
	}

	defer client.Disconnect(ctx)

	students := client.Database(conf.Mongo.Db).Collection(conf.Mongo.Col)
	//questions := *database.Collection("questions")

	if discord {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := students.FindOne(ctx, bson.M{"DiscordID": id}).Decode(&student)
		if err != nil {
			return nil, err
		}
	} else {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := students.FindOne(ctx, bson.M{"ID": id}).Decode(&student)
		if err != nil {
			return nil, err
		}
	}

	return student, nil
}
