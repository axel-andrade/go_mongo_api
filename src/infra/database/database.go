package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx = context.TODO()
var mongoClient *mongo.Client

func ConnectDB() {
	var err error

	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI"))
	mongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	/*
		Aqui, você chamou o método e passou um para ele junto com uma preferência de leitura primária usando ,
		que informa ao cliente MongoDB como ler as operações para os membros do conjunto de réplicas.Ping()contextreadpref.Primary()
	*/
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	// defer client.Disconnect(ctx)
}

func CloseDB() error {
	mongoClient.Disconnect(ctx)
	return nil
}

func GetDB() *mongo.Database {
	return mongoClient.Database(os.Getenv("DB_NAME"))
}
