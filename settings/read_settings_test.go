package settings

import (
	"encoding/json"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	testFilePath := ".././sample.settings.json"
	//defer os.Remove(testFilePath)

	testSettings := Settings{
		MqttSettings: MqttSettings{
			MqttBrokerURL: "mqtt://localhost:1883",
			User:          "testuser",
			Password:      "testpassword",
			ClientId:      "testclient",
			Topics:        map[string]byte{"topic1": 1, "topic2": 2},
		},
		MongodbSettings: MongodbSettings{
			Url:        "mongodb://localhost:27017",
			AuthSource: "admin",
			Username:   "testuser",
			Password:   "testpassword",
			DbName:     "testdb",
			Collection: "testcollection",
		},
	}

	file, err := os.Create(testFilePath)
	if err != nil {
		t.Fatalf("No se pudo crear el archivo de prueba: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(&testSettings); err != nil {
		t.Fatalf("No se pudo escribir la configuración de prueba en el archivo: %v", err)
	}

	settings, err := ReadFile(&testFilePath)
	if err != nil {
		t.Fatalf("Error al leer el archivo de configuración: %v", err)
	}

	if settings.MqttSettings.MqttBrokerURL != testSettings.MqttSettings.MqttBrokerURL ||
		settings.MqttSettings.User != testSettings.MqttSettings.User ||
		settings.MqttSettings.Password != testSettings.MqttSettings.Password ||
		settings.MqttSettings.ClientId != testSettings.MqttSettings.ClientId ||
		len(settings.MqttSettings.Topics) != len(testSettings.MqttSettings.Topics) ||
		settings.MongodbSettings.Url != testSettings.MongodbSettings.Url ||
		settings.MongodbSettings.AuthSource != testSettings.MongodbSettings.AuthSource ||
		settings.MongodbSettings.Username != testSettings.MongodbSettings.Username ||
		settings.MongodbSettings.Password != testSettings.MongodbSettings.Password ||
		settings.MongodbSettings.DbName != testSettings.MongodbSettings.DbName ||
		settings.MongodbSettings.Collection != testSettings.MongodbSettings.Collection {
		t.Errorf("La configuración leída no coincide con la configuración de prueba")
	}
}
