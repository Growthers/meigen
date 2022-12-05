package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type DB struct {
	client *mongo.Client
}

type Meigen struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

func (db *DB) getURI() (uri string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	if uri = os.Getenv("DB_URI"); uri == "" {
		log.Fatal("DB_URI is not set")
	}
	return uri
}

func (db *DB) Connect(ctx context.Context) (err error) {
	opt := options.Client().ApplyURI(db.getURI())
	if err = opt.Validate(); err != nil {
		return err
	}
	db.client, err = mongo.Connect(ctx, opt)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Disconnect(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}

func (db *DB) Ping(ctx context.Context) (err error) {
	if err = db.client.Ping(ctx, nil); err != nil {
		return err
	}
	fmt.Println("Ping to MongoDB successful")
	return nil
}

func (db *DB) AddMeigen(ctx context.Context, a string, t string) (err error) {
	DB_NAME := os.Getenv("DB_NAME")
	COLLECTION_NAME := os.Getenv("DB_COLLECTION_NAME")
	collection := db.client.Database(DB_NAME).Collection(COLLECTION_NAME)
	m := Meigen{}

	m.ID = db.CountMeigen(ctx) + 1
	m.Author = a
	m.Text = t

	_, err = collection.InsertOne(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) CountMeigen(ctx context.Context) int {
	DB_NAME := os.Getenv("DB_NAME")
	COLLECTION_NAME := os.Getenv("DB_COLLECTION_NAME")
	collection := db.client.Database(DB_NAME).Collection(COLLECTION_NAME)
	count, _ := collection.CountDocuments(ctx, bson.D{})
	return int(count)
}
