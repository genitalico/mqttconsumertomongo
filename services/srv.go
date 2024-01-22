package services

import (
	"context"
	"encoding/json"
	"fmt"
	"mqttconsumertomongo/interfaces"
	"mqttconsumertomongo/settings"

	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient interfaces.IMongoClient
var settingsValues settings.Settings
var mqttClient interfaces.IMqttClient

func SetSettingsValues(settings settings.Settings) {
	settingsValues = settings
}

func MessageHandler(client mqtt.Client, msg mqtt.Message) {

	var data = map[string]interface{}{}
	err := json.Unmarshal(msg.Payload(), &data)

	if err != nil {
		data["topic"] = msg.Topic()
		data["payload"] = string(msg.Payload())
		mongoClient.Database(settingsValues.MongodbSettings.DbName).Collection(settingsValues.MongodbSettings.Collection).InsertOne(context.Background(), data)

		return
	}
	data["topic"] = msg.Topic()
	collectionValue, isCollection := data["collection"]
	if isCollection {
		collection := fmt.Sprintf("%v", collectionValue)
		mongoClient.Database(settingsValues.MongodbSettings.DbName).Collection(string(collection)).InsertOne(context.Background(), data)
	} else {
		mongoClient.Database(settingsValues.MongodbSettings.DbName).Collection(settingsValues.MongodbSettings.Collection).InsertOne(context.Background(), data)
	}
}

func MqttConnection(mqttSettings settings.MqttSettings, mqttt *interfaces.Mqtt) (err error) {

	mqttBrokerURL := mqttSettings.MqttBrokerURL
	user := mqttSettings.User
	password := mqttSettings.Password
	clientId := mqttSettings.ClientId
	opts := mqtt.NewClientOptions()
	opts.AddBroker(mqttBrokerURL)
	opts.SetClientID(clientId)
	opts.SetUsername(user)
	opts.SetPassword(password)
	opts.SetKeepAlive(10 * time.Second)
	opts.SetAutoReconnect(true)

	if mqttt == nil {
		mqttClient = mqtt.NewClient(opts)
	} else {
		mqttClient = mqttt.NewClient(opts)
	}

	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func MqttSubscribe(messageHandler mqtt.MessageHandler, topics map[string]byte) (err error) {

	token := mqttClient.SubscribeMultiple(topics, MessageHandler)

	if token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func MongoConnection(mongodbSettings settings.MongodbSettings, mongodb *interfaces.Mongo) (err error) {
	authSource := mongodbSettings.AuthSource
	username := mongodbSettings.Username
	password := mongodbSettings.Password
	url := mongodbSettings.Url
	clientOptions := options.Client().ApplyURI(url).
		SetAuth(options.Credential{
			AuthSource: authSource,
			Username:   username,
			Password:   password,
		})

	// Connect to MongoDB

	if mongodb == nil {
		mongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	} else {
		mongoClient, err = mongodb.Connect(context.TODO(), clientOptions)
	}

	if err != nil {
		return err
	}

	// Check the connection
	err = mongoClient.Ping(context.TODO(), nil)

	if err != nil {
		return err
	}

	return nil
}
