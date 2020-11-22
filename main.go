package main

import (
	_ "github/four-servings/meonzi/accounts"
	"net/http"
)

func main() {
	http.ListenAndServe(":5000", nil)
}
