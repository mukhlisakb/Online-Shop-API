# Onlie Shop Project

1. Jalankan Docker untuk database PostgreSQL

```
docker run --name onlineshopdb -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DATABASE=database -d -p 5432:5432 postgres:16
```

2. Buat Env Variable dalam terminal

```
export DB_URI=postgres://postgres:password@localhost:5432/database?sslmode=disable
```

3. Jalankan Program

```
go run .
```
