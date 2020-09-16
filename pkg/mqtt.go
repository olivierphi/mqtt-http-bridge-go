package pkg

import (
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ConnectToMqttBrokerAndSubscribeToTopics(rawUrl string, topics []Topic) (client mqtt.Client, error error) {
	client, error = ConnectToMqttBroker(rawUrl)
	if error != nil {
		return nil, error
	}
	error = SubscribeToTopics(client, topics)
	if error != nil {
		return nil, error
	}
	return
}

func ConnectToMqttBroker(rawUrl string) (mqtt.Client, error) {
	urlStruct, error := url.Parse(rawUrl)
	if error != nil {
		return nil, error
	}
	clientOpts := mqtt.NewClientOptions()
	// clientOpts.AddBroker()
	clientOpts.ClientID = "hellogo"
	clientOpts.CleanSession = false
	clientOpts.Servers = []*url.URL{urlStruct}
	client := mqtt.NewClient(clientOpts)
	clientConnectionToken := client.Connect()
	clientConnectionToken.Wait()
	if !client.IsConnected() {
		return nil, fmt.Errorf("Can't connect to %s", rawUrl)
	}
	error = clientConnectionToken.Error()
	if error != nil {
		return nil, error
	}
	fmt.Printf("MQTT client connected to '%s'.\n", rawUrl)

	return client, nil
}

type Topic struct {
	Name string
	Qos  byte
}

func SubscribeToTopics(client mqtt.Client, topics []Topic) error {
	for _, topic := range topics {

		subscriptionToken := client.Subscribe(topic.Name, topic.Qos, mqttIncomingMessageHandler)
		error := subscriptionToken.Error()
		if error != nil {
			return error
		}
		subscriptionToken.Wait()
		fmt.Printf("Subscribed to topic '%s' with QOS %v.\n", topic.Name, topic.Qos)
	}
	return nil
}

func mqttIncomingMessageHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received MQTT message:\n%v\n", string(msg.Payload()))
}
