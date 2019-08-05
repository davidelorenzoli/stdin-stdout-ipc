package main

import (
	"lab/parent-child-ipc/ipc"
	"log"
	"os"
	"strconv"
	"time"
)

var daemonIpc ipc.DaemonIpc

func main() {
	daemonIpc = ipc.NewDaemonIpc(os.Stdin, os.Stdout)

	log.Printf("Start waiting for messages")

	go daemonIpc.ListenForMessages(messageHandler)

	log.Printf("Start sending messages")

	for i := 0; i < 20; i++ {
		message := "message-" + strconv.Itoa(i) + "\n"
		err := daemonIpc.SendMessage(message)

		if err != nil {
			log.Printf("Failed to send message. Error %s", err)
		}

		time.Sleep(time.Second)
	}
}

func messageHandler(message string) {
	log.Printf("Received message: %s", string(message))

	err := daemonIpc.SendMessage("ack " + message)

	if err != nil {
		log.Printf("Failed to send message. Error %s", err)
	}
}
