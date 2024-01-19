# Distributed Logging

### Quick start
#### 1. Deploy Memphis Broker
```bash
$ curl -s https://memphisdev.github.io/memphis-docker/docker-compose.yml -o docker-compose.yml \
&& docker compose -f docker-compose.yml -p memphis up -d
```

#### 2. Deploy Postgres
```bash
$ docker run -d \
-p 5433:5432 \
-e POSTGRES_USER=gogger \
-e POSTGRES_PASSWORD=gogger1! \
-e POSTGRES_DB=gogger \
--name postgres postgres:alpine
```

$ curl --location --request POST 'http://127.0.0.1:4444/auth/authenticate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "root",
    "password": "memphis",
    "token_expiry_in_minutes": 6000000,
    "refresh_token_expiry_in_minutes": 100000
}'

{"expires_in":123932266980000,"jwt":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoxLCJjb25uZWN0aW9uX3Rva2VuIjoiIiwiZXhwIjoyMDY1NTM3NzgzLCJwYXNzd29yZCI6Im1lbXBoaXMiLCJ1c2VybmFtZSI6InJvb3QifQ.FYwGMcX93weZlT3QBq5Rv-jHtEu7JtDo8bOV1QXKyC0","jwt_refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoxLCJjb25uZWN0aW9uX3Rva2VuIjoiIiwiZXhwIjoxNzExNTM3NzgzLCJwYXNzd29yZCI6Im1lbXBoaXMiLCJ0b2tlbl9leHAiOjIwNjU1Mzc3ODMsInVzZXJuYW1lIjoicm9vdCJ9.4Wmkp9sWAlfeQeSHPVv1TcVUquUIulBYkyqvA_w_KL4","refresh_token_expires_in":123932266980000}
```