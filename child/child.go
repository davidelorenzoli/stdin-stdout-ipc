package main

import (
	"lab/parent-child-ipc/ipc"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	daemonIpc := ipc.NewDaemonIpc(os.Stdin, os.Stdout)

	log.Printf("Start waiting for messages")
	go func() {
		daemonIpc.ListenForMessages()
	}()

	log.Printf("Start sending messages")

	for i := 0; i < 20; i++ {
		message := "message-" + strconv.Itoa(i) + "\n"
		err := daemonIpc.SendMessage(message)

		if err != nil {
			log.Println("Failed to send message", err)
		}

		time.Sleep(time.Second)
	}
}
