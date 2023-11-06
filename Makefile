run:
	RDB_ADDR=localhost:6379 RDB_PWD= PG_HOST=localhost PG_PORT=5433 PG_PWD=root PG_USER=postgres PG_DB=postgres go run cmd/budgeter/main.go