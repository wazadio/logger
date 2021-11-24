package main

import (
	"context"
	"fmt"
	"log"

	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func create(data, koleksi string) {
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoConfig.MongoServer)

	// Connect to MongoDB (server)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// konversi string "data" ke struct "newData"

	if koleksi == mongoConfig.SystemLogCollection {
		newData := SystemLoggerStruct{}
		json.Unmarshal([]byte(data), &newData)

		collection := client.Database(mongoConfig.Database).Collection(koleksi)
		insertResult, err := collection.InsertOne(context.TODO(), newData)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Inserted a single document SystemLogs: ", insertResult.InsertedID)
	}

	if koleksi == mongoConfig.SettlementLogCollection {
		newData := SettlementLogStruct{}
		json.Unmarshal([]byte(data), &newData)
		fmt.Println(newData)

		collection := client.Database(mongoConfig.Database).Collection(koleksi)
		insertResult, err := collection.InsertOne(context.TODO(), newData)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Inserted a single document to SettlementLogs: ", insertResult.InsertedID)
	}

	if koleksi == mongoConfig.CreditTransferLogCollection {
		newData := CreditTransferLogStruct{}
		json.Unmarshal([]byte(data), &newData)
		fmt.Println(newData)

		collection := client.Database(mongoConfig.Database).Collection(koleksi)
		insertResult, err := collection.InsertOne(context.TODO(), newData)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Inserted a single document to BalanceLogs: ", insertResult.InsertedID)
	}

	// switch koleksi {
	// case mongoConfig.SystemLogCollection:
	// 	newData := SystemLoggerStruct{}
	// case mongoConfig.SettlementLogCollection:
	// 	newData := SettlementLogStruct{}
	// case mongoConfig.CreditTransferLogCollection:
	// 	newData := CreditTransferLogStruct{}
	// }

	// json.Unmarshal([]byte(data), &newData)

	// collection := client.Database(mongoConfig.Database).Collection(koleksi)
	// insertResult, err := collection.InsertOne(context.TODO(), newData)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// json.Unmarshal([]byte(data), &newData)
	// fmt.Println(newData)

	// collection := client.Database(mongoConfig.Database).Collection(koleksi)
	// insertResult, err := collection.InsertOne(context.TODO(), newData)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	//disconnect dari server db
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func kafkaConsumer(koleksi, topic string) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaConfig.KafkaAddress,
		"group.id":          "myGroup",
		"auto.offset.reset": "latest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			// fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			// fmt.Println(string(msg.Value))
			create(string(msg.Value), koleksi)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

}
