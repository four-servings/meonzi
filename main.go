package main

import (
	"fmt"
	"github/four-servings/meonzi/local"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("app is listening on port 5000")
	// repo := exampleGetAccountRepository()
	// log.Println(repo != nil)

	pubsub := local.NewPubSub()
	pubsub.Subscribe(func(a int) {
		fmt.Println("hello world 1")
	})
	pubsub.Subscribe(func(a int) {
		fmt.Println("hello world 2")
	})
	pubsub.Publish(0)
	time.Sleep(time.Second * 10)
}
