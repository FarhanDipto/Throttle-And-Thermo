package main

import (
    // "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/redis/go-redis/v9"
    "context"
)

var ctx = context.Background()
var rdb *redis.Client

func main() {
    redisAddr := os.Getenv("REDIS_ADDR")
    if redisAddr == "" {
        redisAddr = "redis:6379" 
    }

    rdb = redis.NewClient(&redis.Options{
        Addr: redisAddr,
    })

    http.HandleFunc("/weather", weatherHandler)

    log.Println("[Thermo] Weather API running on :8081")
    http.ListenAndServe(":8081", nil)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
    city := r.URL.Query().Get("city")
    if city == "" {
        http.Error(w, "City is required as query param, e.g., ?city=Dhaka", http.StatusBadRequest)
        return
    }

    data, err := rdb.Get(ctx, city).Result()
    if err == redis.Nil {
        http.Error(w, "City not found", http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, data)
}
