package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestConnectionDB(t *testing.T) {

	// Arrange
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	dbpwd := os.Getenv("DB_PASSWORD")

	if strings.TrimSpace(dbhost) == "" ||
		strings.TrimSpace(dbport) == "" ||
		strings.TrimSpace(dbuser) == "" ||
		strings.TrimSpace(dbpwd) == "" {

		t.Skip("Skipping test as DB connection info is not present")
	}

	// mongodb://admin:admin@mongo:27017/

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority", dbuser, dbpwd, dbhost, dbport)
	t.Log("Test connection to DB, using URI: " + uri)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			t.Fatal(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		t.Fatal(err)
	}

	t.Log("Connection to DB succeeded")
}
