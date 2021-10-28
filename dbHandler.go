package main

import (
	"context"
	"fmt"
	"log"

	"encoding/json"
	"math/rand"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//struct data-data yang akan dimasukkan ke database
type record struct {
	Timestamp string `json:"timeStamp"`
	Tracenum  string `json:"traceNum"`
	Messageid string `json:"messageId"`
	Status    string `json:"status"`
	Amount    string `json:"InterBankSettlementAmount"`
}

func create(data, koleksi, topic string) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

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
	newData := record{}
	json.Unmarshal([]byte(data), &newData)
	newData.Timestamp = time.Now().Format("20060102") // format yyyyMMdd
	newData.Messageid = fmt.Sprint(rand.Intn(100))
	if topic == "mpc.json.bifast.response" {
		newData.Status = "incoming"
	} else {
		newData.Status = "outgoing"
	}
	fmt.Println(newData)

	collection := client.Database("testdb").Collection("settlement")
	insertResult, err := collection.InsertOne(context.TODO(), newData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// if koleksi == "channel" {
	// 	//konek ke collection <koleksi> di db "dbBifast"
	// 	collection := client.Database("dbBifast").Collection("biller")
	// 	//insert "newData" ke collection "channel"
	// 	insertResult, err := collection.InsertOne(context.TODO(), newData)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	// } else if koleksi == "adaptor" {
	// 	time.Sleep(1 * time.Second)
	// 	collection := client.Database("dbBifast").Collection("biller")
	// 	fmt.Println(newData.Tracenum)
	// 	filter := bson.D{{"tracenum", newData.Tracenum}}
	// 	update := bson.D{
	// 		{"$set", bson.D{{"timestamp", newData.Timestamp}, {"source", newData.Source}, {"destination", newData.Destination}}},
	// 	}
	// 	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	// }

	//disconnect dari server db
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func kafkaConsumer(topic, koleksi string) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
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
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			fmt.Println(string(msg.Value))
			create(string(msg.Value), koleksi, topic)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

}
