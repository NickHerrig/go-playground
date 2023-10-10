package main

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func printMsg(c mqtt.Client, m mqtt.Message) {

	var msg interface{}
	json.Unmarshal([]byte(m.Payload()), &msg)

	val, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(m.Topic())
	fmt.Println(string(val))
}

func main() {
	// setup mqtt client options
	options := mqtt.NewClientOptions()
	options.AddBroker("tcp://192.168.1.30:1883")
	options.SetClientID("example_subscriber")

	// connect to mqtt broker.
	c := mqtt.NewClient(options)
	token := c.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// subscribe to a topic
	token = c.Subscribe("#", 0, printMsg)
	token.Wait()
	fmt.Println("Subscribed to topic")

	// block forever
	select {}
}
