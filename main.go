package main

import (
	"fmt"
)

func main() {

	fmt.Println("mulai")

	go kafkaConsumer("mpc.json.bifast.request", "adaptor")
	kafkaConsumer("mpc.json.bifast.response", "channel")

}
