# Final Task

## Система "Платежі"

### Used:

- server: github.com/labstack/echo
- orm: github.com/uptrace/bun
- test: github.com/stretchr/testify

### Done:

- Domen-driven Design
- BUN
- logs
- e2e testing coverage
- postgreSQL
- validation
- documentation
- error handling
- containerization
- graceful shutdown

### To be Done:

- db transactions
- pagination, sorting
- authentication

### Start:

- **run docker**:
  `docker run --name payments-db -p 5455:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=payments -d postgres
  `
- **run application**: `source .env && go run cmd/main.go`

- **run e2e**: `go test -v tests/e2e/*`

