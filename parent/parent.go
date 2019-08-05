package main

import (
	"lab/parent-child-ipc/ipc"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	executablePath, _ := os.Executable()
	executableDir := path.Dir(executablePath)
	childExecutablePath := path.Join(executableDir, "child")

	daemonIpc := execute(childExecutablePath)

	go daemonIpc.ListenForMessages(messageHandler)

	for i := 0; i < 10; i++ {
		message := "parent-" + strconv.Itoa(i) + "\n"
		err := daemonIpc.SendMessage(message)

		if err != nil {
			log.Printf("Failed to send message. Error %s", err)
		}

		time.Sleep(2 * time.Second)
	}

	log.Printf("Execution completed")
}

func messageHandler(message string) {
	log.Printf("Received message: %s", message)
}

// execute run the given executable. In case of error it calls `os.Exit(1)`
func execute(executablePath string) ipc.DaemonIpc {
	command := exec.Command(executablePath)
	pipeReader, err := command.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	pipeWriter, err := command.StdinPipe()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Launching executable %s", command.Path)

	err = command.Start()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Process started. PID %d", command.Process.Pid)

	return ipc.NewDaemonIpc(pipeReader, pipeWriter)
}
