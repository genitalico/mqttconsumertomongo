package main

import (
	"fmt"
	"mqttconsumertomongo/services"
	"mqttconsumertomongo/settings"

	"os"
)

func main() {

	var err error
	var settingsValues settings.Settings

	args := os.Args[1:]

	if len(args) > 0 {
		fmt.Println("Reading configuration file:", args[0])

		settingsValues, err = settings.ReadFile(&args[0])
	} else {
		settingsValues, err = settings.ReadFile(nil)
	}

	services.SetSettingsValues(settingsValues)

	if err != nil {
		fmt.Println("Error reading the configuration file:", err)
		os.Exit(1)
	}

	mqttSettings := settingsValues.MqttSettings

	err = services.MqttConnection(mqttSettings, nil)

	if err != nil {
		fmt.Println("Error connecting to the MQTT broker:", err)
		os.Exit(1)
	}

	fmt.Println("Connected to the MQTT broker:", mqttSettings.MqttBrokerURL)

	err = services.MongoConnection(settingsValues.MongodbSettings, nil)

	if err != nil {
		fmt.Println("Error connecting to the MongoDB database:", err)
		os.Exit(1)
	}

	fmt.Println("Connected to the MongoDB database:", settingsValues.MongodbSettings.Url)

	err = services.MqttSubscribe(services.MessageHandler, mqttSettings.Topics)

	if err != nil {
		fmt.Println("Error subscribing to the topics:", err)
		os.Exit(1)
	}

	for key, value := range mqttSettings.Topics {
		fmt.Println("Subscribed to the topic:", key, "with QoS:", value)
	}

	select {}
}
