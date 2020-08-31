package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strconv"
	"time"
)

var mongURI = os.Getenv("MONGO_URI")

func updateStudent(id string, discordID string, channelID string, i int , status bool) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongURI))

	if err != nil {
		return err
	}

	defer client.Disconnect(ctx)

	database := client.Database("dtu")
	students := database.Collection("studenter")
	_, err = students.UpdateOne(ctx,
		bson.M{"ID": id},
		bson.D{
			{"$set", bson.M{
				"DiscordID": discordID,
				"Courses."+strconv.Itoa(i)+".OnDiscord": status,
			}},
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
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongURI))
	if err != nil {
		fmt.Println("Could not connect to client:", err)
		return nil, err
	}

	defer client.Disconnect(ctx)

	students := client.Database("dtu").Collection("studenter")
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
