package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get tasks by ID
func getByID(id string, c *mongo.Collection) Task {
	var res Task
	bsonID, err := primitive.ObjectIDFromHex(id)
	err = collection.FindOne(context.TODO(), bson.M{"_id": bsonID}).Decode(&res)
	if err != nil {
		log.Fatal("Error", err)
	}
	fmt.Printf("Found a single document: %+v\n", res)
	return res
}

// Get all tasks
func getAll(c *mongo.Collection) []bson.M {
	cursor, err := c.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var res []bson.M
	if err = cursor.All(context.TODO(), &res); err != nil {
		log.Fatal(err)
	}
	return res
}

// Add a task
func insert(t Task, c *mongo.Collection) {
	res, err := c.InsertOne(context.TODO(), t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)
}

// Delete a task
func delete(id string, c *mongo.Collection) {
	bsonID, err := primitive.ObjectIDFromHex(id)
	res, err := c.DeleteOne(context.TODO(), bson.M{"_id": bsonID})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
}

// Updpate a task
func update(id string, t Task, c *mongo.Collection) {
	bsonID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	update := bson.M{
		"$set": bson.M{
			"name":        t.Name,
			"description": t.Description,
			"status":      t.Status,
		},
	}
	c.FindOneAndUpdate(context.TODO(), bson.M{"_id": bsonID}, update)
}
