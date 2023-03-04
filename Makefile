run_postgres:
	docker run --name arxivhub -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres:15-alpine3.17

create_db:
	docker exec -it postgres createdb --username=root --owner=root arxivhub

drop_db:
	docker exec -it postgres dropdb arxivhub

migrate_up:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/arxivhub?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/arxivhub?sslmode=disable" -verbose down

generate_sqlc:
	sqlc generate