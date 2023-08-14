# User application.

Designed to work with users.

Implemented:
 CRUD
 Migrations
 LoadBalancer
 Configs
 
Testing ___in process
___
## Application launch:
```
docker compose up
```
## Test query file :
```
 testRequest.http
```
___
Design:
- REST API
- Clean Arcitecture

Technologies usage:
- Docker
- Nginx
- PostgreSql

Frameworks:
- PostgresCLient <a href="github.com/jackc/pgx/v5">pgx</a>
- Gin <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>
- Logrus <a href="github.com/sirupsen/logrus">logrus</a>
- Migrations <a href="github.com/golang-migrate/migrate/v4">migrate</a>
