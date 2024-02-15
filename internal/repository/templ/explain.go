package templ

//nolint:lll
var Explain = `├── cmd/
│   ├── main.go             Initial stage of the application will run.
├── database/
│   ├── migrate/
│   │   ├── up.go           Database migrate up, 'go run database/migrate/up.go'.
│   ├── schema_migration/   Contain database migrate sql file.
├── docs/
│   ├── errors.json         Contain all errors list to be deplayed by frontend.
├── internal/
│   ├── controller/         Contain things related to gather input (REST/gRPC/console/etc) and pass input as request to service.
│   │   ├── http/
│   │   ├── grpc/
│   ├── repository/         Contain things related to persistence (CRUD database/redis/etc).
│   ├── router/
│   │   ├── injection.go    Contain dependency injection from controller to service to repository.
│   ├── service/            Contain business logic, this package get input request from controller, this package use repository for things related to persistence.
├── go.mod
└── go.sum
└── README.md
`
