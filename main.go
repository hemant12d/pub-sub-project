package main

import (
	"fmt"
	"time"
)

// available topics
var availableTopics = map[string]string{
	"BTC": "BITCOIN",
	"ETH": "ETHEREUM",
	"DOT": "POLKADOT",
	"SOL": "SOLANA",
}

func main() {

	//construct new broker.
	broker := NewBroker()
	fmt.Println(broker)

	//create new subscriber
	s1 := broker.AddSubscriber()
	// subscribe BTC and ETH to s1.
	broker.Subscribe(s1, availableTopics["BTC"])
	broker.Subscribe(s1, availableTopics["ETH"])

	// create new subscriber
	s2 := broker.AddSubscriber()
	// subscribe ETH and SOL to s2.
	broker.Subscribe(s2, availableTopics["ETH"])
	broker.Subscribe(s2, availableTopics["SOL"])

	go (func() {
		// sleep for 5 sec, and then subscribe for topic DOT for s2
		time.Sleep(3 * time.Second)
		broker.Subscribe(s2, availableTopics["DOT"])
	})()

	go (func() {
		// sleep for 5 sec, and then unsubscribe for topic SOL for s2
		time.Sleep(5 * time.Second)
		broker.Unsubscribe(s2, availableTopics["SOL"])
		fmt.Printf("Total subscribers for topic ETH is %v\n", broker.GetSubscribers(availableTopics["ETH"]))
	})()

	go (func() {
		// s;eep for 5 sec, and then unsubscribe for topic SOL for s2
		time.Sleep(10 * time.Second)
		broker.RemoveSubscriber(s2)
		fmt.Printf("Total subscribers for topic ETH is %v\n", broker.GetSubscribers(availableTopics["ETH"]))
	})()

	// Concurrently publish the values.
	go pricePublisher(broker)

	// Concurrently listens from s1.
	go s1.Listen()

	// Concurrently listens from s2.
	go s2.Listen()

	// to prevent terminate
	fmt.Scanln()
	fmt.Println("Done!")
}
