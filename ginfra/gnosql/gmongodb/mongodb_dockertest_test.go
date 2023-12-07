package gmongodb

import (
	"context"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra"
)

func TestMongoDockerTest(t *testing.T) {
	dockerTest := ginfra.InitDockerTest()
	defer dockerTest.CleanUp()

	mongoDockerTestConf := MongoDockerTestConf{}

	var mongoClient *mongo.Client

	dockerTest.NewContainer(mongoDockerTestConf.ImageVersion(dockerTest, ""), func(res *dockertest.Resource) error {
		time.Sleep(10 * time.Second)
		conn, err := mongoDockerTestConf.ConnectClient(res)
		gcommon.PanicIfError(err)
		mongoClient = conn

		return nil
	})
	// query := bson.D{
	// 	{
	// 		Key: "$and", Value: bson.A{
	// 		bson.D{{Key: "field1", Value: "value1"}},
	// 		bson.D{{Key: "field2", Value: "value2"}},
	// 	},
	// 	},
	// 	{
	// 		Key: "$or", Value: bson.A{
	// 		bson.D{{Key: "field3", Value: "value3"}},
	// 		bson.D{{Key: "field4", Value: "value4"}},
	// 	},
	// 	},
	// }

	t.Run("Sample Integration Test for MongoDB", func(t *testing.T) {
		mongodbConn := mongoClient.Database("exampleDatabase")
		err := mongodbConn.CreateCollection(context.Background(), "exampleCollection")

		if err != nil {
			t.Errorf("Failed to create collection: %v", err)
		}
		collection := mongodbConn.Collection("exampleCollection")

		document := bson.M{"name": "example", "value": "entry"}
		insertResult, err := collection.InsertOne(context.Background(), document)
		if err != nil {
			t.Errorf("Failed to insert document: %v", err)
		}

		if insertResult.InsertedID == nil {
			t.Error("No ID found for inserted document")
		}

		// Retrieve and validate
		filter := bson.M{"name": "example"}
		var result bson.M
		err = collection.FindOne(context.Background(), filter).Decode(&result)
		if err != nil {
			t.Errorf("Failed to find document: %v", err)
		}

		if result["name"] != "example" || result["value"] != "entry" {
			t.Errorf("Document does not match expected values: %v", result)
		}
	})
}
