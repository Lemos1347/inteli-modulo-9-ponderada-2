package subscriber

import (
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var MessagePubHandler MQTT.MessageHandler = func(_ MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Recebido dado solar: %s do tópico: %s as %s \n", msg.Payload(), msg.Topic(), time.Now().Format(time.RFC3339))
}

func RunSub(clientId string, topic string, callback MQTT.MessageHandler, end ...chan struct{}) {
	if len(end) == 0 {
		end = []chan struct{}{
			make(chan struct{}),
		}
	}

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID(clientId)
	opts.SetDefaultPublishHandler(callback)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	fmt.Println("Subscriber está rodando. Pressione CTRL+C para sair.")
	select {
	case <-end[0]:
		close(end[0])
		return
	}
}
