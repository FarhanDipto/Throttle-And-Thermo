# Throttle & Thermo 🌤️⚙️

This project demonstrates how I combined a **Go-based Rate Limiter** (`Throttle`) with a **Weather API Microservice** (`Thermo`) to build a rate-limited public weather service. It uses **Redis** as a backend store for both request counting and weather data caching, all containerized via Docker.

---

## 🧠 Project Background

A while ago, I built a **simple rate limiter** in Go while learning about middleware and concurrency. Recently, I revisited that code and decided to give it real-world utility by connecting it with a live weather-serving API — thus giving birth to **Throttle & Thermo**.

## 🔧 Features

- ✅ Microservice architecture with `Throttle` and `Thermo`
- ✅ Rate limiting per API key using Redis (Token Bucket strategy)
- ✅ Weather API serving **realistic weather data** by city
- ✅ Redis-backed weather store for dynamic updates (no hardcoding!)
- ✅ Docker Compose setup for easy local development
- ✅ Clear API usage with `curl` examples
- ✅ Fully forkable and customizable

## 🐳 Tech Stack

- **Go (Golang)** – core language for both microservices
- **Redis** – for both rate limiting and weather data caching
- **Docker & Docker Compose** – containerized development
- **curl** – used for simple API interaction testing

## 🗂️ Project Structure

```text
ThrottleAndThermo/
├── docker-compose.yml
├── thermo/
│   ├── main.go
│   ├── Dockerfile
├── throttle/
│   ├── main.go
│   ├── Dockerfile
├── init/
│   └── seed_redis.sh
```

## 🚀 Getting Started

### Prerequisites

- [Docker](https://www.docker.com/)
- [Git](https://git-scm.com/)

### Clone the Repo

```bash
git clone https://github.com/FarhanDipto/Throttle-And-Thermo.git
cd ThrottleAndThermo
```

### Start the Services

```bash
docker-compose up --build
```

This will:

- Start Redis
- Build and run the `Throttle` and `Thermo` services
- Auto-load sample weather data into Redis

## 🌐 API Usage

### Rate-Limited Weather Request

All requests must include an API key:

```bash
curl -X GET "http://localhost:8081/weather?city=Dhaka" -H "X-API-Key: my-key"
```

Replace `"Dhaka"` with any famous cities. If the city is not passed:

```
City is required as query param, e.g., ?city=Dhaka
```

If the API key exceeds the limit:

```
Rate limit exceeded
```

## 🧠 How It Works

- `Throttle` service intercepts all requests and **limits the number per minute per API key** using Redis.
- If allowed, the request is **proxied** to the `Thermo` service.
- `Thermo` reads the weather data for the city from Redis and returns it.
- All services are fully Dockerized and communicate over a shared Docker network.

## 📦 Persisted Weather Data

The weather information is **loaded at startup** into Redis using `preload_weather.go` in the `thermo` service. These can be updated or extended by modifying that file.

The Redis data is also stored persistently in `redis_data/` so it's not lost when containers shut down.

## 🔐 Rate Limiting Details

- **Algorithm**: Token Bucket
- **Key**: API key from header (`X-API-Key`)
- **Limit**: 5 requests per minute (adjustable in code)
- **Store**: Redis

## 🔄 Restarting the App

To run the app again after reboot:

```bash
docker-compose up
```

No rebuild is needed unless you've changed code.


## 🤝 Contribution

Feel free to fork, extend, and PR. If you're learning Go, Docker, or Redis!

## 🧹 Clean Up

To stop and remove containers and volumes:

```bash
docker-compose down -v
```