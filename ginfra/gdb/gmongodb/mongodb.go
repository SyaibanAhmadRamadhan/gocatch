package gmongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func OpenConnMongoClient(uri string) (*mongo.Client, error) {
	clientOpt := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpt)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func OpenConnMongoDB(uri string, dbName string) (*mongo.Database, error) {
	client, err := OpenConnMongoClient(uri)
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)

	return db, nil
}
