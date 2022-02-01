postgres:
	 docker run --name mypostgres -p 5432:5432 -e POSTGRES_PASSWORD=presensi_0ke_Deh -e POSTGRES_USER=root -d postgres:12-alpine
createdb: 
	docker exec -it mypostgres createdb --username=root --owner=root presensi_sekolah
dropdb:
	docker exec -it mypostgres dropdb presensi_sekolah
migrateup:
	migrate -path db/migrate -database "postgresql://root:presensi_0ke_Deh@localhost:5432/presensi_sekolah?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrate -database "postgresql://root:presensi_0ke_Deh@localhost:5432/presensi_sekolah?sslmode=disable" -verbose down 

.PHONY: postgres createdb dropdb migrateup migratedown
