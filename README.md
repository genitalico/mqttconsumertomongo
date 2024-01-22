# MQTT Consumer To Mongodb

Consumidor para MQTT simple que guarda todos los datos que llegan a los tópicos en una base de datos Mongodb

[Se publico una entrada sobre este código aquí.](https://80bits.blog/index.php/2024/01/22/mqtt-en-go-consumidor-que-almacena-mensajes-en-mongodb/)

## Compilación 

Se puede usar el archivo Makefile usando los siguientes targets:

- `build.linux`
- `build.windows`
- `build.darwin`

```bash
make build.linux
```

Nota: Aunque es posible compilar para Windows desde un sistema que no sea Windows, los archivos Makefile no funcionarán en Windows a menos que se utilice Windows Subsystem for Linux (WSL).

## Configurar archivo settings.json

```json
{
    "mqttSettings": {
        "mqttBrokerURL": "mqtt://localhost:1883",
        "user": "testuser",
        "password": "testpassword",
        "clientId": "testclient",
        "topics": {
            "topic1": 1,
            "topic2": 2
        }
    },
    "mongodbSettings": {
        "url": "mongodb://localhost:27017",
        "authSource": "admin",
        "username": "testuser",
        "password": "testpassword",
        "dbName": "testdb",
        "collection": "testcollection"
    }
}
```