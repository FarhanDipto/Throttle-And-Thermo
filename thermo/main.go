package main

import (
	"log"
	"net/http"
	"thermo/handlers"
)

func main() {
	http.HandleFunc("/weather", handlers.GetWeather)

	log.Println("[Thermo] Weather API running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
