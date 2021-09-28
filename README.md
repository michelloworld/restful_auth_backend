# RESTFul Authentication Backend

### How to run with docker compose

1. update .env file

```
ACCESS_TOKEN_KEY=YOUR_TOKEN_EKY
ACCESS_TOKEN_LIFE=600
REFRESH_TOKEN_KEY=YOUR_REFRESH_KEY
REFRESH_TOKEN_LIFE=86400

DB_HOST=db
DB_PORT=5432
DB_USER=my_db
DB_PASSWORD=password
DB_NAME=restful_auth

DOCKER_DB_USER=my_db
DOCKER_DB_PASSWORD=password
DOCKER_DB_NAME=restful_auth
```

2. run docker compose

```
docker-compose -f docker-compose.dev.yml up --build
```

3. run migration

```
docker exec -it restful-auth-backend-app-dev go run migrations/migration/migration.go
```

4. run seed

```
docker exec -it restful-auth-backend-app-dev go run migrations/seed/seed.go
```

5. access

```
http://localhost:4000
```