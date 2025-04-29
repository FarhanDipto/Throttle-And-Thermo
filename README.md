```markdown
# Throttle & Thermo 🌤️⚙️

This project demonstrates how I combined a **Go-based Rate Limiter** (`Throttle`) with a **Weather API Microservice** (`Thermo`) to build a rate-limited public weather service. It uses **Redis** as a backend store for both request counting and weather data caching, all containerized via Docker.

---

## 🧠 Project Background

A while ago, I built a **simple rate limiter** in Go while learning about middleware and concurrency. Recently, I revisited that code and decided to give it real-world utility by connecting it with a live weather-serving API — thus giving birth to **Throttle & Thermo**.

---

## 🔧 Features

- ✅ Microservice architecture with `Throttle` and `Thermo`
- ✅ Rate limiting per API key using Redis (Token Bucket strategy)
- ✅ Weather API serving **realistic weather data** by city
- ✅ Redis-backed weather store for dynamic updates (no hardcoding!)
- ✅ Docker Compose setup for easy local development
- ✅ Clear API usage with `curl` examples
- ✅ Fully forkable and customizable

---

## 🐳 Tech Stack

- **Go (Golang)** – core language for both microservices
- **Redis** – for both rate limiting and weather data caching
- **Docker & Docker Compose** – containerized development
- **curl** – used for simple API interaction testing

---

## 🗂️ Project Structure

```
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

---

## 🚀 Getting Started

### 1. Prerequisites

- [Docker](https://www.docker.com/)
- [Git](https://git-scm.com/)

### 2. Clone the Repo

```bash
git clone https://github.com/your-username/ThrottleAndThermo.git
cd ThrottleAndThermo
```

### 3. Start the Services

```bash
docker-compose up --build
```

This will:

- Start Redis
- Build and run the `Throttle` and `Thermo` services
- Auto-load sample weather data into Redis

---

## 🌐 API Usage

### 🔑 Rate-Limited Weather Request

All requests must include an API key:

```bash
curl -X GET "http://localhost:8081/weather?city=Dhaka" -H "X-API-Key: my-key"
```

Replace `"Dhaka"` with any supported city:
- Dhaka
- Tokyo
- London
- Paris
- New York
- Delhi

If the city is not passed:
```
City is required as query param, e.g., ?city=Dhaka
```

If the API key exceeds the limit:
```
Rate limit exceeded
```

---

## 🧠 How It Works

- `Throttle` service intercepts all requests and **limits the number per minute per API key** using Redis.
- If allowed, the request is **proxied** to the `Thermo` service.
- `Thermo` reads the weather data for the city from Redis and returns it.
- All services are fully Dockerized and communicate over a shared Docker network.

---

## 📦 Persisted Weather Data

The weather information is **loaded at startup** into Redis using `preload_weather.go` in the `thermo` service. These can be updated or extended by modifying that file.

The Redis data is also stored persistently in `redis_data/` so it's not lost when containers shut down.

---

## 🔄 Restarting the App

To run the app again after reboot:

```bash
docker-compose up
```

No rebuild is needed unless you've changed code.

---

## 🔐 Rate Limiting Details

- **Algorithm**: Token Bucket
- **Key**: API key from header (`X-API-Key`)
- **Limit**: 5 requests per minute (adjustable in code)
- **Store**: Redis

---

## 💡 Future Ideas

- Add a real weather API data puller (e.g., OpenWeatherMap)
- Track usage per key and provide dashboard
- Add authentication and user registration

---

## 🤝 Contribution

Feel free to fork, extend, and PR. If you're learning Go, Docker, or Redis — this repo is made for you!

---

## 🧹 Clean Up

To stop and remove containers and volumes:

```bash
docker-compose down -v
```

---

## 📬 Contact

Built with ❤️ by [Your Name](mailto:your.email@example.com)

```

---

Let me know if you want me to generate `weather.sh` or a preload example for another city.