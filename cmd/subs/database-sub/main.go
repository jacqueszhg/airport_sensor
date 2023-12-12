package main

import (
	"Airport/internal/pkg/config"
	mqttConfig "Airport/internal/pkg/mqtt"
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	configSub := config.GetSubonfig("./databaseConfig.yml")

	// Initialize Influx DB
	db := createDb()
	writeApi := createWriteAPI(db)

	// Connect to mqtt
	client := mqttConfig.Connect(configSub.MQTT.Protocol+"://"+configSub.MQTT.Url+":"+configSub.MQTT.Port, configSub.MQTT.Id)

	// Subscribe to all sensors
	client.Subscribe("airport/temperature", byte(configSub.MQTT.QOSLevel), func(client mqtt.Client, message mqtt.Message) {
		onDataReceived(message, writeApi)
	})

	client.Subscribe("airport/wind", byte(configSub.MQTT.QOSLevel), func(client mqtt.Client, message mqtt.Message) {
		onDataReceived(message, writeApi)
	})

	client.Subscribe("airport/pressure", byte(configSub.MQTT.QOSLevel), func(client mqtt.Client, message mqtt.Message) {
		onDataReceived(message, writeApi)
	})

	fmt.Printf("finish")
	wg.Wait()
	db.Close()
}

/*
 * Creating the influx client with the corresponding token and url to connect with database
 */
func createDb() influxdb.Client {
	token := "TOKEN_INFLUX_DB"
	url := "URL_INFLUX_DB"
	return influxdb.NewClient(url, token)
}

/*
 * Creating a WriteAPIBlocking with our database and organization/bucket
 */
func createWriteAPI(db influxdb.Client) api.WriteAPIBlocking {
	org := "airport"
	bucket := "Sensors"
	return db.WriteAPIBlocking(org, bucket)
}

/*
 * On any data received (from every sensors), add a point into the database
 */
func onDataReceived(message mqtt.Message, api api.WriteAPIBlocking) {
	var r mqttConfig.MessageSensorPublisher
	err := json.Unmarshal(message.Payload(), &r)
	if err != nil {
		fmt.Println("Can't deserislize", message.Payload())
	}

	tags := map[string]string{
		"airport": r.AirportCode,
		"id":      fmt.Sprintf("%d", r.SensorId),
	}

	fields := map[string]interface{}{
		"value": r.Value,
	}

	point := write.NewPoint(r.SensorType, tags, fields, time.Now())

	if err := api.WritePoint(context.Background(), point); err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)
	fmt.Println("Sent!")
}
