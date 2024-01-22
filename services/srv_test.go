package services

import (
	"mqttconsumertomongo/interfaces"
	"mqttconsumertomongo/settings"
	"testing"
)

var testSettings = settings.Settings{
	MqttSettings: settings.MqttSettings{
		MqttBrokerURL: "mqtt://localhost:1883",
		User:          "testuser",
		Password:      "testpassword",
		ClientId:      "testclient",
		Topics:        map[string]byte{"topic1": 1, "topic2": 2},
	},
	MongodbSettings: settings.MongodbSettings{
		Url:        "mongodb://localhost:27017",
		AuthSource: "admin",
		Username:   "testuser",
		Password:   "testpassword",
		DbName:     "testdb",
		Collection: "testcollection",
	},
}

func TestMongoConnection(t *testing.T) {
	mongoSimulatedClient := &interfaces.Mongo{}

	MongoConnection(testSettings.MongodbSettings, mongoSimulatedClient)
}

func TestMqttConnection(t *testing.T) {

	mqttSimulatedClient := &interfaces.Mqtt{}

	MqttConnection(testSettings.MqttSettings, mqttSimulatedClient)

	MqttSubscribe(MessageHandler, testSettings.MqttSettings.Topics)
}
