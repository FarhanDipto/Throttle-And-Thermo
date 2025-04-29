package handlers

import (
	"encoding/json"
	"net/http"
	"throttle/db"
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

func CheckRateLimit(w http.ResponseWriter, r *http.Request) {
	var req RateLimitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	remaining, allowed := db.CheckRate(req.APIKey, req.Action, req.Limit, req.WindowSeconds)

	resp := RateLimitResponse{
		Allowed:   allowed,
		Remaining: remaining,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
