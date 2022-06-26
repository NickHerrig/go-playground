package main

import (
	"encoding/json"
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

	// create and marshal topic message to json
	type Message struct {
		Battery float64
	}
	m := Message{0.85}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	for {
		token = c.Publish("/ames/hq/ntf/d/pump/battery", 0, false, b)
		token.Wait()
		time.Sleep(time.Second)
	}

}
