package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	fmt.Println("mulai")
	readConfigJSON(&mongoConfig, &kafkaConfig)

	go kafkaConsumer(mongoConfig.CreditTransferLogCollection, kafkaConfig.CreditTransferLogTopic)
	go kafkaConsumer(mongoConfig.SystemLogCollection, kafkaConfig.SystemLogTopic)
	kafkaConsumer(mongoConfig.SettlementLogCollection, kafkaConfig.SettlementLogTopic)

}

func readConfigJSON(mongoConfigStruct *MongoConfig, kafkaConfigStruct *KafkaConfig) {
	// Get config file
	log.Println("Opening config file")
	file, _ := os.Open("./config.json")
	defer file.Close()

	// Read config file
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("ReadFile error: ", err)
	}

	// Unmarshal config
	err = json.Unmarshal(b, &mongoConfigStruct)
	if err != nil {
		log.Fatal("Json Unmarshall error: ", err)
	}
	// Unmarshal kafka config
	err = json.Unmarshal(b, &kafkaConfigStruct)
	if err != nil {
		log.Fatal("Json Unmarshall error: ", err)
	}
}
