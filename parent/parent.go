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
	executablePath, _ := os.Executable()
	executableDir := path.Dir(executablePath)
	childExecutablePath := path.Join(executableDir, "child")

	processOutputReader := execute(childExecutablePath)

	read(processOutputReader)

	log.Printf("Execution completed")
}

// execute run the given executable. In case of error it calls `os.Exit(1)`
func execute(executablePath string) io.ReadCloser {
	command := exec.Command(executablePath)
	pipeReader, _ := command.StdoutPipe()

	log.Printf("Launching executable %s", command.Path)

	err := command.Start()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Process started. PID %d", command.Process.Pid)

	return pipeReader
}

// read read from ReadCloser and log its content
func read(readCloser io.ReadCloser) {
	reader := bufio.NewReader(readCloser)

	for hasNext := true; hasNext; {
		line, _ := reader.ReadString('\n')

		hasNext = line != ""

		if hasNext {
			log.Printf("Received message: %s", string(line))
		}
	}
}
