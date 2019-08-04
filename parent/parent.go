package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

func main() {
	executablePath, err := os.Executable()
	executableDir := path.Dir(executablePath)
	childExecutablePath := path.Join(executableDir, "child")

	command := exec.Command(childExecutablePath)
	pipeReader, _ := command.StdoutPipe()

	log.Printf("Launching executable %s", command.Path)
	err = command.Start()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Process started. PID %d", command.Process.Pid)

	readProcessOutput(pipeReader)

	log.Printf("Terminated")
}

func readProcessOutput(readCloser io.ReadCloser) {
	reader := bufio.NewReader(readCloser)

	for hasNext := true; hasNext; {
		line, _ := reader.ReadString('\n')

		hasNext = line != ""

		if hasNext {
			log.Printf("Received message: %s", string(line))
		}
	}
}
