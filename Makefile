postgres:
	docker run --name trustbank -p 5432:5432 -e POSTGRES_USER=trustbank -e POSTGRES_PASSWORD=trustapi -d postgres:alpine3.18
 
createdb:
	docker exec -it trustbank createdb --username=trustbank --owner=trustbank trustbankdb

migrateup: 
	migrate -path db/migration -database "postgresql://trustbank:trustapi@localhost:5432/trustbankdb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://trustbank:trustapi@localhost:5432/trustbankdb?sslmode=disable" -verbose down

migrateclean:
	migrate -path db/migration -database "postgresql://trustbank:trustapi@localhost:5432/trustbankdb?sslmode=disable" -verbose force 0000001

dropdb:
	docker exec -it trustbank dropdb --username=trustbank --owner=trustbank trustbankdb

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

sqlcwin:
	docker run --rm -v C:\Users\JosephLamptey\Desktop\Projects\trustbank:/src -w /src sqlc/sqlc generate

.PHONY:createdb postgres migrateup migratedown sqlc test