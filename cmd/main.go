package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	server, err := newServer()
	if err != nil {
		logrus.Fatalf("falied to create server: %v", err)
	}

	server.start()
}
