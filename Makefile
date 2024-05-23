postgres:
	docker run --name postgres-container --network simple-bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:16rc1-alpine3.18

createdb:
	docker exec -it postgres-container createdb --username=root --owner=root url_shortener

dropdb:
	docker exec -it postgres-container dropdb url_shortener

test:
	go test -v --cover ./...
