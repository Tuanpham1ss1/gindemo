postgres:
	docker run --name microservicegin -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=12345 -d postgres:17.2
createdb:
	docker exec -it microservicegin createdb --username=postgres --owner=postgres gin_bank
dropdb:
	docker exec -it microservicegin dropdb gin_bank

sqlc:
	sqlc generate
server:
	go run main.go


.PHONY: postgres createdb dropdb sqlc server