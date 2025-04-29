package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type RateLimitRequest struct {
	APIKey        string `json:"api_key"`
	Action        string `json:"action"`
	Limit         int    `json:"limit"`
	WindowSeconds int    `json:"window_seconds"`
}

type RateLimitResponse struct {
	Allowed   bool `json:"allowed"`
	Remaining int  `json:"remaining"`
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "Missing API Key", http.StatusUnauthorized)
		return
	}

	throttleURL := os.Getenv("THROTTLE_URL")
	if throttleURL == "" {
		throttleURL = "http://throttle:8080/check"
	}

	rateReq := RateLimitRequest{
		APIKey:        apiKey,
		Action:        "get_weather",
		Limit:         5,       
		WindowSeconds: 60,      
	}

	bodyBytes, _ := json.Marshal(rateReq)

	resp, err := http.Post(throttleURL, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Println("Error contacting Throttle:", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	var rateResp RateLimitResponse
	json.Unmarshal(respBody, &rateResp)

	if !rateResp.Allowed {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"city": "Dhaka",
		"temp": "24°C",
	})
}
