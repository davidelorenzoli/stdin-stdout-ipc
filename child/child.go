package main

import (
	"os"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		message := "message-" + strconv.Itoa(i) + "\n"
		sendMessage(message)
		time.Sleep(time.Second)
	}
}

func sendMessage(message string) {
	_, err := os.Stdout.Write([]byte(message))

	if err != nil {
		_, _ = os.Stderr.Write([]byte(err.Error()))
	}
}
