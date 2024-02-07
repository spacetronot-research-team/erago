package explain

import "fmt"

func Explain() {
	fmt.Println(explainText)
}

var explainText = `
├── cmd/                Initial stage of the application will run.
├── internal/           Core package of the application and contains the implementation of various business logic.
│   ├──  controller/    This package is only to gather input (REST/gRPC/console/etc) and pass input as request to service.
│   │   ├──  http/
│   │   ├──  grpc/
│   ├──  service/       This package contain business logic, this package get input request from controller, this package use repository for things related to persistence.
│   ├──  repository/    This package only for things related to persistence (CRUD database/redis/etc).
├── go.mod
└── go.sum
`
