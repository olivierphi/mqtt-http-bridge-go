package main

import (
	"fmt"

	"github.com/DrBenton/mqtt-http-bridge-go/pkg"
)

func main() {
	go connectToMqttBrokers()
	pkg.StartHttpServer(8080)
}

func connectToMqttBrokers() {
	topicName := "test/e1CzbjHBW5GaiAy"
	mqttClient, error := pkg.ConnectToMqttBrokerAndSubscribeToTopics("tcp://broker.hivemq.com:1883", []pkg.Topic{{Name: topicName, Qos: 1}})
	if error != nil {
		panic(error)
	}
	fmt.Printf("%v", mqttClient.IsConnected())

}
