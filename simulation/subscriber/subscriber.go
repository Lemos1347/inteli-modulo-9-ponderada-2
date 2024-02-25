package subscriber

import (
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var MessagePubHandler MQTT.MessageHandler = func(_ MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Recebido dado solar: %s do tópico: %s as %s \n", msg.Payload(), msg.Topic(), time.Now().Format(time.RFC3339))
}

func RunSub(topic string, callback MQTT.MessageHandler) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_subscriber")
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
	select {}
}

