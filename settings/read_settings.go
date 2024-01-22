package settings

import (
	"encoding/json"
	"os"
)

type Settings struct {
	MqttSettings    MqttSettings    `json:"mqttSettings"`
	MongodbSettings MongodbSettings `json:"mongodbSettings"`
}

type MqttSettings struct {
	MqttBrokerURL string          `json:"mqttBrokerURL"`
	User          string          `json:"user"`
	Password      string          `json:"password"`
	ClientId      string          `json:"clientId"`
	Topics        map[string]byte `json:"topics"`
}

type MongodbSettings struct {
	Url        string `json:"url"`
	AuthSource string `json:"authSource"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	DbName     string `json:"dbName"`
	Collection string `json:"collection"`
}

func ReadFile(path *string) (Settings, error) {

	if path == nil {
		path = new(string)
		*path = "settings.json"
	}

	file, err := os.Open(*path)
	if err != nil {
		return Settings{}, err
	}
	defer file.Close()

	var settings Settings
	err = json.NewDecoder(file).Decode(&settings)
	if err != nil {
		return Settings{}, err
	}

	return settings, nil
}
