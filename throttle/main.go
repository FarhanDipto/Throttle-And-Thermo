package main

import (
	"log"
	"net/http"
	"throttle/handlers"
)

func main() {
	http.HandleFunc("/check", handlers.CheckRateLimit)

	log.Println("[Throttle] Rate Limiter API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
