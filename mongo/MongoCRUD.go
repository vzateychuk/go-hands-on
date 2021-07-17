package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {

	//region Connect to MongoDB

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Set client options

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil) // Check the connection
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	//endregion
	//region Getting collection

	collection := client.Database("test").Collection("trainers")
	fmt.Println("Database(test).Collection(trainers):", collection)

	//endregion
	//region Insert documents

	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}

	// To insert a single document, use the collection.InsertOne() method:
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// To insert multiple documents at a time, the collection.InsertMany() method will take a slice of objects:
	trainers := []interface{}{misty, brock}
	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	//endregion
	//region Update document
	filter := bson.D{{"name", "Ash"}} // Will match the document where the name is Ash and will increment Ash's age by 1

	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	//endregion
	//region FindOne document

	var result Trainer // create a value into which the result can be decoded

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
	//endregion
	//region FindMany

	findOptions := options.Find()
	findOptions.SetLimit(2) // Pass these options to the Find method

	// Here's an array in which you can store the decoded documents
	var results []*Trainer

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor, iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	//endregion
	//region DeleteMany

	// deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}}) удалит все без исключения
	deleteFilter := bson.D{
		{"name", "Misty"},
	} // Will match the document where the name is Ash and will increment Ash's age by 1

	deleteResult, err := collection.DeleteMany(context.TODO(), deleteFilter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	//endregion
	//region Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
	//endregion
}
