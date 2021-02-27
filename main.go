package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("app is listening on port 5000")
	repo := exampleGetAccountRepository()
	log.Println(repo != nil)
}
