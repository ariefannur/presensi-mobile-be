### Presensi Mobile Api

#### Docker 
```
docker run --rm -d \
    --name dev-postgres \
    --network presensi-mobile\
    -e POSTGRES_USER=admin \
    -e POSTGRES_PASSWORD=password \
    -e POSTGRES_DB=presensi_db \
    -v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
    -p 5432:5432 \
    postgres
    ```