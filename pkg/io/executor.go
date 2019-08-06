package io

import (
	"lab/parent-child-ipc/pkg/ipc"
	"log"
	"os/exec"
)

// execute run the given executable. In case of error it calls `os.Exit(1)`
func Execute(executablePath string) ipc.DaemonIpc {
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
