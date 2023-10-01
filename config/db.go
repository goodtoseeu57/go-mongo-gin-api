package config

import (
	"context"
	"log"
	"my-api/module/models"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository interface {
	FindOneNote(ctx context.Context, filter bson.M) (*models.Note, error)
}

var NotesCollection *mongo.Collection
var UsersCollection *mongo.Collection

func LoadDBConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	uri := os.Getenv("MONGO_CONNECTION_STRING")

	if uri == "" {
		panic("MONGO_CONNECTION_STRING environment variable is not set")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	NotesCollection = client.Database("notes-go-api").Collection("notes")
	UsersCollection = client.Database("notes-go-api").Collection("users")

	var result bson.M
	if err := client.Database("notes-go-api").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB!")
}
