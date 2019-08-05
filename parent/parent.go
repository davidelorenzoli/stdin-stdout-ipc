package main

import (
	"lab/parent-child-ipc/ipc"
	"log"
	"os"
	"os/exec"
	"path"
)

func main() {
	executablePath, _ := os.Executable()
	executableDir := path.Dir(executablePath)
	childExecutablePath := path.Join(executableDir, "child")

	daemonIpc := execute(childExecutablePath)

	daemonIpc.ListenForMessages(messageHandler)
	log.Printf("Execution completed")
}

func messageHandler(message string) {
	log.Printf("Received message: %s", string(message))
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
