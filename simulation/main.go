package main

import (
	"fmt"
	"os"

	"github.com/Lemos1347/inteli-modulo-9-ponderada-2/simulation/publisher"
	"github.com/Lemos1347/inteli-modulo-9-ponderada-2/simulation/subscriber"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("\033[31mMissing csv path\033[0m")
		os.Exit(1)
	}
	go subscriber.RunSub("go_sub", "sensors/solar_sensor", subscriber.MessagePubHandler)
	go publisher.PubMessage("go_pub", "sensors/solar_sensor", os.Args[1], make(chan struct{}))

	select {}
}
