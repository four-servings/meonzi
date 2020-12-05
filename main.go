package main

import (
	_ "github/four-servings/meonzi/accounts"
	"log"
	"net/http"
)

func main() {
	log.Println("app is listening on port 5000")
	http.ListenAndServe(":5000", nil)
}
