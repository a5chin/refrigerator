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
  - 全素材取得 API: http://localhost:8080/api/v1/ingredients?min=10&max=100
    - min（Optional）: 重さの最小値
    - max（Optional）: 重さの最大値
  ```json
  {
    "ingredients": [
      {
        "id": "string",
        "name": "string",
        "nutritions": [
          {
            "id": "string",
            "name": "string"
          }
        ],
        "weight": 0
      }
    ]
  }
  ```
  - 素材取得 API: http://localhost:8080/api/v1/ingredients/01Z02KW4CH0QAYDWE2D26X0JVD
  ```json
  {
    "ingredient": {
      "id": "string",
      "name": "string",
      "nutritions": [
        {
          "id": "string",
          "name": "string"
        }
      ],
      "weight": 0
    }
  }
  ```
  - 素材更新 API: http://localhost:8080/api/v1/ingredients/01Z02KW4CH0QAYDWE2D26X0JVD?weight=80
    - weight: 更新後する重さ
  ```json
  {
    "message": "string"
  }
  ```
  - 全栄養素取得 API: http://localhost:8080/api/v1/nutritions
  ```json
  {
    "nutritions": [
      {
        "id": "string",
        "name": "string"
      }
    ]
  }
  ```
  - 栄養素取得 API: http://localhost:8080/api/v1/nutritions/01H3M84GCPQWTACHPE77WTJ20Y
  ```json
  {
    "nutritions": [
      {
        "id": "string",
        "name": "string"
      }
    ]
  }
  ```
