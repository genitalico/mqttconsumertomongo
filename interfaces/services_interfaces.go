package interfaces

import (
	"context"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type IMongoClient interface {
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
}

type MongoClient struct {
}

func (m *MongoClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return nil
}
func (m *MongoClient) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	return nil
}

type Mongo struct {
}

func (m *Mongo) Connect(ctx context.Context, opts ...*options.ClientOptions) (IMongoClient, error) {

	mongoClient := &MongoClient{}

	return mongoClient, nil
}

// Mqtt interface

type IMqttClient interface {
	Connect() mqtt.Token
	SubscribeMultiple(filters map[string]byte, callback mqtt.MessageHandler) mqtt.Token
}

type MqttToken struct {
}

func (m *MqttToken) Wait() bool {
	return true
}

func (m *MqttToken) WaitTimeout(d time.Duration) bool {
	return true
}

func (m *MqttToken) Done() <-chan struct{} {
	return nil
}

func (m *MqttToken) Error() error {
	return nil
}

type Mqtt struct {
}

func (m *Mqtt) Connect() mqtt.Token {

	mqttToken := &MqttToken{}
	return mqttToken

}

func (m *Mqtt) SubscribeMultiple(filters map[string]byte, callback mqtt.MessageHandler) mqtt.Token {
	mqttToken := &MqttToken{}
	return mqttToken
}

func (m *Mqtt) NewClient(o *mqtt.ClientOptions) IMqttClient {

	client := &Mqtt{}

	return client
}
