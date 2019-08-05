package ipc

import (
	"bufio"
	"io"
	"log"
)

type DaemonIpc struct {
	reader io.Reader
	writer io.Writer
}

type MessageHandler func(message string)

/*
NewDaemonIpc create an instance of NewDaemonIpc
		reader reads from a data source
		writer writes to a data source
*/
func NewDaemonIpc(reader io.Reader, writer io.Writer) DaemonIpc {
	return DaemonIpc{reader, writer}
}

/*
sendMessage not blocking function that sends a message to the standard output
		message the message to be sent
*/
func (daemonIpc DaemonIpc) SendMessage(message string) error {
	_, err := daemonIpc.writer.Write([]byte(message))

	if err != nil {
		log.Println("Failed to send message", err)
	}

	return err
}

// listenForMessages blocking function that listens for messages from the reader
func (daemonIpc DaemonIpc) ListenForMessages(messageHandler MessageHandler) {
	reader := bufio.NewReader(daemonIpc.reader)

	for hasNext := true; hasNext; {
		line, _ := reader.ReadString('\n')

		hasNext = line != ""

		if hasNext {
			messageHandler(line)
		}
	}
}
