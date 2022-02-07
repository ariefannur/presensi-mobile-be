include .env

postgres:
	 docker run --name mypostgres -p 5432:5432 -e POSTGRES_PASSWORD=${DB_PASSWORD} -e POSTGRES_USER=${DB_USER} -d postgres:12-alpine
createdb: 
	docker exec -it mypostgres createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}
newmigate:
	migrate create -ext sql -dir db/migrate -seq ${schema}	
dropdb:
	docker exec -it mypostgres dropdb ${DB_NAME}
migrateup:
	migrate -path db/migrate -database "postgresql://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrate -database "postgresql://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -verbose down 

.PHONY: postgres createdb dropdb migrateup migratedown newmigate
