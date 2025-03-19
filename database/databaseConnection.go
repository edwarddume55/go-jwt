package database

import(
	"fmt"
	"log"
	"time"
	"os"
	"context"
	"github.com/joho/golan.ev"
	"go.mongodb.org/mongo.driver/mongo"
	"go/mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.client{
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env files")
	}
	MongoDb := os.Getenv("MONGODB_URL")

	client, err :=mongo.NewClient(options.client().ApplyURI(MongoDb))
	if err != nil{
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err= client.Connect(ctx)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return client
}

var client *mongo.client = DBinstance

func OpenCollection(client *mongo.Client, collectionName string ) *mongo.Collection{
	var collection *mongo.collection = client.Database("cluster0").collection(collectionName)
	return collection
}