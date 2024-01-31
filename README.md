# gogger (go + logger)
Distributed Logging
#### - producer server 📋
[![Build Status](https://github.com/Slowhigh/gogger/actions/workflows/ci-producer.yml/badge.svg?branch=main)](https://github.com/features/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/Slowhigh/gogger/producer)](https://goreportcard.com/report/github.com/Slowhigh/gogger/producer)
[![codebeat badge](https://codebeat.co/badges/ecde5eef-54fa-412f-bbe4-3b63e488809c)](https://codebeat.co/projects/github-com-slowhigh-gogger-producer)
#### - consumer server 📋
[![Build Status](https://github.com/Slowhigh/gogger/actions/workflows/ci-consumer.yml/badge.svg?branch=main)](https://github.com/features/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/Slowhigh/gogger/consumer)](https://goreportcard.com/report/github.com/Slowhigh/gogger/consumer)
[![codebeat badge](https://codebeat.co/badges/128b9146-2e6f-4490-8346-3bf7f17c6bf5)](https://codebeat.co/projects/github-com-slowhigh-gogger-consumer)

## Quick start 🚀
#### 1. Deploy Memphis Broker
```bash
$ curl -s https://memphisdev.github.io/memphis-docker/docker-compose.yml -o docker-compose.yml \
&& docker compose -f docker-compose.yml -p memphis up -d
```

#### 2. Deploy Postgres
```bash
$ docker run -d \
-p 5432:5432 \
-e POSTGRES_USER=gogger \
-e POSTGRES_PASSWORD=gogger1! \
-e POSTGRES_DB=gogger \
--name postgres postgres:alpine
```

#### 3. Run Consumer Server
```bash
$ cd ./consumer/ && go run ./cmd/server/

# 2024/01/28 23:13:12 INFO start consuming messages
```

#### 4. Run Producer Server (new terminal)
```bash
$ cd ./producer/ && go run ./cmd/server/

# [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
# 
# [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
#  - using env:   export GIN_MODE=release
#  - using code:  gin.SetMode(gin.ReleaseMode)
# 
# [GIN-debug] POST   /log/access               --> github.com/Slowhigh/gogger/producer/infra/router.NewRouter.func1 (3 handlers)
# [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
# Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details. 
# [GIN-debug] Listening and serving HTTP on :5000
```

#### 5. Send the HTTP message to the Producer (new terminal)
```bash
$ curl --location 'http://localhost:5000/log/access' \
--header 'Content-Type: application/json' \
--data '{
    "timestamp": "2006-01-02T15:04:05Z",
    "is_normal_mode": true,
    "is_login": true,
    "user_name": "john",
    "device_name": "slowhigh",
    "ip": "192.168.0.1"
}'
```

#### 6. Check the message on the Station page
- initial account - id: `root` / pw: `memphis`
- go to http://localhost:9000/stations/access-message
  ![image](https://github.com/Slowhigh/gogger/assets/37216082/2462a2e5-e428-4aac-a9d9-6f56f8a19e84)


