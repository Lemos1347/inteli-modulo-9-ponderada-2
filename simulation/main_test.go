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
	err := make(chan string)
	// defer close(err)

	testTimeOut, testTimeOutCancel := context.WithTimeout(context.Background(), time.Second*20)
	defer testTimeOutCancel()

	endSub := make(chan struct{}, 1)
	endPub := make(chan struct{}, 1)

	go publisher.PubMessage("go_test_pub_1", "test/TestSendingMessages", "../data/dados_sensor_radiacao_solar.csv", endPub)

	testCallback := func(_ MQTT.Client, msg MQTT.Message) {
		if string(msg.Payload()) == "" || msg.Topic() != "test/TestSendingMessages" {
			err <- "Message didn't come"
			return
		}
		t.Log(fmt.Sprintf("Received data: %s\n", msg.Payload()))
	}

	go subscriber.RunSub("go_test_sub_1", "test/TestSendingMessages", testCallback, endSub)

	select {
	case msg := <-err:
		endSub <- struct{}{}
		endPub <- struct{}{}
		t.Fatal(msg)
	case <-testTimeOut.Done():
		endSub <- struct{}{}
		endPub <- struct{}{}
		return
	}
}

func TestMessageAcertivity(t *testing.T) {
	err := make(chan string)
	// defer close(err)

	testTimeOut, testTimeOutCancel := context.WithTimeout(context.Background(), time.Second*20)
	defer testTimeOutCancel()

	endSub := make(chan struct{}, 1)
	endPub := make(chan struct{}, 1)

	ch := make(chan string, 20)

	go publisher.PubMessage("go_test_pub_2", "test/TestMessageAcertivity", "../data/dados_sensor_radiacao_solar.csv", endPub, ch)

	testCallback := func(_ MQTT.Client, msg MQTT.Message) {
		valueSended := <-ch
		valueReceived := string(msg.Payload())

		if valueReceived != valueSended || msg.Topic() != "test/TestMessageAcertivity" {
			err <- fmt.Sprintf("Wrong message sent\nSend: %s\nReceived: %s", valueSended, valueReceived)
			return
		}
		t.Log(fmt.Sprintf("Received data: %s\n", msg.Payload()))
	}

	go subscriber.RunSub("go_test_sub_2", "test/TestMessageAcertivity", testCallback, endSub)

	select {
	case msg := <-err:
		endSub <- struct{}{}
		endPub <- struct{}{}
		t.Fatal(msg)
		return
	case <-testTimeOut.Done():
		endSub <- struct{}{}
		endPub <- struct{}{}
		return
	}
}

func TestSendingMessagesTime(t *testing.T) {
	start := time.Now()

	err := make(chan string)
	// defer close(err)

	testTimeOut, testTimeOutCancel := context.WithTimeout(context.Background(), time.Second*20)
	defer testTimeOutCancel()

	endSub := make(chan struct{}, 1)
	endPub := make(chan struct{}, 1)

	go publisher.PubMessage("go_test_pub_3", "test/TestSendingMessagesTime", "../data/dados_sensor_radiacao_solar.csv", endPub)

	testTimeCallback := func(_ MQTT.Client, msg MQTT.Message) {
		waitTime := time.Since(start).Seconds()
		limit := float64(time.Second * 6)
		if waitTime > limit {
			err <- "Message took too long"
			return
		}
		t.Log(fmt.Sprintf("Message took: %f s\n", waitTime))
		start = time.Now()
	}

	go subscriber.RunSub("go_test_sub_3", "test/TestSendingMessagesTime", testTimeCallback, endSub)

	select {
	case msg := <-err:
		endSub <- struct{}{}
		endPub <- struct{}{}
		t.Fatal(msg)
		return
	case <-testTimeOut.Done():
		endSub <- struct{}{}
		endPub <- struct{}{}
		return
	}

}
