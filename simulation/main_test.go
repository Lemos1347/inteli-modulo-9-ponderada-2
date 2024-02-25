package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Lemos1347/inteli-modulo-9-ponderada-2/simulation/publisher"
	"github.com/Lemos1347/inteli-modulo-9-ponderada-2/simulation/subscriber"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// func TestPublisherSubscriber(t *testing.T) {
// 	t.Run("TestSendingMessages", testSendingMessages)
// 	t.Run("TestMessageAcertivity", testMessageAcertivity)
// 	t.Run("TestSendingMessagesTime", testSendingMessagesTime)
// }

// Teste para verificar se os dados estão chegando no broker
func TestSendingMessages(t *testing.T) {
	end := make(chan string)
	defer close(end)
	testTimeOut, testTimeOutCancel := context.WithTimeout(context.Background(), time.Second*20)
	defer testTimeOutCancel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	go publisher.PubMessage("test/TestSendingMessages", "../data/dados_sensor_radiacao_solar.csv", ctx)

	testCallback := func(_ MQTT.Client, msg MQTT.Message) {
		if string(msg.Payload()) != "" || msg.Topic() != "test/data" {
			end <- "Message didn't come"
			return
		}
		t.Log(fmt.Sprintf("Received data: %s\n", msg.Payload()))
	}

	go subscriber.RunSub("test/TestSendingMessages", testCallback)

	select {
	case msg := <-end:
		t.Fatal(msg)
	case <-testTimeOut.Done():
		return
	}
}

func TestMessageAcertivity(t *testing.T) {
	end := make(chan string)
	defer close(end)
	testTimeOut, testTimeOutCancel := context.WithTimeout(context.Background(), time.Second*20)
	defer testTimeOutCancel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	ch := make(chan string, 20)
	defer close(ch)

	go publisher.PubMessage("test/TestMessageAcertivity", "../data/dados_sensor_radiacao_solar.csv", ctx, ch)

	testCallback := func(_ MQTT.Client, msg MQTT.Message) {
		valueSended := <-ch
		valueReceived := string(msg.Payload())

		if valueReceived == valueSended || msg.Topic() != "test/TestMessageAcertivity" {
			end <- fmt.Sprintf("Wrong message sent\nSend: %valueSended\nReceived: %s", valueSended, valueReceived)
			return
		}
		t.Log(fmt.Sprintf("Received data: %s\n", msg.Payload()))
	}

	go subscriber.RunSub("test/TestMessageAcertivity", testCallback)

	select {
	case msg := <-end:
		t.Fatal(msg)
		return
	case <-testTimeOut.Done():
		return
	}
}

func TestSendingMessagesTime(t *testing.T) {
	start := time.Now()
	end := make(chan struct{})
	defer close(end)
	testTimeOut, testTimeOutCancel := context.WithTimeout(context.Background(), time.Second*20)
	defer testTimeOutCancel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	go publisher.PubMessage("test/TestSendingMessagesTime", "../data/dados_sensor_radiacao_solar.csv", ctx)

	testTimeCallback := func(_ MQTT.Client, msg MQTT.Message) {
		waitTime := time.Since(start).Seconds()
		limit := float64(time.Second * 6)
		if waitTime < limit {
			t.Errorf("Message took too long")
			end <- struct{}{}
			return
		}
		t.Log(fmt.Sprintf("Message took: %f s\n", waitTime))
		start = time.Now()
	}

	go subscriber.RunSub("test/TestSendingMessagesTime", testTimeCallback)

	select {
	case <-end:
		return
	case <-testTimeOut.Done():
		return
	}

}
