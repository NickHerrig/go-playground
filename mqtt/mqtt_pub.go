package main

import (
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// setup mqtt client options
	options := mqtt.NewClientOptions()
	options.AddBroker("tcp://localhost:1883")
	options.SetClientID("example_publisher")

	// connect to mqtt broker.
	c := mqtt.NewClient(options)
	token := c.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		token = c.Publish("autodoser/state", 0, false, "alive")
		token.Wait()
		time.Sleep(time.Second)
	}

}
