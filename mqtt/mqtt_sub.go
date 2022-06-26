package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func printMsg(c mqtt.Client, m mqtt.Message) {
	fmt.Println(m.Topic(), m.Payload())
}

func main() {
	// setup mqtt client options
	options := mqtt.NewClientOptions()
	options.AddBroker("tcp://localhost:1883")
	options.SetClientID("example_subscriber")

	// connect to mqtt broker.
	c := mqtt.NewClient(options)
	token := c.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// subscribe to a topic
	token = c.Subscribe("/ames/hq/ntf/d/pump/battery", 0, printMsg)
	token.Wait()
	fmt.Println("Subscribed to topic")

	// block forever
	select {}
}
