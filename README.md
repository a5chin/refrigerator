# Refrigerator

## docker-compose structure
- backend:
  - go: 1.20.0
  - air: 1.43.0
  - db
    - mysql: 8.0.28

## How to develop
### Copy .env.example as .env
```sh
cp backend/config/.env.example backend/config/.env
```

### Copy this to backend/config/.env.env
```sh
# .env
HOSTNAME="127.0.0.1"
PORT="8080"
DB_HOSTNAME="db"
DB_USER="root"
DB_PWD=""
DB_NAME="contents"
DB_PORT="3306"
```

### Run backend
```sh
$ docker-compose up
```
- backend: http://localhost:8080
