### Presensi Mobile Api

#### Docker
```
// for run all env and start service
docker compose up
// for delete stop docker container
docker compose down 
```

#### Makefile 
```
// create postgres
make postgres
// create db
make createdb
// migrate up
make migrateup
// migrate down
make migratedown
```

#### Database Diagram
![ERD](/Presensi%20Diagram.png)

#### Environment Variable
```
DB_NAME=
DB_PASSWORD=
DB_USER=

DB_SERVER_URL=postgres://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}
DB_MAX_CONNECTIONS=00
DB_MAX_IDLE_CONNECTIONS=00
DB_MAX_LIFETIME_CONNECTIONS=00

GROUP_PATH="api/v1"

JWT_SECRET=
JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=00
```
