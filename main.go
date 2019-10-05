package main

import (
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/openfaas-incubator/connector-sdk/types"
)

type connectorConfig struct {
	*types.ControllerConfig
	BrokerURL string
}

func main() {
	credentials := types.GetCredentials()
	config := buildConnectorConfig()

	controller := types.NewController(credentials, config.ControllerConfig)

	controller.BeginMapBuilder()

	client := connectMqtt(config, controller)

	subscribeToTopics(client, controller)

}

func createMqttCLientOptions(config connectorConfig, controller *types.Controller) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.BrokerURL)
	opts.SetAutoReconnect(true)
	opts.SetClientID("openfaas")

	opts.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
		log.Println(message)

		payload := message.Payload()

		controller.Invoke(message.Topic(), &payload)
	})

	return opts
}

func connectMqtt(config connectorConfig, controller *types.Controller) mqtt.Client {
	opts := createMqttCLientOptions(config, controller)
	client := mqtt.NewClient(opts)

	token := client.Connect()
	token.Wait()

	if token.Error() != nil {
		log.Fatal(token.Error())
	}

	return client
}

func subscribeToTopics(client mqtt.Client, controller *types.Controller) {
	subscribe := func(topic string) {
		log.Printf("Subscribe to topic %v.", topic)
		client.Subscribe(topic, 2, nil)
	}

	unsubscribe := func(topic string) {
		log.Printf("Unubscribe from topic %v.", topic)
		client.Unsubscribe(topic)
	}

	topics := NewTopics(subscribe, unsubscribe)

	for {
		topics.UpdateTopics(controller.Topics())

		time.Sleep(controller.Config.RebuildInterval)
	}
}

func buildConnectorConfig() connectorConfig {

	brokerURL := "tcp://mqtt-broker:1883"
	if val, exists := os.LookupEnv("mqtt_url"); exists {
		brokerURL = val
	}

	gatewayURL := "http://gateway:8080"
	if val, exists := os.LookupEnv("gateway_url"); exists {
		gatewayURL = val
	}

	upstreamTimeout := time.Second * 30
	rebuildInterval := time.Second * 3

	if val, exists := os.LookupEnv("upstream_timeout"); exists {
		parsedVal, err := time.ParseDuration(val)
		if err == nil {
			upstreamTimeout = parsedVal
		}
	}

	if val, exists := os.LookupEnv("rebuild_interval"); exists {
		parsedVal, err := time.ParseDuration(val)
		if err == nil {
			rebuildInterval = parsedVal
		}
	}

	printResponse := false
	if val, exists := os.LookupEnv("print_response"); exists {
		printResponse = (val == "1" || val == "true")
	}

	printResponseBody := false
	if val, exists := os.LookupEnv("print_response_body"); exists {
		printResponseBody = (val == "1" || val == "true")
	}

	delimiter := ","
	if val, exists := os.LookupEnv("topic_delimiter"); exists {
		if len(val) > 0 {
			delimiter = val
		}
	}

	asynchronousInvocation := false
	if val, exists := os.LookupEnv("asynchronous_invocation"); exists {
		asynchronousInvocation = (val == "1" || val == "true")
	}

	return connectorConfig{
		ControllerConfig: &types.ControllerConfig{
			UpstreamTimeout:          upstreamTimeout,
			GatewayURL:               gatewayURL,
			PrintResponse:            printResponse,
			PrintResponseBody:        printResponseBody,
			RebuildInterval:          rebuildInterval,
			TopicAnnotationDelimiter: delimiter,
			AsyncFunctionInvocation:  asynchronousInvocation,
		},
		BrokerURL: brokerURL,
	}
}
