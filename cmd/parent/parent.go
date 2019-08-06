package main

import (
	"lab/parent-child-ipc/pkg/io"
	"lab/parent-child-ipc/pkg/ipc"
	"log"
	"os"
	"path"
	"strconv"
	"time"
)

var daemonIpc ipc.DaemonIpc

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	executablePath, _ := os.Executable()
	executableDir := path.Dir(executablePath)
	childExecutablePath := path.Join(executableDir, "child")

	daemonIpc = io.Execute(childExecutablePath)

	go daemonIpc.ListenForMessages(messageHandler)

	sendMessages()

	log.Printf("Execution completed")
}

func sendMessages() {
	for i := 0; i < 10; i++ {
		message := "parent-" + strconv.Itoa(i) + "\n"
		err := daemonIpc.SendMessage(message)

		if err != nil {
			log.Printf("Failed to send message. Error %s", err)
		}

		time.Sleep(2 * time.Second)
	}
}

func messageHandler(message string) {
	log.Printf("Received message: %s", message)
}
