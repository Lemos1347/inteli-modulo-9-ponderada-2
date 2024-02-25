package publisher

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/Lemos1347/inteli-modulo-9-ponderada-2/simulation/publisher/sensors"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// function to generate a random sleep time
func randSleep() {
	sleepTime := rand.Intn(5) + 1
	time.Sleep(time.Duration(sleepTime) * time.Second)
}

// function to publish a messagem in a given topic
func PubMessage(topic string, csvPath string, timeOut context.Context, ch ...chan string) {
	// connecting to a broker
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// loop to emit the messages
	for {
		select {
		case <-timeOut.Done():
			fmt.Println("\033[35mPublisher encerrado! \033[0m")
			return
		default:
			// Getting the readings of a given sensor
			solarReading, err := sensors.GenerateReading(csvPath)
			if len(ch) > 0 {
				ch[0] <- solarReading
			}
			if err == nil {
				randSleep()
				fmt.Printf("\033[33mDado lido: %s  \033[0m\n", solarReading)
				token := client.Publish(topic, 1, false, solarReading)
				token.Wait()
			} else {
				fmt.Printf("\033[31m%s\033[0m\n", err.Error())
				break
			}
		}
	}

}
